package sessions

import (
	"clicktrainer/internal/players"
	"clicktrainer/internal/targets"
)

var (
	sessions = make(map[string]*Session)
)

func Get() Session {
	return Session{
		Players: players.GetList(),
		Targets: targets.GetList(),
	}
}
