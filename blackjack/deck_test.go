package blackjack_test

import (
	"testing"

	"github.com/aixoio/go-blackjack/blackjack"
)

func TestNewDeck(t *testing.T) {
	deck := blackjack.NewDeck()

	if len(deck.Cards()) != 52 {
		t.Fail()
	}

}
