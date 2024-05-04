package game

var MouseIsOverHomeButton = false

func UpdateHome() {
	MouseIsOverHomeButton = MouseIsOver(268, 300, 128, 32)
}
