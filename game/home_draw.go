package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

func DrawHome(screen *ebiten.Image) {
	DrawTextWithPoppinsRegularAt("Blackjack", color.White, 200, 75, screen, 48)
}
