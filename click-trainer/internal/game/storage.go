package gamedata

import (
	"clicktrainer/internal/players"
	"clicktrainer/internal/targets"
)

type GameData struct {
	Players []*players.Player
	Targets []*targets.Target
}

func Get() GameData {
	return GameData{
		Players: players.GetList(),
		Targets: targets.GetList(),
	}
}
