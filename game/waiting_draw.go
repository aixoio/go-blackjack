package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func DrawWaiting(screen *ebiten.Image) {
	DrawFlipedCard(screen, 10, 10)
	if DEBUG {
		DrawCard(screen, 10, 10, DealerHand.Cards()[0])
	}
	DrawCard(screen, 50, 10, DealerHand.Cards()[1])
}
