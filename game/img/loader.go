package img

import (
	"bytes"
	"image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

func LoadImages() error {
	var err error

	cardDecoded, err := png.Decode(bytes.NewBuffer(card_png))
	if err != nil {
		return err
	}
	CARD = ebiten.NewImageFromImage(cardDecoded)

	err = loadSpades()
	if err != nil {
		return err
	}

	err = loadClubs()
	if err != nil {
		return err
	}

	err = loadDiamonds()
	if err != nil {
		return err
	}

	err = loadHearts()
	if err != nil {
		return err
	}

	return nil
}
