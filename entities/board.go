package entities

import "fmt"

type Board struct {
	Board [][]Tile
}

func (b *Board) New() Board {
	return Board{Board: [][]Tile{
		{Tile{}, Tile{}, Tile{}, Tile{}, Tile{}},
		{Tile{}, Tile{}, Tile{}, Tile{}},
		{Tile{}, Tile{}, Tile{}, Tile{}, Tile{}},
		{Tile{}, Tile{}, Tile{}, Tile{}},
		{Tile{}, Tile{}, Tile{}, Tile{}, Tile{}},
	}}
}

func (b *Board) PlaceAToken(t Token, nbColumn int, nbTile int) {
	b.Board[nbColumn][nbTile].Place(t)
}

func (b *Board) Print() {
	for line := 0; line < 5; line++ {
		for col := 0; col < 5; col++ {
			fmt.Print("|")
			for stack := 0; stack < 3; stack++ {
				if (col == 1 && line == 4) || (col == 3 && line == 4) {
					fmt.Print(" ")
				} else {
					fmt.Print(b.Board[col][line].Stack[stack].Print())
				}
			}
		}
		fmt.Println("|")
	}
}
