package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jmechavez/gotest2023/domain"
	"github.com/jmechavez/gotest2023/service"
)

func Start() {
	// Create a new router using the Gorilla Mux package
	//mux := http.NewServeMux()
	router := mux.NewRouter()

	// Wiring: Initialize the PlayerHandlers with a new PlayerService and PlayerRepository
	//phS := PlayerHandlers{service.NewPlayerService(domain.NewPlayerRepositoryStub())}
	ph := PlayerHandlers{service.NewPlayerService(domain.NewPlayerRepositoryDb())}

	// Define routes
	router.HandleFunc("/greet", GreetHandler).Methods(http.MethodGet)
	router.HandleFunc("/players", ph.GetAllPlayers).Methods(http.MethodGet)
	router.HandleFunc("/players/{player_id:[0-9]+}", ph.GetPlayer).Methods(http.MethodGet)
	//router.HandleFunc("/players", phS.GetAllPlayers).Methods(http.MethodGet)

	// Starting server

	// Retrieve the server address and port from environment variables
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")

	// Check if either address or port is empty
	if address == "" || port == "" {
		log.Fatal("Error: SERVER_ADDRESS or SERVER_PORT is empty")
	}

	// Combine address and port to form the server address
	serverAdd := fmt.Sprintf("%s:%s", address, port)

	// Start the HTTP server
	err := http.ListenAndServe(serverAdd, router)

	// Check for errors during server startup
	if err != nil {
		// Log the error and exit the program
		log.Fatal("Error starting server:", err)
	}
}
