package game

func (g *Game) Update() error {
	mouse.UpdatePosition()

	if CURRENT_STATE == HOME {
		UpdateHome()
	}

	return nil
}
