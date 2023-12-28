//rest api handlers

package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmechavez/gotest2023/service"
)

func GreetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Player")
}

type PlayerHandlers struct {
	service service.PlayerService
}

// xml and jason
// func writeResponse(w http.ResponseWriter, code int, data interface{}) {
// 	w.Header().Add("Content-Type", "application/json")
// 	w.WriteHeader(code)
// 	//XML
// 	if r.Header.Get("Content-Type") == "application/xml" {
// 		w.Header().Add("Content-Type", "application/xml")
// 		xml.NewEncoder(w).Encode(players)
// 	} else {
// 		//JSON
// 		w.Header().Add("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(players)
// 	}
// }

func (ph *PlayerHandlers) GetAllPlayers(w http.ResponseWriter, r *http.Request) {
	// playerTest := []players.Player{
	// 	{Id: uuid.New(), Name: "John", Age: 21, Game: "Basketball"},

	players, err := ph.service.GetAllPlayers()
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, players)
		// w.Header().Add("Content-Type", "application/json")
		// w.WriteHeader(http.StatusOK)
		// json.NewEncoder(w).Encode(player)
	}
}

func (ph *PlayerHandlers) GetPlayer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["player_id"]

	player, err := ph.service.GetPlayerByID(id)
	if err != nil {
		// w.WriteHeader(err.Code)
		// fmt.Fprintf(w, err.Message)
		// w.Header().Add("Content-Type", "application/json")
		// w.WriteHeader(err.Code)
		// json.NewEncoder(w).Encode(err.AsMessage())

		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, player)
		// w.Header().Add("Content-Type", "application/json")
		// w.WriteHeader(http.StatusOK)
		// json.NewEncoder(w).Encode(player)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
