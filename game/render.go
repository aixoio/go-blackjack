package game

import (
	"image/color"

	"github.com/aixoio/go-blackjack/blackjack"
	"github.com/aixoio/go-blackjack/game/img"
	"github.com/hajimehoshi/ebiten/v2"
)

func DrawBGColor(screen *ebiten.Image) {
	bg := ebiten.NewImage(WIDTH, HEIGHT)
	bg.Fill(color.RGBA{R: 0, G: 110, B: 4, A: 0})
	screen.DrawImage(bg, nil)
}

func DrawFlipedCard(screen *ebiten.Image, x, y int) {
	DrawImageAt(x, y, screen, img.CARD)
}

func DrawCard(screen *ebiten.Image, x, y int, card blackjack.Card) {

	switch card.Suit {
	case blackjack.SUIT_CLUBS:
		switch card.Face.ExtraFaceDat() {
		case blackjack.ACE_DAT:
			DrawImageAt(x, y, screen, img.CLUBS_ACE)
		case blackjack.JACK_DAT:
			DrawImageAt(x, y, screen, img.CLUBS_JACK)
		case blackjack.KING_DAT:
			DrawImageAt(x, y, screen, img.CLUBS_KING)
		case blackjack.QUEEN_DAT:
			DrawImageAt(x, y, screen, img.CLUBS_QUEEN)
		}
	case blackjack.SUIT_SPADE:
		switch card.Face.ExtraFaceDat() {
		case blackjack.ACE_DAT:
			DrawImageAt(x, y, screen, img.SPADES_ACE)
		case blackjack.JACK_DAT:
			DrawImageAt(x, y, screen, img.SPADES_JACK)
		case blackjack.KING_DAT:
			DrawImageAt(x, y, screen, img.SPADES_KING)
		case blackjack.QUEEN_DAT:
			DrawImageAt(x, y, screen, img.SPADES_QUEEN)
		}
		return
	case blackjack.SUIT_DIAMOND:
		switch card.Face.ExtraFaceDat() {
		case blackjack.ACE_DAT:
			DrawImageAt(x, y, screen, img.DIAMONDS_ACE)
		case blackjack.JACK_DAT:
			DrawImageAt(x, y, screen, img.DIAMONDS_JACK)
		case blackjack.KING_DAT:
			DrawImageAt(x, y, screen, img.DIAMONDS_KING)
		case blackjack.QUEEN_DAT:
			DrawImageAt(x, y, screen, img.DIAMONDS_QUEEN)
		}
		return
	case blackjack.SUIT_HEART:
		switch card.Face.ExtraFaceDat() {
		case blackjack.ACE_DAT:
			DrawImageAt(x, y, screen, img.HEARTS_ACE)
		case blackjack.JACK_DAT:
			DrawImageAt(x, y, screen, img.HEARTS_JACK)
		case blackjack.KING_DAT:
			DrawImageAt(x, y, screen, img.HEARTS_KING)
		case blackjack.QUEEN_DAT:
			DrawImageAt(x, y, screen, img.HEARTS_QUEEN)
		}
		return
	}

}
