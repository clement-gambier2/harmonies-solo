package entities

import "fmt"

type Tile struct {
	Stack [3]Token
}

func (t *Tile) Place(token Token) {
	if t.Stack[0] == (Token{}) {
		t.Stack[0] = token
	} else if t.Stack[1] == (Token{}) {
		t.Stack[1] = token
	} else {
		t.Stack[2] = token
	}
}

func (t *Tile) Print() {
	for i := 0; i < 3; i++ {
		fmt.Print(t.Stack[i].Print(), " ")
		fmt.Print(t.Stack[i].Print(), " ")
		fmt.Print(t.Stack[i].Print(), " ")
	}
}
