package blackjack

type Card struct {
	Face FaceDat
	Suit Suit
}

type Suit = int

const NUMBER_OF_SUITES int = 4
const NUMBER_OF_CARDS_PER_SUITE int = 13

const (
	SUIT_HEART Suit = iota
	SUIT_DIAMOND
	SUIT_SPADE
	SUIT_CLUBS
)

var SUITES = [4]Suit{SUIT_HEART, SUIT_DIAMOND, SUIT_SPADE, SUIT_CLUBS}

type ExtraFaceDat = int

const (
	ACE_DAT ExtraFaceDat = iota
	JACK_DAT
	QUEEN_DAT
	KING_DAT
	NONE_DAT
)

type FaceDat struct {
	value    int
	extraDat ExtraFaceDat
}

func NewFaceDat(value int, ex ExtraFaceDat) FaceDat {
	return FaceDat{value: value, extraDat: ex}
}

func (f *FaceDat) Value() int {
	return f.value
}

func (f *FaceDat) ExtraFaceDat() ExtraFaceDat {
	return f.extraDat
}

func (c *Card) IsRed() bool {
	return c.Suit == SUIT_DIAMOND || c.Suit == SUIT_HEART
}
