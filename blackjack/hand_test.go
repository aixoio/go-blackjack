package blackjack_test

import (
	"testing"

	"github.com/aixoio/go-blackjack/blackjack"
)

func TestHandCount(t *testing.T) {

	hand := blackjack.NewHand()
	hand.AddCard(blackjack.Card{
		Face: blackjack.NewFaceDat(4, blackjack.NONE_DAT),
		Suit: blackjack.SUIT_CLUBS,
	})
	hand.AddCard(blackjack.Card{
		Face: blackjack.NewFaceDat(3, blackjack.NONE_DAT),
		Suit: blackjack.SUIT_SPADE,
	})

	if hand.CountValue() != 7 {
		t.Fail()
	}

}

func TestHandCountAce(t *testing.T) {

	hand := blackjack.NewHand()
	hand.AddCard(blackjack.Card{
		Face: blackjack.NewFaceDat(1, blackjack.ACE_DAT),
		Suit: blackjack.SUIT_CLUBS,
	})
	hand.AddCard(blackjack.Card{
		Face: blackjack.NewFaceDat(1, blackjack.ACE_DAT),
		Suit: blackjack.SUIT_SPADE,
	})

	if hand.CountValue() != 12 {
		t.Fail()
	}

}
