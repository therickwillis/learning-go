package players

import "sync"

var (
	players   = make(map[string]*Player)
	playersMu sync.Mutex
)

func Add(id string, name string) error {
	playersMu.Lock()
	players[id] = &Player{ID: id, Name: name}
	playersMu.Unlock()
	return nil
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

func UpdateScore(id string, points int) {
	playersMu.Lock()
	if p, e := players[id]; e {
		p.Score += points
	}
	playersMu.Unlock()
}

func ValidateSession(sessionId string) bool {
	playersMu.Lock()
	defer playersMu.Unlock()
	_, exists := players[sessionId]
	return exists
}
