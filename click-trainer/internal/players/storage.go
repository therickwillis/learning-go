package players

import (
	utility "clicktrainer/internal"
	"sync"
)

var (
	players   = make(map[string]*Player)
	allReady  = false
	playersMu sync.Mutex
)

func Add(id string, name string) error {
	playersMu.Lock()
	players[id] = &Player{ID: id, Name: name, Color: utility.RandomColorHex()}
	playersMu.Unlock()
	return nil
}

func Get(id string) *Player {
	playersMu.Lock()
	defer playersMu.Unlock()
	return players[id]
}

func GetList() []*Player {
	playersMu.Lock()
	playerList := make([]*Player, 0, len(players))
	for _, p := range players {
		playerList = append(playerList, p)
	}
	playersMu.Unlock()
	return playerList
}

func UpdateScore(id string, points int) *Player {
	playersMu.Lock()
	defer playersMu.Unlock()
	if p, e := players[id]; e {
		p.Score += points
		return p
	}
	return nil
}

func SetReady(id string, isReady bool) *Player {
	playersMu.Lock()
	defer playersMu.Unlock()
	if p, e := players[id]; e {
		p.Ready = isReady
		return p
	}
	return nil
}

func AllReady() bool {
	playersMu.Lock()
	defer playersMu.Unlock()
	all := true
	for _, player := range players {
		if all {
			all = player.Ready
		}
		break
	}
	return all
}

func ValidateSession(sessionId string) bool {
	playersMu.Lock()
	defer playersMu.Unlock()
	_, exists := players[sessionId]
	return exists
}
