package main

import (
	"os"

	"github.com/aixoio/go-blackjack/game"
)

func main() {
	game.Start(len(os.Args) >= 2)
}
