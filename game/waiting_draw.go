package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

func DrawWaiting(screen *ebiten.Image) {
	if !DEBUG {
		DrawFlipedCard(screen, 10, 10)
	} else {
		DrawCard(screen, 10, 10, DealerHand.Cards()[0])
		fmt.Println(DealerHand.Cards()[0])
	}
	DrawCard(screen, 50, 10, DealerHand.Cards()[1])
}
