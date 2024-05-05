package game

import (
	"github.com/aixoio/go-blackjack/game/money"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var MouseIsOverBet10Btn = false
var MouseIsOverBet100Btn = false
var MouseIsOverBet1000Btn = false
var MouseIsOverBet10KBtn = false
var MouseIsOverBetAllInBtn = false
var MouseIsOverBetClearBtn = false
var MouseIsOverPlayBtn = false

func UpdateBetting() {
	MouseIsOverBet10Btn = MouseIsOver(100, 300, 192, 32)
	MouseIsOverBet100Btn = MouseIsOver(324, 300, 192, 32)
	MouseIsOverBet1000Btn = MouseIsOver(100, 364, 192, 32)
	MouseIsOverBet10KBtn = MouseIsOver(324, 364, 192, 32)
	MouseIsOverBetAllInBtn = MouseIsOver(100, 428, 192, 32)
	MouseIsOverBetClearBtn = MouseIsOver(324, 428, 192, 32)
	MouseIsOverPlayBtn = MouseIsOver(100, 492, 416, 32)

	if MouseIsOverBet10Btn && inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		money.MakeBet(10)
	}

	if MouseIsOverBet100Btn && inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		money.MakeBet(100)
	}

	if MouseIsOverBet1000Btn && inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		money.MakeBet(1000)
	}

	if MouseIsOverBet10KBtn && inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		money.MakeBet(10000)
	}

	if MouseIsOverBetAllInBtn && inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) && money.Balence > 0 {
		money.MakeBet(money.Balence)
	}

	if MouseIsOverBetClearBtn && inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		money.ClearBet()
	}

	if MouseIsOverPlayBtn && inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		CURRENT_STATE = WAITING
	}

}
