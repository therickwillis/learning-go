package targets

import (
	"math/rand"
	"sync"
)

var (
	targets   = make(map[int]*Target)
	nextID    = 1
	targetsMu sync.Mutex
)

const (
	gameHeight    = 400
	gameWidth     = 600
	minTargetSize = 50
	maxTargetSize = 100
)

func Add() int {
	targetsMu.Lock()
	id := nextID
	nextID++
	targetSize := rand.Intn(maxTargetSize-minTargetSize) + minTargetSize
	targets[id] = &Target{
		ID:   id,
		X:    rand.Intn(gameWidth - targetSize),
		Y:    rand.Intn(gameHeight - targetSize),
		Size: targetSize,
	}
	targetsMu.Unlock()
	return id
}

func Kill(id int) {
	targetsMu.Lock()
	if t, e := targets[id]; e {
		t.Dead = true
	}
	targetsMu.Unlock()
}

func GetList() []*Target {
	targetsMu.Lock()
	targetList := make([]*Target, 0, len(targets))
	for _, t := range targets {
		if !t.Dead {
			targetList = append(targetList, t)
		}
	}
	targetsMu.Unlock()
	return targetList
}
