package main

import (
	app "github.com/jmechavez/gotest2023/app"
	"github.com/jmechavez/gotest2023/logger"
)

func main() {

	logger.Info("Starting the application....")
	app.Start()

	// reader := bufio.NewReader(os.Stdin)

	// var playerList []players.Player

	// for {
	// 	var player players.Player

	// 	fmt.Print("Enter Name: ")
	// 	player.Name, _ = reader.ReadString('\n')
	// 	player.Name = strings.TrimSpace(player.Name)

	// 	fmt.Print("Enter Age: ")
	// 	fmt.Scanln(&player.Age)

	// 	fmt.Print("What games do you play: ")
	// 	player.Game, _ = reader.ReadString('\n')
	// 	player.Game = strings.TrimSpace(player.Game)

	// 	fmt.Println(player)

	// 	playerList = append(playerList, player)

	// 	fmt.Print("Do you want to enter details for another player? (yes/no): ")
	// 	var input string
	// 	fmt.Scanln(&input)

	// 	if strings.ToLower(input) != "yes" {
	// 		break
	// 	}
	// }

	// // Specify the absolute directory path
	// directoryPath := "/home/gotestserver/golangtest/project1/gotest2023/db"
	// filePath := filepath.Join(directoryPath, "players.csv")

	// // Write playerList to CSV file outside the loop
	// file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()

	// writer := csv.NewWriter(file)
	// defer writer.Flush()

	// // Write header to CSV file
	// if err := writer.Write([]string{"Name", "Age", "Game"}); err != nil {
	// 	panic(err)
	// }
	// // Write player data to CSV file
	// for _, p := range playerList {
	// 	writer.Write([]string{p.Name, fmt.Sprint(p.Age), p.Game})
	// }
}
