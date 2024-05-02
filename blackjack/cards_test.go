package blackjack_test

import (
	"testing"

	"github.com/aixoio/go-blackjack/blackjack"
)

func TestCardIsRed(t *testing.T) {
	red := blackjack.Card{
		Face: blackjack.NewFaceDat(1),
		Suit: blackjack.SUIT_HEART,
	}

	if !red.IsRed() {
		t.Fail()
	}
}
