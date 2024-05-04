package game

func (g *Game) Update() error {
	if CURRENT_STATE == HOME {
		UpdateHome()
	}

	return nil
}
