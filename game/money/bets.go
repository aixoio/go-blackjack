package money

var betting int64 = 0

func Betting() int64 {
	return betting
}

func MakeBet(a int64) {
	betting += a
	Balence -= a
}

func WonBet() {
	Balence += betting * 2
	betting = 0
}

func LoseBet() {
	betting = 0
}

func Push() {
	Balence += betting
	betting = 0
}

func ClearBet() {
	Balence += betting
	betting = 0
}
