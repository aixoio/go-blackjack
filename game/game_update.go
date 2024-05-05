package game

func (g *Game) Update() error {
	mouse.UpdatePosition()

	if CURRENT_STATE == HOME {
		UpdateHome()
	}

	if CURRENT_STATE == WAITING {
		WaitingUpdate()
	}

	if CURRENT_STATE == PLAYING {
		UpdatePlaying()
	}

	if CURRENT_STATE == BETTING {
		UpdateBetting()
	}

	return nil
}
