package handlers

import (
	"fmt"
	"net/http"
	"strings"
	"text/template"
)

var pages = map[string]*template.Template{}

func init() {
	baseTmpl := template.Must(template.ParseFiles("templates/base.html"))
	for _, child := range []string{"register", "join", "lobby", "game"} {
		t, _ := baseTmpl.Clone()
		_, err := t.ParseFiles(fmt.Sprintf("templates/%s.html", child))
		if err != nil {
			fmt.Println("Error parsing templates", err)
		}
		pages[child] = t
	}
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[Handle:Index] Request Received")

	idCookie, _ := r.Cookie("player_id")
	nameCookie, _ := r.Cookie("player_name")

	if err := idCookie.Valid(); err == nil {
		if err := nameCookie.Valid(); err == nil {
			if err := pages["join"].ExecuteTemplate(w, "base", nameCookie.Value); err != nil {
				fmt.Println(err.Error())
				http.Error(w, "Error rendering join template", http.StatusInternalServerError)
			}
			return
		}
	}

	if err := pages["register"].ExecuteTemplate(w, "base", nil); err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func HandleGame(w http.ResponseWriter, r *http.Request) {
	if err := validateCookies(w, r); err != nil {
		fmt.Println("[Handle:Game] failed cookie validation")
	}

	// establish the session
	parts := strings.Split(r.URL.Path, "/")
	id := parts[2]
	if len(id) > 0 {
		fmt.Printf("Game Session Found:%s\n", id)
	} else {
		// newId := M
		// sessions.NewSession(newId)
		// http.Redirect(w, r, fmt.Sprintf("/g/%s", newId), http.StatusSeeOther)
		return
	}

	// pass the game data to the game template
	if err := tmpl.ExecuteTemplate(w, "game", nil); err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// func handleTarget(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("[Handle:Target] Request Received")
// 	cookie, err := r.Cookie("player_id")
// 	if err != nil {
// 		http.Error(w, "Not Registered", http.StatusBadRequest)
// 		return
// 	}
// 	parts := strings.Split(r.URL.Path, "/")

// 	// Target
// 	strTargetId := parts[2]
// 	targetId, err := strconv.Atoi(strTargetId)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	targets.Kill(targetId)
// 	time.AfterFunc(500*time.Millisecond, func() {
// 		newTarget := targets.Add()
// 		var buf bytes.Buffer
// 		if err := tmpl.ExecuteTemplate(&buf, "target", newTarget); err != nil {
// 			log.Println(err)
// 		}

// 		fmt.Println("sending OOB for new Target")
// 		BroadcastOOB("newTarget", buf.String())
// 	})

// 	// Player
// 	strPoints := parts[3]
// 	points, err := strconv.Atoi(strPoints)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	playerId := cookie.Value
// 	player := players.UpdateScore(playerId, points)

// 	targetOOB := fmt.Sprintf(`<div id="target_%d" hx-swap-oob="delete"></div>`, targetId)
// 	playerOOB := fmt.Sprintf(`<div id="player_score_%s" hx-swap-oob="innerHTML">%d</div>`, player.ID, player.Score)
// 	BroadcastOOB("swap", targetOOB+playerOOB)
// }

// func handleRegister(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("[Handle:Register] Request Received")
// 	if err := r.ParseForm(); err != nil {
// 		fmt.Println(err.Error())
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	id := uuid.New().String()
// 	name := r.FormValue("name")
// 	fmt.Println("Registering name:", name)

// 	http.SetCookie(w, &http.Cookie{
// 		Name:     "player_id",
// 		Value:    id,
// 		Path:     "/",
// 		HttpOnly: true,
// 	})
// 	http.SetCookie(w, &http.Cookie{
// 		Name:     "player_name",
// 		Value:    name,
// 		Path:     "/",
// 		HttpOnly: true,
// 	})

// 	if err := players.Add(id, name); err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	var buf bytes.Buffer
// 	if err := tmpl.ExecuteTemplate(&buf, "scoreboard", players.GetList()); err != nil {
// 		log.Println(err)
// 	}
// 	//playersOOB := fmt.Sprintf(`<div id="scoreboard" hx-oob-swap="outerHTML">%s</div>`, buf.String())
// 	//BroadcastOOB("swap", playersOOB)
// 	BroadcastOOB("scoreboard", buf.String())

// 	http.Redirect(w, r, "/", http.StatusSeeOther)
// }

func validateCookies(w http.ResponseWriter, r *http.Request) error {
	idCookie, _ := r.Cookie("player_id")
	if err := idCookie.Valid(); err != nil {
		fmt.Println("[ValidateCookie] No player cookie found")
		http.Redirect(w, r, "/", http.StatusPermanentRedirect)
		return err
	}
	return nil
}
