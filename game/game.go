package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const WIDTH, HEIGHT = 512, 512

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Clear()
	DrawBGColor(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return WIDTH, HEIGHT
}
