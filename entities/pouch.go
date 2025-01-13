package entities

import (
	"math/rand"
	"time"
)

type Pouch struct {
	Token []Token
}

func (p *Pouch) New() *Pouch {
	var pouch []Token
	pouch = append(pouch, CreateTokens("Blue", 23)...)
	pouch = append(pouch, CreateTokens("Gray", 23)...)
	pouch = append(pouch, CreateTokens("Brown", 21)...)
	pouch = append(pouch, CreateTokens("Green", 19)...)
	pouch = append(pouch, CreateTokens("Yellow", 19)...)
	pouch = append(pouch, CreateTokens("Red", 15)...)
	return &Pouch{
		Token: pouch,
	}
}

func (p *Pouch) TakeTokens() TokenBoard {
	var tokenBoard TokenBoard
	shuffledTokens := shuffleTokens(p.Token)
	tokenBoard.Stack1 = [3]Token{shuffledTokens[0], shuffledTokens[1], shuffledTokens[2]}
	tokenBoard.Stack2 = [3]Token{shuffledTokens[3], shuffledTokens[4], shuffledTokens[5]}
	tokenBoard.Stack3 = [3]Token{shuffledTokens[6], shuffledTokens[7], shuffledTokens[8]}

	p.Token = p.Token[9:]
	return tokenBoard
}

func shuffleTokens(tokens []Token) []Token {
	rand.Seed(time.Now().UnixNano())
	shuffled := make([]Token, len(tokens))
	copy(shuffled, tokens)
	rand.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})
	return shuffled
}
