package img

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func LoadImages(cards_path string) error {
	var err error

	CARD, _, err = ebitenutil.NewImageFromFile(fmt.Sprintf("%s/card.png", cards_path))
	if err != nil {
		return err
	}

	err = loadSpades(cards_path)
	if err != nil {
		return err
	}

	err = loadClubs(cards_path)
	if err != nil {
		return err
	}

	err = loadDiamonds(cards_path)
	if err != nil {
		return err
	}

	err = loadHearts(cards_path)
	if err != nil {
		return err
	}

	return nil
}
