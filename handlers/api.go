package api

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	players "github.com/jmechavez/gotest2023/db"
)

func GreetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Player")
}

func GetAllPlayers(w http.ResponseWriter, r *http.Request) {
	playerTest := []players.Player{
		{Name: "John", Age: 21, Game: "Basketball"},
	}

	//XML
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(playerTest)
	} else {
		//JSON
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(playerTest)
	}
}
