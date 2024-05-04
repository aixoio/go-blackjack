package img

import (
	"bytes"
	"image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

func loadDiamonds() error {
	var err error

	cardDecoded, err := png.Decode(bytes.NewBuffer(diamonds_png_ACE))
	if err != nil {
		return err
	}
	DIAMONDS_ACE = ebiten.NewImageFromImage(cardDecoded)

	cardDecoded, err = png.Decode(bytes.NewBuffer(diamonds_png_2))
	if err != nil {
		return err
	}
	DIAMONDS_2 = ebiten.NewImageFromImage(cardDecoded)

	cardDecoded, err = png.Decode(bytes.NewBuffer(diamonds_png_3))
	if err != nil {
		return err
	}
	DIAMONDS_3 = ebiten.NewImageFromImage(cardDecoded)

	cardDecoded, err = png.Decode(bytes.NewBuffer(diamonds_png_4))
	if err != nil {
		return err
	}
	DIAMONDS_4 = ebiten.NewImageFromImage(cardDecoded)

	cardDecoded, err = png.Decode(bytes.NewBuffer(diamonds_png_5))
	if err != nil {
		return err
	}
	DIAMONDS_5 = ebiten.NewImageFromImage(cardDecoded)

	cardDecoded, err = png.Decode(bytes.NewBuffer(diamonds_png_6))
	if err != nil {
		return err
	}
	DIAMONDS_6 = ebiten.NewImageFromImage(cardDecoded)

	cardDecoded, err = png.Decode(bytes.NewBuffer(diamonds_png_7))
	if err != nil {
		return err
	}
	DIAMONDS_7 = ebiten.NewImageFromImage(cardDecoded)

	cardDecoded, err = png.Decode(bytes.NewBuffer(diamonds_png_8))
	if err != nil {
		return err
	}
	DIAMONDS_8 = ebiten.NewImageFromImage(cardDecoded)

	cardDecoded, err = png.Decode(bytes.NewBuffer(diamonds_png_9))
	if err != nil {
		return err
	}
	DIAMONDS_9 = ebiten.NewImageFromImage(cardDecoded)

	cardDecoded, err = png.Decode(bytes.NewBuffer(diamonds_png_10))
	if err != nil {
		return err
	}
	DIAMONDS_10 = ebiten.NewImageFromImage(cardDecoded)

	cardDecoded, err = png.Decode(bytes.NewBuffer(diamonds_png_JACK))
	if err != nil {
		return err
	}
	DIAMONDS_JACK = ebiten.NewImageFromImage(cardDecoded)

	cardDecoded, err = png.Decode(bytes.NewBuffer(diamonds_png_KING))
	if err != nil {
		return err
	}
	DIAMONDS_KING = ebiten.NewImageFromImage(cardDecoded)

	cardDecoded, err = png.Decode(bytes.NewBuffer(diamonds_png_QUEEN))
	if err != nil {
		return err
	}
	DIAMONDS_QUEEN = ebiten.NewImageFromImage(cardDecoded)

	return err
}
