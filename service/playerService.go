package service

import "github.com/jmechavez/gotest2023/domain"

type PlayerService interface {
	GetAllPlayer() ([]domain.Player, error)
}

type DefaultPlayerService struct {
	repo domain.PlayerRepository
}

func (s DefaultPlayerService) GetAllPlayer() ([]domain.Player, error) {
	return s.repo.FindAll()
}

func NewPlayerService(repository domain.PlayerRepository) DefaultPlayerService {
	return DefaultPlayerService{repository}
}
