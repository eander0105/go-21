package main

import (
	game "github.com/eander0105/go-21/game"
)

func main() {
	game := game.NewGame()
	game.AddPlayer("Emil")

	game.Play()
}
