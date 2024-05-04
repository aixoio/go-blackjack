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

	if CURRENT_STATE == WAITING {
		DrawWaiting(screen)
	}

	if CURRENT_STATE == PLAYING {
		DrawPlaying(screen)
	}

	if DEBUG {
		DebugDraw(screen)
	}
}
