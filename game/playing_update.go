package game

import (
	"github.com/aixoio/go-blackjack/game/money"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var MouseIsOverContinuteButton = false

func UpdatePlaying() {
	MouseIsOverContinuteButton = MouseIsOver(400, 588, 192, 32)

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) && MouseIsOverContinuteButton {
		NewHands()
		money.SaveBalenceToFile(DEBUG)
		CURRENT_STATE = BETTING
	}
}
