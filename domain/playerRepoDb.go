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
	"github.com/joho/godotenv"
)

type PlayerRepositoryDB struct {
	client *sql.DB
}

// ById implements PlayerRepository.
//
//	func (*PlayerRepositoryDB) ById(string) (*Player, error) {
//		panic("unimplemented")
//	}
//
// * Find all players
func (d *PlayerRepositoryDB) FindAll() ([]Player, error) {

	findAllSql := "select player_id,name,age,game,status from players" //check all players in SQLdatabase
	rows, err := d.client.Query(findAllSql)                            //send querry to the database

	if err != nil {
		log.Println("Error while querying player table" + err.Error())
		return nil, err
	}

	players := make([]Player, 0)
	for rows.Next() {
		var p Player
		err := rows.Scan(&p.Id, &p.Name, &p.Age, &p.Game, &p.Status)
		if err != nil {
			log.Println("Error while scanning players table" + err.Error())
			return nil, err
		}
		players = append(players, p)
	}

	return players, nil
}

// * Find Players by ID
func (d PlayerRepositoryDB) ById(id string) (*Player, *errorCust.AppError) {
	playerSql := "select player_id, name, age, game, status from players where player_id = ?" //check all players table in SQLdatabase
	row := d.client.QueryRow(playerSql, id)                                                   //send querryrow to the database

	var p Player
	err := row.Scan(&p.Id, &p.Name, &p.Age, &p.Game, &p.Status)

	if err == sql.ErrNoRows {
		log.Println("No rows found for players with ID:", id)
		return nil, errorCust.NewNotFoundError("Player not found")
	} else if err != nil {
		log.Println("Error while scanning:", err.Error())
		return nil, errorCust.NewUnexpectedError("Unexpected Error")
	}
	return &p, nil
}

// ! connect to mySQL database with godotenv file
func NewPlayerRepositoryDb() *PlayerRepositoryDB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	client, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	// Set connection pool settings.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return &PlayerRepositoryDB{client}
}
