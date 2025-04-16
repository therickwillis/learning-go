package sessions

import (
	"clicktrainer/internal/players"
	"clicktrainer/internal/targets"
)

type Session struct {
	Players []*players.Player
	Targets []*targets.Target
}
