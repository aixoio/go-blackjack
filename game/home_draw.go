package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

func DrawHome(screen *ebiten.Image) {
	DrawTextWithPoppinsRegularAt("Blackjack", color.White, 10, 10, screen, 48)
}
