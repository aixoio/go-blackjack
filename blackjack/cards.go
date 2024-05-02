package blackjack

type Card struct {
	Face FaceDat
	Suit Suit
}

type Suit = int

const (
	SUIT_HEART Suit = iota
	SUIT_DIAMOND
	SUIT_SPADE
	SUIT_CLUBS
)

type FaceDat struct {
	value int
}

func NewFaceDat(value int) FaceDat {
	return FaceDat{value: value}
}

func (f *FaceDat) Value() int {
	return f.value
}

func (c *Card) IsRed() bool {
	return c.Suit == SUIT_DIAMOND || c.Suit == SUIT_HEART
}
