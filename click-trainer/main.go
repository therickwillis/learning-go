package main

import (
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

	tmpl = template.Must(template.ParseFiles("templates/game.html", "templates/join.html"))

	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/register", handleRegister)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Server listening on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err, err.Error())
	}
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[Handle:Index] Request Received")

	cookie, err := r.Cookie("player_id")
	if err == nil && validateSession(cookie.Value) {
		gameData := getGameData()
		if err := tmpl.ExecuteTemplate(w, "game", gameData); err != nil {
			fmt.Println(err.Error())
			http.Error(w, "Error rendering game view", http.StatusInternalServerError)
		}
		return
	}

	if err := tmpl.ExecuteTemplate(w, "join", nil); err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func validateSession(sessionId string) bool {
	playersMu.Lock()
	defer playersMu.Unlock()
	_, exists := players[sessionId]
	return exists
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

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
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
