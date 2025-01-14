package main

import (
	"fmt"
	"github.com/harmonies/entities"
)

func main() {
	fmt.Println("Harmonies Solo Mode")
	e := entities.Pouch{}
	pouch := e.New()
	e2 := entities.Board{}
	board := e2.New()
	game := entities.Game{
		Pouch:      pouch,
		TokenBoard: pouch.TakeTokens(),
		Board:      board,
	}
	game.TokenBoard.Print()
	fmt.Println(len(game.Pouch.Token))

	stackChosen := game.TokenBoard.Stack2

	game.Board.PlaceAToken(stackChosen[0], 0, 0)
	game.Board.PlaceAToken(stackChosen[1], 0, 0)
	game.Board.PlaceAToken(stackChosen[2], 1, 0)
	game.Board.Print()
}
