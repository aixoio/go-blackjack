package img

import (
	"bytes"
	"image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

func loadHearts() error {
	var err error

	cardDecoded, err := png.Decode(bytes.NewBuffer(hearts_png_ACE))
	if err != nil {
		return err
	}
	HEARTS_ACE = ebiten.NewImageFromImage(cardDecoded)

	cardDecoded, err = png.Decode(bytes.NewBuffer(hearts_png_2))
	if err != nil {
		return err
	}
	HEARTS_2 = ebiten.NewImageFromImage(cardDecoded)

	cardDecoded, err = png.Decode(bytes.NewBuffer(hearts_png_3))
	if err != nil {
		return err
	}
	HEARTS_3 = ebiten.NewImageFromImage(cardDecoded)

	cardDecoded, err = png.Decode(bytes.NewBuffer(hearts_png_4))
	if err != nil {
		return err
	}
	HEARTS_4 = ebiten.NewImageFromImage(cardDecoded)

	cardDecoded, err = png.Decode(bytes.NewBuffer(hearts_png_5))
	if err != nil {
		return err
	}
	HEARTS_5 = ebiten.NewImageFromImage(cardDecoded)

	cardDecoded, err = png.Decode(bytes.NewBuffer(hearts_png_6))
	if err != nil {
		return err
	}
	HEARTS_6 = ebiten.NewImageFromImage(cardDecoded)

	cardDecoded, err = png.Decode(bytes.NewBuffer(hearts_png_7))
	if err != nil {
		return err
	}
	HEARTS_7 = ebiten.NewImageFromImage(cardDecoded)

	cardDecoded, err = png.Decode(bytes.NewBuffer(hearts_png_8))
	if err != nil {
		return err
	}
	HEARTS_8 = ebiten.NewImageFromImage(cardDecoded)

	cardDecoded, err = png.Decode(bytes.NewBuffer(hearts_png_9))
	if err != nil {
		return err
	}
	HEARTS_9 = ebiten.NewImageFromImage(cardDecoded)

	cardDecoded, err = png.Decode(bytes.NewBuffer(hearts_png_10))
	if err != nil {
		return err
	}
	HEARTS_10 = ebiten.NewImageFromImage(cardDecoded)

	cardDecoded, err = png.Decode(bytes.NewBuffer(hearts_png_JACK))
	if err != nil {
		return err
	}
	HEARTS_JACK = ebiten.NewImageFromImage(cardDecoded)

	cardDecoded, err = png.Decode(bytes.NewBuffer(hearts_png_KING))
	if err != nil {
		return err
	}
	HEARTS_KING = ebiten.NewImageFromImage(cardDecoded)

	cardDecoded, err = png.Decode(bytes.NewBuffer(hearts_png_QUEEN))
	if err != nil {
		return err
	}
	HEARTS_QUEEN = ebiten.NewImageFromImage(cardDecoded)

	return err
}
