package dto

// Player response represents a player in the system.
type PlayerResponse struct {
	PlayerId string `json:"player_id"`
	Name     string `json:"full_name"`
	Age      int    `json:"age" `
	Game     string `json:"game"`
	Status   string `json:"status" `
}
