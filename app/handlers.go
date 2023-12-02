//rest api handlers

package app

import (
	"encoding/json"
	"encoding/xml"
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

func (ph *PlayerHandlers) GetAllPlayers(w http.ResponseWriter, r *http.Request) {
	// playerTest := []players.Player{
	// 	{Id: uuid.New(), Name: "John", Age: 21, Game: "Basketball"},

	players, _ := ph.service.GetAllPlayer()

	//XML
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(players)
	} else {
		//JSON
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(players)
	}
}
func (ph *PlayerHandlers) GetPlayer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["player_id"]

	player, err := ph.service.GetPlayer(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, err.Error())
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(player)
	}
}
