// players/player.go
package domain

import "github.com/jmechavez/gotest2023/errorCust"

// Player represents a player in the system.
type Player struct {
	Id     string `json:"player_id" xml:"id" db:"player_id"`
	Name   string `json:"full_name" xml:"name"`
	Age    int    `json:"age" xml:"age"`
	Game   string `json:"game" xml:"game"`
	Status string `json:"status" xml:"status"`
}

// PlayerRepository defines the interface for interacting with player data.
type PlayerRepository interface {
	FindAll() ([]Player, *errorCust.AppError)   // Retrieve all players
	ById(string) (*Player, *errorCust.AppError) // Retrieve a player by ID
}
