package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"sync"

	"github.com/google/uuid"
)

const (
	gameWidth  = 600
	gameHeight = 400
	targetSize = 50
)

type Player struct {
	Name  string
	Score int
}

type Target struct {
	X int
	Y int
}

type GameData struct {
	Players []*Player
	Target  Target
}

var (
	players   = make(map[string]*Player)
	playersMu sync.Mutex

	currentTarget Target
	targetMu      sync.Mutex

	tmpl *template.Template
)

func main() {
	targetMu.Lock()
	currentTarget = Target{
		X: rand.Intn(gameWidth - targetSize),
		Y: rand.Intn(gameHeight - targetSize),
	}
	targetMu.Unlock()

	tmpl = template.Must(template.ParseFiles("templates/index.html", "templates/game.html"))

	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/register", handleRegister)

	fmt.Println("Server Listening on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[Handle:Index] Request Received")
	if err := tmpl.Execute(w, nil); err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	cookie, err := r.Cookie("player_id")
	if err != nil {
		return
	}
	fmt.Println("Found Session", cookie.Value)

	playersMu.Lock()
	if players[cookie.Value] != nil {
		enterGame(w)
	}
	playersMu.Unlock()
}

func handleRegister(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[Handle:Register] Request Received")
	if err := r.ParseForm(); err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id := uuid.New().String()
	name := r.FormValue("name")
	fmt.Println("Registering name:", name)

	http.SetCookie(w, &http.Cookie{
		Name:     "player_id",
		Value:    id,
		Path:     "/",
		HttpOnly: true,
	})

	playersMu.Lock()
	players[id] = &Player{Name: name}
	playersMu.Unlock()

	enterGame(w)
}

func enterGame(w http.ResponseWriter) {
	gameData := getGameData()
	if err := tmpl.ExecuteTemplate(w, "game", gameData); err != nil {
		fmt.Println(err.Error())
	}

	var buf bytes.Buffer
	buf.WriteString(`<div id="join" hx-swap-oob="true" style="display:none"></div>`)
	w.Header().Set("Content-Type", "text/html")
	if _, err := w.Write(buf.Bytes()); err != nil {
		fmt.Println(err.Error())
	}
}

func getGameData() GameData {
	playersMu.Lock()
	defer playersMu.Unlock()
	targetMu.Lock()
	defer targetMu.Unlock()

	playerList := make([]*Player, 0, len(players))
	for _, p := range players {
		playerList = append(playerList, p)
	}
	return GameData{
		Players: playerList,

		Target: currentTarget,
	}
}
