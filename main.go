package main

import (
	"fmt"
	"github.com/harmonies/entities"
)

func main() {
	fmt.Println("Harmonies Solo Mode")
	e := entities.Pouch{}
	pouch := e.New()
	game := entities.Game{
		Pouch:      pouch,
		TokenBoard: pouch.TakeTokens(),
	}
	game.TokenBoard.Print()
	fmt.Println(len(game.Pouch.Token))

	game.TokenBoard = game.Pouch.TakeTokens()
	game.TokenBoard.Print()
	fmt.Println(len(game.Pouch.Token))

}
