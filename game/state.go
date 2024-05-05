package game

type GameState = int

const (
	HOME GameState = iota
	WAITING
	PLAYING
	BETTING
)

var CURRENT_STATE GameState = HOME
