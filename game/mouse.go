package game

import "github.com/hajimehoshi/ebiten/v2"

var mouse MouseDat = MouseDat{X: 0, Y: 0}

type MouseDat struct {
	X int
	Y int
}

func (m *MouseDat) UpdatePosition() {
	mx, my := ebiten.CursorPosition()
	m.X = mx
	m.Y = my
}
