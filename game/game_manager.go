package game

import "github.com/aixoio/go-blackjack/blackjack"

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
	PlayerHand.AddCard(Deck.PopCard())
	PlayerCount = PlayerHand.CountValue()
}
