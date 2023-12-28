// players/player.go
package domain

import "github.com/jmechavez/gotest2023/errorCust"

type Player struct {
	Id     string `json:"player_id" xml:"id" db:"player_id"`
	Name   string `json:"full_name" xml:"name"`
	Age    int    `json:"age" xml:"age"`
	Game   string `json:"game" xml:"game"`
	Status string `json:"status" xml:"status"`
}

type PlayerRepository interface {
	FindAll() ([]Player, *errorCust.AppError)   //retreive all players
	ById(string) (*Player, *errorCust.AppError) //retrieve player_id only
}
