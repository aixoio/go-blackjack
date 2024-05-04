package game

import "github.com/aixoio/go-blackjack/blackjack"

var Deck blackjack.Deck
var DealerHand, PlayerHand blackjack.Hand

func InitBlackjack() {
	Deck = blackjack.NewDeck()
	DealerHand = blackjack.NewHand()
	PlayerHand = blackjack.NewHand()
	Deck.Shuffle()

	DealerHand.AddCard(Deck.PopCard())
	DealerHand.AddCard(Deck.PopCard())

	PlayerHand.AddCard(Deck.PopCard())
	PlayerHand.AddCard(Deck.PopCard())

}
