package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmechavez/gotest2023/domain"
	"github.com/jmechavez/gotest2023/service"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func Start() {
	// Create a new router using the Gorilla Mux package
	//mux := http.NewServeMux()
	router := mux.NewRouter()

	// Wiring: Initialize the PlayerHandlers with a new PlayerService and PlayerRepository

	dbClient := getDbclient()
	playerRepositoryDb := domain.NewPlayerRepositoryDb(dbClient)
	accountRepositoryDb := domain.NewPlayerRepositoryDb(dbClient)
	//accountRepositoryDB = domain.NewAccountRepositoryDB(dbClient)
	ph := PlayerHandlers{service.NewPlayerService(playerRepositoryDb)}
	//phS := PlayerHandlers{service.NewPlayerService(domain.NewPlayerRepositoryStub())}
	ah := AccountHandler{service.NewAccountService(accountRepositoryDb)}

	// Define routes
	router.
		HandleFunc("/greet", GreetHandler).
		Methods(http.MethodGet)
	router.
		HandleFunc("/players", ph.GetAllPlayers).
		Methods(http.MethodGet)
	router.
		HandleFunc("/players/{player_id:[0-9]+}", ph.GetPlayer).
		Methods(http.MethodGet)
	router.
		HandleFunc("/players/{player_id:[0-9]+}/account", ah.NewAccount).
		Methods(http.MethodPost)
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

func getDbclient() *sqlx.DB {

	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Retrieve database connection details from environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Construct the data source name for the MySQL connection
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Open a connection to the MySQL database
	client, err := sqlx.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err.Error())
		return nil
	}

	// Set connection pool settings.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
