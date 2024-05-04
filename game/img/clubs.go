package img

import (
	"bytes"
	"image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

func loadClubs() error {
	var err error

	cardDecoded, err := png.Decode(bytes.NewBuffer(clubs_png_ACE))
	if err != nil {
		return err
	}
	CLUBS_ACE = ebiten.NewImageFromImage(cardDecoded)

	cardDecoded, err = png.Decode(bytes.NewBuffer(clubs_png_2))
	if err != nil {
		return err
	}
	CLUBS_2 = ebiten.NewImageFromImage(cardDecoded)

	cardDecoded, err = png.Decode(bytes.NewBuffer(clubs_png_3))
	if err != nil {
		return err
	}
	CLUBS_3 = ebiten.NewImageFromImage(cardDecoded)

	cardDecoded, err = png.Decode(bytes.NewBuffer(clubs_png_4))
	if err != nil {
		return err
	}
	CLUBS_4 = ebiten.NewImageFromImage(cardDecoded)

	cardDecoded, err = png.Decode(bytes.NewBuffer(clubs_png_5))
	if err != nil {
		return err
	}
	CLUBS_5 = ebiten.NewImageFromImage(cardDecoded)

	cardDecoded, err = png.Decode(bytes.NewBuffer(clubs_png_6))
	if err != nil {
		return err
	}
	CLUBS_6 = ebiten.NewImageFromImage(cardDecoded)

	cardDecoded, err = png.Decode(bytes.NewBuffer(clubs_png_7))
	if err != nil {
		return err
	}
	CLUBS_7 = ebiten.NewImageFromImage(cardDecoded)

	cardDecoded, err = png.Decode(bytes.NewBuffer(clubs_png_8))
	if err != nil {
		return err
	}
	CLUBS_8 = ebiten.NewImageFromImage(cardDecoded)

	cardDecoded, err = png.Decode(bytes.NewBuffer(clubs_png_9))
	if err != nil {
		return err
	}
	CLUBS_9 = ebiten.NewImageFromImage(cardDecoded)

	cardDecoded, err = png.Decode(bytes.NewBuffer(clubs_png_10))
	if err != nil {
		return err
	}
	CLUBS_10 = ebiten.NewImageFromImage(cardDecoded)

	cardDecoded, err = png.Decode(bytes.NewBuffer(clubs_png_JACK))
	if err != nil {
		return err
	}
	CLUBS_JACK = ebiten.NewImageFromImage(cardDecoded)

	cardDecoded, err = png.Decode(bytes.NewBuffer(clubs_png_KING))
	if err != nil {
		return err
	}
	CLUBS_KING = ebiten.NewImageFromImage(cardDecoded)

	cardDecoded, err = png.Decode(bytes.NewBuffer(clubs_png_QUEEN))
	if err != nil {
		return err
	}
	CLUBS_QUEEN = ebiten.NewImageFromImage(cardDecoded)

	return err
}
