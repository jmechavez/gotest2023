package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmechavez/gotest2023/domain"
	"github.com/jmechavez/gotest2023/service"
)

func Start() {

	//mux := http.NewServeMux()
	router := mux.NewRouter()

	//wiring
	//ph := PlayerHandlers{service.NewPlayerService(domain.NewPlayerRepositoryStub())}
	ph := PlayerHandlers{service.NewPlayerService(domain.NewPlayerRepositoryDb())}

	// define routes
	router.HandleFunc("/greet", GreetHandler).Methods(http.MethodGet)
	router.HandleFunc("/players", ph.GetAllPlayers).Methods(http.MethodGet)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))

}
