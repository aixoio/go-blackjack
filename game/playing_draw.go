package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

func DrawPlaying(screen *ebiten.Image) {
	x := 10

	for _, card := range DealerHand.Cards() {
		DrawCard(screen, x, 10, card)
		x += 40
	}

	x = 10

	for _, card := range PlayerHand.Cards() {
		DrawCard(screen, x, 300, card)
		x += 40
	}

	DrawTextWithPoppinsRegularAt(fmt.Sprintf("%d", DealerCount), color.White, 30, 265, screen, 24)
	DrawTextWithPoppinsRegularAt(fmt.Sprintf("%d", PlayerCount), color.White, 30, 588, screen, 24)

	DrawPlayAgainBtn(screen)

	DrawPlayingInfo(screen)
}

func DrawPlayingInfo(screen *ebiten.Image) {

	if DealerCount > 21 && PlayerCount > 21 {
		DrawTextWithPoppinsRegularAt("Push", color.White, 260, 250, screen, 48)
	} else if DealerCount > 21 {
		DrawTextWithPoppinsRegularAt("Player wins", color.White, 260, 250, screen, 48)
	} else if PlayerCount > 21 {
		DrawTextWithPoppinsRegularAt("Dealer wins", color.White, 260, 250, screen, 48)
	} else if DealerCount == PlayerCount {
		DrawTextWithPoppinsRegularAt("Push", color.White, 260, 250, screen, 48)
	} else if DealerCount > PlayerCount {
		DrawTextWithPoppinsRegularAt("Dealer wins", color.White, 260, 250, screen, 48)
	} else if DealerCount < PlayerCount {
		DrawTextWithPoppinsRegularAt("Player wins", color.White, 260, 250, screen, 48)
	}

}

// Buttion locaiton: 400, 588
//
// Button dimension: 192x32
func DrawPlayAgainBtn(screen *ebiten.Image) {
	locationGeoM := ebiten.GeoM{}
	locationGeoM.Translate(400, 588)
	btnImg := ebiten.NewImage(192, 32)
	btnImg.Fill(color.RGBA{R: 51, G: 51, B: 51, A: 255})
	if MouseIsOverContinuteButton {
		btnImg.Fill(color.RGBA{R: 150, G: 150, B: 150, A: 255})
		DrawTextWithPoppinsRegularAt("Play again", color.RGBA{R: 51, G: 51, B: 51, A: 255}, 48, 4, btnImg, 18)
	} else {
		DrawTextWithPoppinsRegularAt("Play again", color.White, 48, 4, btnImg, 18)
	}

	screen.DrawImage(btnImg, &ebiten.DrawImageOptions{
		GeoM: locationGeoM,
	})
}
