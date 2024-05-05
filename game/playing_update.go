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

		if DealerCount > 21 && PlayerCount > 21 {
			money.Push()
		} else if DealerCount > 21 {
			money.WonBet()
		} else if PlayerCount > 21 {
			money.LoseBet()
		} else if DealerCount == PlayerCount {
			money.Push()
		} else if DealerCount > PlayerCount {
			money.LoseBet()
		} else if DealerCount < PlayerCount {
			money.WonBet()
		}

		money.SaveBalenceToFile(DEBUG)

		NewHands()
		CURRENT_STATE = BETTING
	}
}
