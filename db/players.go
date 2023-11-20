// players/player.go
package players

type Player struct {
	Name string `json:"full_name" xml:"name"`
	Age  int    `json:"age" xml:"age"`
	Game string `json:"game" xnk:"game"`
}
