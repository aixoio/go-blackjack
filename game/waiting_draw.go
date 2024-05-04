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
}
