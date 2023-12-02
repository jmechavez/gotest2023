package domain

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type PlayerRepositoryDB struct {
	client *sql.DB
}

// ById implements PlayerRepository.
//func (*PlayerRepositoryDB) ById(string) (*Player, error) {
//	panic("unimplemented")
//}

func (d *PlayerRepositoryDB) FindAll() ([]Player, error) {

	findAllSql := "select player_id,name,age,game,status from players"

	rows, err := d.client.Query(findAllSql)
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

func (d PlayerRepositoryDB) ById(id string) (*Player, error) {
	playerSql := "select player_id, name, age, game, status from players where player_id = ?"
	row := d.client.QueryRow(playerSql, id)

	var p Player
	err := row.Scan(&p.Id, &p.Name, &p.Age, &p.Game, &p.Status)

	if err == sql.ErrNoRows {
		log.Println("No rows found for player with ID:", id)
		return nil, nil // Return nil for both player and error
	} else if err != nil {
		log.Println("Error while scanning:", err.Error())
		return nil, err
	}

	return &p, nil
}

// ! connect to mySQL database
func NewPlayerRepositoryDb() *PlayerRepositoryDB {
	client, err := sql.Open("mysql", "root:3@tcp(localhost:33061)/gotest2023")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return &PlayerRepositoryDB{client}
}
