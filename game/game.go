package game

const WIDTH, HEIGHT = 640, 640

var DEBUG = false

type Game struct{}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return WIDTH, HEIGHT
}
