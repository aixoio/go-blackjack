package game

import (
	"image/color"

	"github.com/aixoio/go-blackjack/game/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

func DrawImageAt(x, y int, screen, img *ebiten.Image) {
	locationGeoM := ebiten.GeoM{}
	locationGeoM.Translate(float64(x), float64(y))

	screen.DrawImage(img, &ebiten.DrawImageOptions{
		GeoM: locationGeoM,
	})
}

func DrawTextWithPoppinsRegularAt(str string, color color.Color, x, y int, screen *ebiten.Image, size float64) {
	drawOptions := &text.DrawOptions{}

	drawOptions.GeoM.Translate(float64(x), float64(y))

	text.Draw(screen, str, &text.GoTextFace{
		Source: fonts.PoppinsRegular,
		Size:   size,
	}, drawOptions)
}
