package game

import "github.com/hajimehoshi/ebiten/v2"

var MouseIsOverHomeButton = false

func UpdateHome() {
	MouseIsOverHomeButton = MouseIsOver(268, 300, 128, 32)

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && MouseIsOverHomeButton {
		CURRENT_STATE = WAITING
	}

}
