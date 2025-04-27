package sessions

import (
	"sync"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

type ctxKey string

const (
	sessionCtxKey = ctxKey("session")
)

type Manager struct {
	sessions map[string]*Session
	mu       sync.RWMutex
}

func NewManager() *Manager {
	return &Manager{
		sessions: make(map[string]*Session),
	}
}

func (m *Manager) NewRandomID() string {
	id, _ := gonanoid.New(6)
	return id
}

func (m *Manager) CreateSession(id string) *Session {
	m.mu.Lock()
	defer m.mu.Unlock()
	s := NewSession(id)
	m.sessions[id] = s
	return s
}

func (m *Manager) Get(id string) (*Session, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	s, ok := m.sessions[id]
	return s, ok
}

func (m *Manager) RemoveSession(id string) {
	m.mu.Lock()
	delete(m.sessions, id)
	m.mu.Unlock()
}
