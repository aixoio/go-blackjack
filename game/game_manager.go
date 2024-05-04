package game

import (
	"github.com/aixoio/go-blackjack/blackjack"
)

var Deck blackjack.Deck
var DealerHand, PlayerHand blackjack.Hand
var PlayerCount, DealerCount int

func InitBlackjack() {
	Deck = blackjack.NewDeck()
	DealerHand = blackjack.NewHand()
	PlayerHand = blackjack.NewHand()
	Deck.Shuffle()

	DealerHand.AddCard(Deck.PopCard())
	DealerHand.AddCard(Deck.PopCard())

	PlayerHand.AddCard(Deck.PopCard())
	PlayerHand.AddCard(Deck.PopCard())

	DealerCount = DealerHand.CountValue()
	PlayerCount = PlayerHand.CountValue()

}

func HitMe() {
	if Deck.TotalCards()-1 < 0 {
		InitBlackjack()
	}
	PlayerHand.AddCard(Deck.PopCard())
	PlayerCount = PlayerHand.CountValue()
}

func DealerPlay() {
	for DealerCount < PlayerCount && DealerCount < 20 {
		DealerHand.AddCard(Deck.PopCard())
		DealerCount = DealerHand.CountValue()
	}
}

func NewHands() {

	if Deck.TotalCards()-4 < 0 {
		InitBlackjack()
		return
	}

	DealerHand = blackjack.NewHand()
	PlayerHand = blackjack.NewHand()

	DealerHand.AddCard(Deck.PopCard())
	DealerHand.AddCard(Deck.PopCard())

	PlayerHand.AddCard(Deck.PopCard())
	PlayerHand.AddCard(Deck.PopCard())

	DealerCount = DealerHand.CountValue()
	PlayerCount = PlayerHand.CountValue()
}
