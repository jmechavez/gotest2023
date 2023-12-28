//service interfaces - Business side / domain

package service

import (
	"github.com/jmechavez/gotest2023/domain"
	"github.com/jmechavez/gotest2023/errorCust"
)

// PlayerService defines the interface for interacting with player data.
type PlayerService interface {
	GetAllPlayers() ([]domain.Player, *errorCust.AppError)      // Get all players
	GetPlayerByID(string) (*domain.Player, *errorCust.AppError) // Get a player by ID
}

// DefaultPlayerService is the default implementation of the PlayerService interface.
type DefaultPlayerService struct {
	repo domain.PlayerRepository
}

// GetAllPlayers retrieves all players.
func (s DefaultPlayerService) GetAllPlayers() ([]domain.Player, *errorCust.AppError) {
	return s.repo.FindAll()
}

// GetPlayerByID retrieves a player by their ID.
func (s DefaultPlayerService) GetPlayerByID(id string) (*domain.Player, *errorCust.AppError) {
	return s.repo.ById(id)
}

// NewPlayerService creates a new instance of DefaultPlayerService.
func NewPlayerService(repository domain.PlayerRepository) DefaultPlayerService {
	return DefaultPlayerService{repository}
}
