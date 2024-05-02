package blackjack

type Hand struct {
	cards []Card
}

func NewHand() Hand {
	return Hand{cards: []Card{}}
}

func (h *Hand) Cards() []Card {
	return h.cards
}

func (h *Hand) Clear() {
	h.cards = []Card{}
}

func (h *Hand) AddCard(c Card) {
	h.cards = append(h.cards, c)
}

func (h *Hand) CountValue() int {
	count := 0

	aces := 0

	for _, card := range h.Cards() {
		if card.Face.extraDat == ACE_DAT {
			aces++
			continue
		}
		count += card.Face.Value()
	}

	for i := 0; i < aces; i++ {
		if count+11 > 21 {
			count++
			continue
		}
		count += 11
	}

	return count
}
