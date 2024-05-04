package game

import (
	"github.com/aixoio/go-blackjack/blackjack"
	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Clear()
	DrawBGColor(screen)
	DrawFlipedCard(screen, 10, 10)
	DrawCard(screen, 50, 10, blackjack.Card{Face: blackjack.NewFaceDat(10, blackjack.JACK_DAT), Suit: blackjack.SUIT_CLUBS}) // 40 px diff
	DrawCard(screen, 100, 10, blackjack.Card{Face: blackjack.NewFaceDat(10, blackjack.KING_DAT), Suit: blackjack.SUIT_HEART})
	DrawCard(screen, 100, 10, blackjack.Card{Face: blackjack.NewFaceDat(7, blackjack.NONE_DAT), Suit: blackjack.SUIT_HEART})

	if DEBUG {
		DebugDraw(screen)
	}
}
