package db

import "github.com/google/uuid"

type PlayerRepositoryStub struct {
	players []Player
}

func (s PlayerRepositoryStub) FindAll() ([]Player, error) {
	return s.players, nil
}

func NewPlayerRepositoryStub() PlayerRepositoryStub {
	players := []Player{
		{Id: uuid.New(), Name: "John", Age: 21, Game: "Basketball"},
		{Id: uuid.New(), Name: "Michael", Age: 22, Game: "Volleyball"},
	}
	return PlayerRepositoryStub{players}
}
