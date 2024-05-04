package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var MouseIsOverHitButton = false
var MouseIsOverStandButton = false

func WaitingUpdate() {
	MouseIsOverHitButton = MouseIsOver(80, 588, 128, 32)
	MouseIsOverStandButton = MouseIsOver(240, 588, 128, 32)

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) && MouseIsOverHitButton {
		HitMe()
	}

}
