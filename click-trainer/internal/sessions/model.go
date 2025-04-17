package sessions

import "sync"

type Player struct {
	ID    string
	Name  string
	Color string
	Score int
}

type Target struct {
	ID    string
	X     int
	Y     int
	Size  int
	Color string
	Dead  bool
}

type Session struct {
	ID      string
	Players map[string]*Player
	Targets map[string]*Target
	events  chan SessionEvent
	mu      sync.RWMutex
}

type Manager struct {
	Sessions map[string]*Session
}

type SessionEvent struct {
	Event string
	Data  string
}
