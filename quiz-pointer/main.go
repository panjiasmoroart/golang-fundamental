package main

import "fmt"

type Gamer struct {
	Name  string
	Games []string
}

// method AddGame
func (gamer *Gamer) AddGame(value string) {
	gamer.Games = append(gamer.Games, value)
}

func main() {
	gamer := Gamer{Name: "Panji"}

	gamer.AddGame("Mario")
	gamer.AddGame("Mortal Combat")
	gamer.AddGame("Super Contra")

	for _, game := range gamer.Games {
		fmt.Println(game)
	}
}
