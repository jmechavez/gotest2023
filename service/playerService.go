package service

import "github.com/jmechavez/gotest2023/db"

type PlayerService interface {
	GetAllPlayer() ([]db.Player, error)
}

type DefaultPlayerService struct {
	repo db.PlayerRepository
}

func (s DefaultPlayerService) GetAllPlayer() ([]db.Player, error) {
	return s.repo.FindAll()
}

func NewPlayerService(repository db.PlayerRepository) DefaultPlayerService {
	return DefaultPlayerService{repository}
}
