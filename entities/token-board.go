package entities

import "fmt"

type TokenBoard struct {
	Stack1 [3]Token
	Stack2 [3]Token
	Stack3 [3]Token
}

func (b TokenBoard) Print() {
	for i := 0; i < 3; i++ {
		fmt.Print(b.Stack1[i].Print(), " ")
		fmt.Print(b.Stack2[i].Print(), " ")
		fmt.Print(b.Stack3[i].Print(), " ")
		fmt.Println()
	}

}
