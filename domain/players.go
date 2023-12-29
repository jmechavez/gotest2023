package domain

import (
	"github.com/jmechavez/gotest2023/dto"
	"github.com/jmechavez/gotest2023/errorCust"
)

// Player represents a player in the system.
type Player struct {
	PlayerId string `db:"player_id"`
	Name     string
	Age      int
	Game     string
	Status   string
}

// statusAsText converts the numeric status to a human-readable text.
func (p Player) statusAsText() string {
	statusAsText := "active"
	if p.Status == "0" {
		statusAsText = "inactive"
	}
	return statusAsText
}

// Todto converts Player to PlayerResponse DTO.
func (p Player) Todto() dto.PlayerResponse {
	return dto.PlayerResponse{
		PlayerId: p.PlayerId,
		Name:     p.Name,
		Age:      p.Age,
		Game:     p.Game,
		Status:   p.statusAsText(),
	}
}

// PlayerRepository defines the interface for interacting with player data.
type PlayerRepository interface {
	FindAll() ([]Player, *errorCust.AppError)   // Retrieve all players
	ById(string) (*Player, *errorCust.AppError) // Retrieve a player by ID
}
