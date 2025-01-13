package entities

import "fmt"

type Token struct {
	Color string
}

func (t Token) New(color string) Token {
	return Token{
		Color: color,
	}
}

func CreateTokens(color string, count int) []Token {
	token := Token{}.New(color)
	tokens := make([]Token, count)
	for i := range tokens {
		tokens[i] = token
	}
	return tokens
}

func (t Token) Print() string {
	colorMap := map[string]string{
		"Red":    "\033[31m",
		"Green":  "\033[32m",
		"Yellow": "\033[33m",
		"Blue":   "\033[34m",
		"Gray":   "\033[90m",
		"Brown":  "\033[38;5;94m",
	}
	colorCode, exists := colorMap[t.Color]
	if exists {
		return fmt.Sprintf("%s%s\033[0m", colorCode, t.Color)
	} else {
		return fmt.Sprintf("%s", t.Color)
	}
}
