package game

func MouseIsOver(x, y, width, height int) bool {
	return mouse.X >= x && mouse.X <= x+width &&
		mouse.Y >= y && mouse.Y <= y+height
}
