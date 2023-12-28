//database adapter

package domain

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmechavez/gotest2023/errorCust"
	"github.com/jmechavez/gotest2023/logger"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

type PlayerRepositoryDB struct {
	client *sqlx.DB
}

// ById implements PlayerRepository.
//
//	func (*PlayerRepositoryDB) ById(string) (*Player, error) {
//		panic("unimplemented")
//	}
//
// * Find all players
func (d *PlayerRepositoryDB) FindAll() ([]Player, *errorCust.AppError) {

	findAllSql := "select player_id,name,age,game,status from players" //check all players in SQLdatabase
	rows, err := d.client.Query(findAllSql)                            //send querry to the database

	if err != nil {
		logger.Error("Error while querying player table" + err.Error())
		return nil, errorCust.NewNotFoundError("Unexpected database error")
	}

	defer rows.Close()

	players := make([]Player, 0)
	err = sqlx.StructScan(rows, &players)
	if err != nil {
		logger.Error("Error while scanning players table" + err.Error())
		return nil, errorCust.NewNotFoundError("Unexpected database error")
	}

	// for rows.Next() {
	// 	var p Player
	// 	err := rows.Scan(&p.Id, &p.Name, &p.Age, &p.Game, &p.Status)
	// 	if err != nil {
	// 		logger.Error("Error while scanning players table" + err.Error())
	// 		return nil, errorCust.NewNotFoundError("Unexpected database error")
	// 	}
	// 	players = append(players, p)
	// }

	return players, nil
}

// FindAll retrieves all players from the database.
func (d PlayerRepositoryDB) ById(id string) (*Player, *errorCust.AppError) {
	playerSql := "select player_id, name, age, game, status from players where player_id = ?" //check all players table in SQLdatabase
	//row := d.client.QueryRow(playerSql, id)                                                   //send querryrow to the database

	var p Player
	err := d.client.Get(&p, playerSql, id)
	//err := row.Scan(&p.Id, &p.Name, &p.Age, &p.Game, &p.Status)

	if err == sql.ErrNoRows {
		log.Println("No rows found for players with ID:", id)
		return nil, errorCust.NewNotFoundError("Player not found")
	} else if err != nil {
		log.Println("Error while scanning:", err.Error())
		return nil, errorCust.NewUnexpectedError("Unexpected Error")
	}
	return &p, nil
}

// ! NewPlayerRepositoryDb creates a new instance of PlayerRepositoryDB
// and connects to a MySQL database using environment variables loaded from a .env file.
func NewPlayerRepositoryDb() *PlayerRepositoryDB {
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

	// Return the PlayerRepositoryDB instance with the connected database client
	return &PlayerRepositoryDB{client}
}
