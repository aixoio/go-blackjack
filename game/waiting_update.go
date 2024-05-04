package game

import "github.com/hajimehoshi/ebiten/v2"

func WaitingUpdate() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		InitBlackjack()
	}
}
