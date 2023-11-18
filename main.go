package main

import (
	"fmt"

	players "github.com/jmechavez/gotest2023/db"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println(players.Player1)

	Player2 := players.Player{
		Name: "Hell",
		Age:  19,
		Game: "Basketball",
	}

	fmt.Println(Player2.Age)
	fmt.Println(players.Player1.Age)
}
