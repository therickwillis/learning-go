package server

import (
	"bytes"
	"clicktrainer/internal/players"
	"clicktrainer/internal/targets"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/google/uuid"
)

var pages = map[string]*template.Template{}

func init() {
	baseTmpl := template.Must(template.ParseFiles("templates/base.html"))
	for _, child := range []string{"register", "join", "lobby", "game", "recap"} {
		t, _ := baseTmpl.Clone()
		t.ParseFiles(fmt.Sprintf("templates/%s.html", child))
		pages[child] = t
	}
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[Handle:Index] Request Received")

	idCookie, _ := r.Cookie("player_id")
	nameCookie, _ := r.Cookie("player_name")

	if err := idCookie.Valid(); err == nil {
		if err := nameCookie.Valid(); err == nil {
			if err := pages["join"].ExecuteTemplate(w, "base", nameCookie.Value); err != nil {
				fmt.Println(err.Error())
				http.Error(w, "Error rendering game view", http.StatusInternalServerError)
			}
			return
		}
	}

	if err := pages["register"].ExecuteTemplate(w, "base", nil); err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleGame(w http.ResponseWriter, r *http.Request) {
	if err := tmpl.ExecuteTemplate(w, "game", nil); err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// func handlePoll(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("[Handle:Poll] Request Received")
// 	if err := tmpl.ExecuteTemplate(w, "gameContent", gamedata.Get()); err != nil {
// 		log.Println(err)
// 	}
// }

func handleTarget(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[Handle:Target] Request Received")
	cookie, err := r.Cookie("player_id")
	if err != nil {
		http.Error(w, "Not Registered", http.StatusBadRequest)
		return
	}
	parts := strings.Split(r.URL.Path, "/")

	// Target
	strTargetId := parts[2]
	targetId, err := strconv.Atoi(strTargetId)
	if err != nil {
		log.Println(err)
	}
	targets.Kill(targetId)
	time.AfterFunc(500*time.Millisecond, func() {
		newTarget := targets.Add()
		var buf bytes.Buffer
		if err := tmpl.ExecuteTemplate(&buf, "target", newTarget); err != nil {
			log.Println(err)
		}

		fmt.Println("sending OOB for new Target")
		BroadcastOOB("newTarget", buf.String())
	})

	// Player
	strPoints := parts[3]
	points, err := strconv.Atoi(strPoints)
	if err != nil {
		log.Println(err)
	}
	playerId := cookie.Value
	player := players.UpdateScore(playerId, points)

	targetOOB := fmt.Sprintf(`<div id="target_%d" hx-swap-oob="delete"></div>`, targetId)
	playerOOB := fmt.Sprintf(`<div id="player_score_%s" hx-swap-oob="innerHTML">%d</div>`, player.ID, player.Score)
	BroadcastOOB("swap", targetOOB+playerOOB)
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
	http.SetCookie(w, &http.Cookie{
		Name:     "player_name",
		Value:    name,
		Path:     "/",
		HttpOnly: true,
	})

	if err := players.Add(id, name); err != nil {
		fmt.Println(err.Error())
	}

	var buf bytes.Buffer
	if err := tmpl.ExecuteTemplate(&buf, "scoreboard", players.GetList()); err != nil {
		log.Println(err)
	}
	//playersOOB := fmt.Sprintf(`<div id="scoreboard" hx-oob-swap="outerHTML">%s</div>`, buf.String())
	//BroadcastOOB("swap", playersOOB)
	BroadcastOOB("scoreboard", buf.String())

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
