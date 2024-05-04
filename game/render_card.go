package game

import (
	"github.com/aixoio/go-blackjack/blackjack"
	"github.com/aixoio/go-blackjack/game/img"
	"github.com/hajimehoshi/ebiten/v2"
)

func DrawFlipedCard(screen *ebiten.Image, x, y int) {
	DrawImageAt(x, y, screen, img.CARD)
}

func DrawCard(screen *ebiten.Image, x, y int, card blackjack.Card) {

	switch card.Suit {
	case blackjack.SUIT_CLUBS:
		switch card.Face.ExtraFaceDat() {
		case blackjack.ACE_DAT:
			DrawImageAt(x, y, screen, img.CLUBS_ACE)
			return
		case blackjack.JACK_DAT:
			DrawImageAt(x, y, screen, img.CLUBS_JACK)
			return
		case blackjack.KING_DAT:
			DrawImageAt(x, y, screen, img.CLUBS_KING)
			return
		case blackjack.QUEEN_DAT:
			DrawImageAt(x, y, screen, img.CLUBS_QUEEN)
			return
		}
		switch card.Face.Value() {
		case 2:
			DrawImageAt(x, y, screen, img.CLUBS_2)
		case 3:
			DrawImageAt(x, y, screen, img.CLUBS_3)
		case 4:
			DrawImageAt(x, y, screen, img.CLUBS_4)
		case 5:
			DrawImageAt(x, y, screen, img.CLUBS_5)
		case 6:
			DrawImageAt(x, y, screen, img.CLUBS_6)
		case 7:
			DrawImageAt(x, y, screen, img.CLUBS_7)
		case 8:
			DrawImageAt(x, y, screen, img.CLUBS_8)
		case 9:
			DrawImageAt(x, y, screen, img.CLUBS_9)
		case 10:
			DrawImageAt(x, y, screen, img.CLUBS_10)

		}
	case blackjack.SUIT_SPADE:
		switch card.Face.ExtraFaceDat() {
		case blackjack.ACE_DAT:
			DrawImageAt(x, y, screen, img.SPADES_ACE)
			return
		case blackjack.JACK_DAT:
			DrawImageAt(x, y, screen, img.SPADES_JACK)
			return
		case blackjack.KING_DAT:
			DrawImageAt(x, y, screen, img.SPADES_KING)
			return
		case blackjack.QUEEN_DAT:
			DrawImageAt(x, y, screen, img.SPADES_QUEEN)
			return
		}
		switch card.Face.Value() {
		case 2:
			DrawImageAt(x, y, screen, img.SPADES_2)
		case 3:
			DrawImageAt(x, y, screen, img.SPADES_3)
		case 4:
			DrawImageAt(x, y, screen, img.SPADES_4)
		case 5:
			DrawImageAt(x, y, screen, img.SPADES_5)
		case 6:
			DrawImageAt(x, y, screen, img.SPADES_6)
		case 7:
			DrawImageAt(x, y, screen, img.SPADES_7)
		case 8:
			DrawImageAt(x, y, screen, img.SPADES_8)
		case 9:
			DrawImageAt(x, y, screen, img.SPADES_9)
		case 10:
			DrawImageAt(x, y, screen, img.SPADES_10)

		}
	case blackjack.SUIT_DIAMOND:
		switch card.Face.ExtraFaceDat() {
		case blackjack.ACE_DAT:
			DrawImageAt(x, y, screen, img.DIAMONDS_ACE)
			return
		case blackjack.JACK_DAT:
			DrawImageAt(x, y, screen, img.DIAMONDS_JACK)
			return
		case blackjack.KING_DAT:
			DrawImageAt(x, y, screen, img.DIAMONDS_KING)
			return
		case blackjack.QUEEN_DAT:
			DrawImageAt(x, y, screen, img.DIAMONDS_QUEEN)
			return
		}

		switch card.Face.Value() {
		case 2:
			DrawImageAt(x, y, screen, img.DIAMONDS_2)
		case 3:
			DrawImageAt(x, y, screen, img.DIAMONDS_3)
		case 4:
			DrawImageAt(x, y, screen, img.DIAMONDS_4)
		case 5:
			DrawImageAt(x, y, screen, img.DIAMONDS_5)
		case 6:
			DrawImageAt(x, y, screen, img.DIAMONDS_6)
		case 7:
			DrawImageAt(x, y, screen, img.DIAMONDS_7)
		case 8:
			DrawImageAt(x, y, screen, img.DIAMONDS_8)
		case 9:
			DrawImageAt(x, y, screen, img.DIAMONDS_9)
		case 10:
			DrawImageAt(x, y, screen, img.DIAMONDS_10)

		}

	case blackjack.SUIT_HEART:
		switch card.Face.ExtraFaceDat() {
		case blackjack.ACE_DAT:
			DrawImageAt(x, y, screen, img.HEARTS_ACE)
			return
		case blackjack.JACK_DAT:
			DrawImageAt(x, y, screen, img.HEARTS_JACK)
			return
		case blackjack.KING_DAT:
			DrawImageAt(x, y, screen, img.HEARTS_KING)
			return
		case blackjack.QUEEN_DAT:
			DrawImageAt(x, y, screen, img.HEARTS_QUEEN)
			return
		}

		switch card.Face.Value() {
		case 2:
			DrawImageAt(x, y, screen, img.HEARTS_2)
		case 3:
			DrawImageAt(x, y, screen, img.HEARTS_3)
		case 4:
			DrawImageAt(x, y, screen, img.HEARTS_4)
		case 5:
			DrawImageAt(x, y, screen, img.HEARTS_5)
		case 6:
			DrawImageAt(x, y, screen, img.HEARTS_6)
		case 7:
			DrawImageAt(x, y, screen, img.HEARTS_7)
		case 8:
			DrawImageAt(x, y, screen, img.HEARTS_8)
		case 9:
			DrawImageAt(x, y, screen, img.HEARTS_9)
		case 10:
			DrawImageAt(x, y, screen, img.HEARTS_10)

		}

	}

}
