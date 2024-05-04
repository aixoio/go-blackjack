package game

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func Start(debug bool) {
	DEBUG = debug
	ebiten.SetWindowSize(WIDTH, HEIGHT)
	ebiten.SetWindowTitle("Blackjack")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
