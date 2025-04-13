package server

import (
	gamedata "clicktrainer/internal/game"
	"clicktrainer/internal/players"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[Handle:Index] Request Received")

	cookie, err := r.Cookie("player_id")
	if err == nil && players.ValidateSession(cookie.Value) {
		gameData := gamedata.Get()
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

func handleTarget(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[Handle:Target] Request Received")
	cookie, err := r.Cookie("player_id")
	if err != nil {
		http.Error(w, "Not Registered", http.StatusBadRequest)
		return
	}
	id := cookie.Value

	players.UpdateScore(id, 1)

	// targetMu.Lock()
	// currentTarget = Target{
	// 	X: rand.Intn(gameWidth - targetSize),
	// 	Y: rand.Intn(gameHeight - targetSize),
	// }
	// targetMu.Unlock()
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

	if err := players.Add(id, name); err != nil {
		fmt.Println(err.Error())
	}
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
