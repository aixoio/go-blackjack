package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

func DrawWaiting(screen *ebiten.Image) {
	DrawFlipedCard(screen, 10, 10)
	if DEBUG {
		DrawCard(screen, 10, 10, DealerHand.Cards()[0])
	}
	DrawCard(screen, 50, 10, DealerHand.Cards()[1])

	x := 10

	for _, card := range PlayerHand.Cards() {
		DrawCard(screen, x, 300, card)
		x += 40
	}

	DrawTextWithPoppinsRegularAt(fmt.Sprintf("%d", PlayerCount), color.White, 30, 588, screen, 24)
	DrawHitButton(screen)
	DrawStandButton(screen)
}

// Button Location: 80, 588
//
// Button dimensions: 128x32
func DrawHitButton(screen *ebiten.Image) {
	locationGeoM := ebiten.GeoM{}
	locationGeoM.Translate(80, 588)
	btnImg := ebiten.NewImage(128, 32)
	btnImg.Fill(color.RGBA{R: 51, G: 51, B: 51, A: 255})
	if MouseIsOverHitButton {
		btnImg.Fill(color.RGBA{R: 150, G: 150, B: 150, A: 255})
		DrawTextWithPoppinsRegularAt("HIT", color.RGBA{R: 51, G: 51, B: 51, A: 255}, 48, 4, btnImg, 18)
	} else {
		DrawTextWithPoppinsRegularAt("HIT", color.White, 48, 4, btnImg, 18)
	}

	screen.DrawImage(btnImg, &ebiten.DrawImageOptions{
		GeoM: locationGeoM,
	})
}

// Button Location: 240, 588
//
// Button dimensions: 128x32
func DrawStandButton(screen *ebiten.Image) {
	locationGeoM := ebiten.GeoM{}
	locationGeoM.Translate(240, 588)
	btnImg := ebiten.NewImage(128, 32)
	btnImg.Fill(color.RGBA{R: 51, G: 51, B: 51, A: 255})
	if MouseIsOverStandButton {
		btnImg.Fill(color.RGBA{R: 150, G: 150, B: 150, A: 255})
		DrawTextWithPoppinsRegularAt("STAND", color.RGBA{R: 51, G: 51, B: 51, A: 255}, 32, 4, btnImg, 18)
	} else {
		DrawTextWithPoppinsRegularAt("STAND", color.White, 32, 4, btnImg, 18)
	}

	screen.DrawImage(btnImg, &ebiten.DrawImageOptions{
		GeoM: locationGeoM,
	})
}
