package game

var MouseIsOverHitButton = false
var MouseIsOverStandButton = false

func WaitingUpdate() {
	MouseIsOverHitButton = MouseIsOver(80, 588, 128, 32)
	MouseIsOverStandButton = MouseIsOver(240, 588, 128, 32)
}
