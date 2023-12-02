// players/player.go
package domain

type Player struct {
	Id     string `json:"player_id" xml:"id"`
	Name   string `json:"full_name" xml:"name"`
	Age    int    `json:"age" xml:"age"`
	Game   string `json:"game" xml:"game"`
	Status string `json:"status" xml:"status"`
}

type PlayerRepository interface {
	FindAll() ([]Player, error)   //retreive all players
	ById(string) (*Player, error) //retrieve player_id only
}
