package game

import (
	"fmt"
	"image/color"

	"github.com/aixoio/go-blackjack/game/money"
	"github.com/hajimehoshi/ebiten/v2"
)

func DrawBetting(screen *ebiten.Image) {

	DrawTextWithPoppinsRegularAt(fmt.Sprintf("You have $%d", money.Balence), color.White, 132, 100, screen, 24)
	DrawTextWithPoppinsRegularAt(fmt.Sprintf("You are betting $%d", money.Betting()), color.White, 132, 150, screen, 24)

	DrawBet10Btn(screen)
	DrawBet100Btn(screen)
	DrawBet1000Btn(screen)
	DrawBet10KBtn(screen)
	DrawBetAllInBtn(screen)
	DrawBetClearBtn(screen)
	DrawPlayBtn(screen)

}

// Button location: 100, 300
//
// Button dimension: 192x32
func DrawBet10Btn(screen *ebiten.Image) {
	locationGeoM := ebiten.GeoM{}
	locationGeoM.Translate(100, 300)
	btnImg := ebiten.NewImage(192, 32)
	btnImg.Fill(color.RGBA{R: 51, G: 51, B: 51, A: 255})
	if MouseIsOverBet10Btn {
		btnImg.Fill(color.RGBA{R: 150, G: 150, B: 150, A: 255})
		DrawTextWithPoppinsRegularAt("Bet $10", color.RGBA{R: 51, G: 51, B: 51, A: 255}, 48, 4, btnImg, 18)
	} else {
		DrawTextWithPoppinsRegularAt("Bet $10", color.White, 48, 4, btnImg, 18)
	}

	screen.DrawImage(btnImg, &ebiten.DrawImageOptions{
		GeoM: locationGeoM,
	})
}

// Button location: 324, 300
//
// Button dimension: 192x32
func DrawBet100Btn(screen *ebiten.Image) {
	locationGeoM := ebiten.GeoM{}
	locationGeoM.Translate(324, 300)
	btnImg := ebiten.NewImage(192, 32)
	btnImg.Fill(color.RGBA{R: 51, G: 51, B: 51, A: 255})
	if MouseIsOverBet100Btn {
		btnImg.Fill(color.RGBA{R: 150, G: 150, B: 150, A: 255})
		DrawTextWithPoppinsRegularAt("Bet $100", color.RGBA{R: 51, G: 51, B: 51, A: 255}, 48, 4, btnImg, 18)
	} else {
		DrawTextWithPoppinsRegularAt("Bet $100", color.White, 48, 4, btnImg, 18)
	}

	screen.DrawImage(btnImg, &ebiten.DrawImageOptions{
		GeoM: locationGeoM,
	})
}

// Button location: 100, 364
//
// Button dimension: 192x32
func DrawBet1000Btn(screen *ebiten.Image) {
	locationGeoM := ebiten.GeoM{}
	locationGeoM.Translate(100, 364)
	btnImg := ebiten.NewImage(192, 32)
	btnImg.Fill(color.RGBA{R: 51, G: 51, B: 51, A: 255})
	if MouseIsOverBet1000Btn {
		btnImg.Fill(color.RGBA{R: 150, G: 150, B: 150, A: 255})
		DrawTextWithPoppinsRegularAt("Bet $1000", color.RGBA{R: 51, G: 51, B: 51, A: 255}, 48, 4, btnImg, 18)
	} else {
		DrawTextWithPoppinsRegularAt("Bet $1000", color.White, 48, 4, btnImg, 18)
	}

	screen.DrawImage(btnImg, &ebiten.DrawImageOptions{
		GeoM: locationGeoM,
	})
}

// Button location: 324, 364
//
// Button dimension: 192x32
func DrawBet10KBtn(screen *ebiten.Image) {
	locationGeoM := ebiten.GeoM{}
	locationGeoM.Translate(324, 364)
	btnImg := ebiten.NewImage(192, 32)
	btnImg.Fill(color.RGBA{R: 51, G: 51, B: 51, A: 255})
	if MouseIsOverBet10KBtn {
		btnImg.Fill(color.RGBA{R: 150, G: 150, B: 150, A: 255})
		DrawTextWithPoppinsRegularAt("Bet $10K", color.RGBA{R: 51, G: 51, B: 51, A: 255}, 48, 4, btnImg, 18)
	} else {
		DrawTextWithPoppinsRegularAt("Bet $10K", color.White, 48, 4, btnImg, 18)
	}

	screen.DrawImage(btnImg, &ebiten.DrawImageOptions{
		GeoM: locationGeoM,
	})
}

// Button location: 100, 428
//
// Button dimension: 192x32
func DrawBetAllInBtn(screen *ebiten.Image) {
	locationGeoM := ebiten.GeoM{}
	locationGeoM.Translate(100, 428)
	btnImg := ebiten.NewImage(192, 32)
	btnImg.Fill(color.RGBA{R: 51, G: 51, B: 51, A: 255})
	if MouseIsOverBetAllInBtn {
		btnImg.Fill(color.RGBA{R: 150, G: 150, B: 150, A: 255})
		DrawTextWithPoppinsRegularAt("All in", color.RGBA{R: 51, G: 51, B: 51, A: 255}, 48, 4, btnImg, 18)
	} else {
		DrawTextWithPoppinsRegularAt("All in", color.White, 48, 4, btnImg, 18)
	}

	screen.DrawImage(btnImg, &ebiten.DrawImageOptions{
		GeoM: locationGeoM,
	})
}

// Button location: 324, 428
//
// Button dimension: 192x32
func DrawBetClearBtn(screen *ebiten.Image) {
	locationGeoM := ebiten.GeoM{}
	locationGeoM.Translate(324, 428)
	btnImg := ebiten.NewImage(192, 32)
	btnImg.Fill(color.RGBA{R: 51, G: 51, B: 51, A: 255})
	if MouseIsOverBetClearBtn {
		btnImg.Fill(color.RGBA{R: 150, G: 150, B: 150, A: 255})
		DrawTextWithPoppinsRegularAt("Clear bet", color.RGBA{R: 51, G: 51, B: 51, A: 255}, 48, 4, btnImg, 18)
	} else {
		DrawTextWithPoppinsRegularAt("Clear bet", color.White, 48, 4, btnImg, 18)
	}

	screen.DrawImage(btnImg, &ebiten.DrawImageOptions{
		GeoM: locationGeoM,
	})
}

// Button location: 100, 492
//
// Button dimension: 416x32
func DrawPlayBtn(screen *ebiten.Image) {
	locationGeoM := ebiten.GeoM{}
	locationGeoM.Translate(100, 492)
	btnImg := ebiten.NewImage(416, 32)
	btnImg.Fill(color.RGBA{R: 51, G: 51, B: 51, A: 255})
	if MouseIsOverPlayBtn {
		btnImg.Fill(color.RGBA{R: 150, G: 150, B: 150, A: 255})
		DrawTextWithPoppinsRegularAt("Play", color.RGBA{R: 51, G: 51, B: 51, A: 255}, 192, 4, btnImg, 18)
	} else {
		DrawTextWithPoppinsRegularAt("Play", color.White, 192, 4, btnImg, 18)
	}

	screen.DrawImage(btnImg, &ebiten.DrawImageOptions{
		GeoM: locationGeoM,
	})
}
