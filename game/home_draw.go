package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

func DrawHome(screen *ebiten.Image) {
	DrawTextWithPoppinsRegularAt("Blackjack", color.White, 200, 75, screen, 48)
	DrawHomePlayButton(screen)
}

// Button location: 268, 300
//
// Button dimenstions: 128x32
func DrawHomePlayButton(screen *ebiten.Image) {
	locationGeoM := ebiten.GeoM{}

	locationGeoM.Translate(268, 300)

	btnImg := ebiten.NewImage(128, 32)
	btnImg.Fill(color.RGBA{R: 51, G: 51, B: 51, A: 255})
	if MouseIsOverHomeButton {
		btnImg.Fill(color.RGBA{R: 150, G: 150, B: 150, A: 255})
		DrawTextWithPoppinsRegularAt("Play", color.RGBA{R: 51, G: 51, B: 51, A: 255}, 48, 4, btnImg, 18)
	} else {
		DrawTextWithPoppinsRegularAt("Play", color.White, 48, 4, btnImg, 18)
	}

	screen.DrawImage(btnImg, &ebiten.DrawImageOptions{
		GeoM: locationGeoM,
	})
}
