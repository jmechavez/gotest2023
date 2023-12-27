//service interfaces - Business side / domain

package service

import (
	"github.com/jmechavez/gotest2023/domain"
	"github.com/jmechavez/gotest2023/errorCust"
)

type PlayerService interface {
	GetAllPlayer() ([]domain.Player, *errorCust.AppError)   //create method  get all players
	GetPlayer(string) (*domain.Player, *errorCust.AppError) //create method get alll player id only
}

type DefaultPlayerService struct {
	repo domain.PlayerRepository
}

func (s DefaultPlayerService) GetAllPlayer() ([]domain.Player, *errorCust.AppError) {
	return s.repo.FindAll()
}

func (s DefaultPlayerService) GetPlayer(id string) (*domain.Player, *errorCust.AppError) {
	return s.repo.ById(id)
}

func NewPlayerService(repository domain.PlayerRepository) DefaultPlayerService {
	return DefaultPlayerService{repository}
}
