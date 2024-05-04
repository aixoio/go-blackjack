package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

func DrawBGColor(screen *ebiten.Image) {
	bg := ebiten.NewImage(WIDTH, HEIGHT)
	bg.Fill(color.RGBA{R: 0, G: 110, B: 4, A: 0})
	screen.DrawImage(bg, nil)
}
