package sessions

import (
	"fmt"
	"math/rand"
	"sync"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

var (
	sessions      = make(map[string]*Session)
	gameHeight    = 400
	gameWidth     = 600
	minTargetSize = 50
	maxTargetSize = 100
)

func NewSession(id string) *Session {
	session := &Session{
		ID:      id,
		Players: make(map[string]*Player),
		Targets: make(map[string]*Target),
		events:  make(chan SessionEvent),
		mu:      sync.RWMutex{},
	}
	return session
}

func (s *Session) AddPlayer(id string, name string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Players[id] = &Player{ID: id, Name: name}
	s.broadcast("addPlayer", id)
}

func (s *Session) RemovePlayer(id string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.Players, id)
	s.broadcast("removePlayer", id)
}

func (s *Session) PlayersSnapshot() []*Player {
	s.mu.RLock()
	defer s.mu.RUnlock()
	list := make([]*Player, 0, len(s.Players))
	for _, p := range s.Players {
		list = append(list, p)
	}
	return list
}

func (s *Session) AddTarget() {
	s.mu.Lock()
	defer s.mu.Unlock()
	if id, err := gonanoid.New(6); err != nil {
		targetSize := rand.Intn(maxTargetSize-minTargetSize) + minTargetSize
		s.Targets[id] = &Target{
			ID:    id,
			X:     rand.Intn(gameWidth - targetSize),
			Y:     rand.Intn(gameHeight - targetSize),
			Color: RandomColorHex(),
			Size:  targetSize,
		}
	}

}

func (s *Session) broadcast(evt, data string) {
	select {
	case s.events <- SessionEvent{Event: evt, Data: data}:
	default:
	}
}

func RandomColorHex() string {
	r := uint8(rand.Intn(248) + 4)
	g := uint8(rand.Intn(248) + 4)
	b := uint8(rand.Intn(248) + 4)
	return fmt.Sprintf("#%02x%02x%02x", r, g, b)
}
