package img

import (
	"bytes"
	"image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

func loadSpades() error {
	var err error

	cardDecoded, err := png.Decode(bytes.NewBuffer(spades_png_ACE))
	if err != nil {
		return err
	}
	SPADES_ACE = ebiten.NewImageFromImage(cardDecoded)

	cardDecoded, err = png.Decode(bytes.NewBuffer(spades_png_2))
	if err != nil {
		return err
	}
	SPADES_2 = ebiten.NewImageFromImage(cardDecoded)

	cardDecoded, err = png.Decode(bytes.NewBuffer(spades_png_3))
	if err != nil {
		return err
	}
	SPADES_3 = ebiten.NewImageFromImage(cardDecoded)

	cardDecoded, err = png.Decode(bytes.NewBuffer(spades_png_4))
	if err != nil {
		return err
	}
	SPADES_4 = ebiten.NewImageFromImage(cardDecoded)

	cardDecoded, err = png.Decode(bytes.NewBuffer(spades_png_5))
	if err != nil {
		return err
	}
	SPADES_5 = ebiten.NewImageFromImage(cardDecoded)

	cardDecoded, err = png.Decode(bytes.NewBuffer(spades_png_6))
	if err != nil {
		return err
	}
	SPADES_6 = ebiten.NewImageFromImage(cardDecoded)

	cardDecoded, err = png.Decode(bytes.NewBuffer(spades_png_7))
	if err != nil {
		return err
	}
	SPADES_7 = ebiten.NewImageFromImage(cardDecoded)

	cardDecoded, err = png.Decode(bytes.NewBuffer(spades_png_8))
	if err != nil {
		return err
	}
	SPADES_8 = ebiten.NewImageFromImage(cardDecoded)

	cardDecoded, err = png.Decode(bytes.NewBuffer(spades_png_9))
	if err != nil {
		return err
	}
	SPADES_9 = ebiten.NewImageFromImage(cardDecoded)

	cardDecoded, err = png.Decode(bytes.NewBuffer(spades_png_10))
	if err != nil {
		return err
	}
	SPADES_10 = ebiten.NewImageFromImage(cardDecoded)

	cardDecoded, err = png.Decode(bytes.NewBuffer(spades_png_JACK))
	if err != nil {
		return err
	}
	SPADES_JACK = ebiten.NewImageFromImage(cardDecoded)

	cardDecoded, err = png.Decode(bytes.NewBuffer(spades_png_KING))
	if err != nil {
		return err
	}
	SPADES_KING = ebiten.NewImageFromImage(cardDecoded)

	cardDecoded, err = png.Decode(bytes.NewBuffer(spades_png_QUEEN))
	if err != nil {
		return err
	}
	SPADES_QUEEN = ebiten.NewImageFromImage(cardDecoded)

	return err
}
