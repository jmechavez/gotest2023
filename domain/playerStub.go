package domain

type PlayerRepositoryStub struct {
	players []Player
}

func (s PlayerRepositoryStub) FindAll() ([]Player, error) {
	return s.players, nil
}

func NewPlayerRepositoryStub() PlayerRepositoryStub {
	players := []Player{
		{Id: "2005", Name: "John", Age: 21, Game: "Basketball", Status: "1"},
		{Id: "2004", Name: "Michael", Age: 22, Game: "Volleyball", Status: "1"},
	}
	return PlayerRepositoryStub{players}
}
