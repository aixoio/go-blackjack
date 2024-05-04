package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Clear()
	DrawBGColor(screen)

	if CURRENT_STATE == HOME {
		DrawHome(screen)
	}

	if DEBUG {
		DebugDraw(screen)
	}
}
