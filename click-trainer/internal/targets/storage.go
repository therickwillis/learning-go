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
	gameHeight = 400
	gameWidth  = 600
	targetSize = 50
)

func Add() int {
	targetsMu.Lock()
	id := nextID
	nextID++
	targets[id] = &Target{
		ID: id,
		X:  rand.Intn(gameWidth - targetSize),
		Y:  rand.Intn(gameHeight - targetSize),
	}
	targetsMu.Unlock()
	return id
}

func GetList() []*Target {
	targetsMu.Lock()
	targetList := make([]*Target, 0, len(targets))
	for _, t := range targets {
		targetList = append(targetList, t)
	}
	targetsMu.Unlock()
	return targetList
}
