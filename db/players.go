// players/player.go
package db

import "github.com/google/uuid"

type Player struct {
	Id   uuid.UUID `json:"id" xml: xml:"id"`
	Name string    `json:"full_name" xml:"name"`
	Age  int       `json:"age" xml:"age"`
	Game string    `json:"game" xnk:"game"`
}

type PlayerRepository interface {
	FindAll() ([]Player, error)
}
