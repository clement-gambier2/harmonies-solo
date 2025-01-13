package entities

import (
	"fmt"
)

type Board struct {
	Column1 [5]Tile
	Column2 [4]Tile
	Column3 [5]Tile
	Column4 [4]Tile
	Column5 [5]Tile
}

func (b *Board) New() *Board {
	return &Board{}
}

func (b *Board) PlaceAToken(t Token, nbColumn int, nbTile int) {
	if nbColumn == 0 {
		b.Column1[nbTile].Place(t)
	} else if nbColumn == 1 {
		b.Column2[nbTile].Place(t)
	} else if nbColumn == 2 {
		b.Column3[nbTile].Place(t)
	} else if nbColumn == 3 {
		b.Column4[nbTile].Place(t)
	} else {
		b.Column5[nbTile].Place(t)
	}
}

func (b *Board) Print() {
	for i := 0; i < 5; i++ {
		fmt.Print(b.Column1[i].Stack)
		if i != 4 {
			fmt.Print(b.Column2[i].Stack)
		} else {
			fmt.Print("vide")
		}
		fmt.Print(b.Column3[i].Stack)
		if i != 4 {
			fmt.Print(b.Column4[i].Stack)
		} else {
			fmt.Print("vide")
		}
		fmt.Print(b.Column5[i].Stack)
		fmt.Println()
	}
}
