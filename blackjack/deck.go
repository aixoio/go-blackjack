package blackjack

type Deck struct {
	cards []Card
}

func (d *Deck) Cards() []Card {
	return d.cards
}

func NewDeck() Deck {
	deck := Deck{cards: []Card{}}

	for i := 0; i < NUMBER_OF_SUITES; i++ {
		for j := 0; j < NUMBER_OF_CARDS_PER_SUITE; j++ {
			card := Card{
				Face: NewFaceDat(j+1, NONE_DAT),
				Suit: SUITES[i],
			}

			switch j {
			case 9: // this a jack
				card.Face.extraDat = JACK_DAT
			case 10: // this a queen
				card.Face.extraDat = QUEEN_DAT
			case 11: // this a king
				card.Face.extraDat = KING_DAT
			case 12: // this a ace
				card.Face.extraDat = ACE_DAT

			}

			deck.cards = append(deck.cards, card)
		}
	}

	return deck

}
