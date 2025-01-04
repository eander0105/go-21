package main

import (
	"github.com/eander0105/go-blackjack/blackjack"
)

func main() {
	game := blackjack.NewGame()
	game.AddPlayer("Emil")

	game.Play()
}
