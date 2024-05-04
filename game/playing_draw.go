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
	DrawTextWithPoppinsRegularAt("Click anything to contine", color.White, 260, 588, screen, 24)

	DrawPlayingInfo(screen)
}

func DrawPlayingInfo(screen *ebiten.Image) {

	if DealerCount > 21 && PlayerCount > 21 {
		DrawTextWithPoppinsRegularAt("Push", color.White, 260, 300, screen, 48)
	} else if DealerCount > 21 {
		DrawTextWithPoppinsRegularAt("Player wins", color.White, 260, 300, screen, 48)
	} else if PlayerCount > 21 {
		DrawTextWithPoppinsRegularAt("Dealer wins", color.White, 260, 300, screen, 48)
	} else if DealerCount == PlayerCount {
		DrawTextWithPoppinsRegularAt("Push", color.White, 260, 300, screen, 48)
	} else if DealerCount > PlayerCount {
		DrawTextWithPoppinsRegularAt("Dealer wins", color.White, 260, 300, screen, 48)
	} else if DealerCount < PlayerCount {
		DrawTextWithPoppinsRegularAt("Player wins", color.White, 260, 300, screen, 48)
	}

}
