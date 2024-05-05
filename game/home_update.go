package game

import (
	"github.com/aixoio/go-blackjack/game/money"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var MouseIsOverHomeButton = false

func UpdateHome() {
	MouseIsOverHomeButton = MouseIsOver(268, 300, 128, 32)

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) && MouseIsOverHomeButton {
		InitBlackjack()
		money.SaveBalenceToFile(DEBUG)
		CURRENT_STATE = BETTING
	}

}
