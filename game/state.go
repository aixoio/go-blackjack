package game

type GameState = int

const (
	HOME GameState = iota
	WAITING
	PLAYING
)

var CURRENT_STATE GameState = HOME
