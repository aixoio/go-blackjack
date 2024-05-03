package img

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func loadSpades(cards_path string) error {
	var err error

	SPADES_ACE, _, err = ebitenutil.NewImageFromFile(fmt.Sprintf("%s/spades-ace.png", cards_path))
	if err != nil {
		return err
	}

	SPADES_JACK, _, err = ebitenutil.NewImageFromFile(fmt.Sprintf("%s/spades-jack.png", cards_path))
	if err != nil {
		return err
	}

	SPADES_KING, _, err = ebitenutil.NewImageFromFile(fmt.Sprintf("%s/spades-king.png", cards_path))
	if err != nil {
		return err
	}

	SPADES_QUEEN, _, err = ebitenutil.NewImageFromFile(fmt.Sprintf("%s/spades-queen.png", cards_path))
	if err != nil {
		return err
	}

	SPADES_2, _, err = ebitenutil.NewImageFromFile(fmt.Sprintf("%s/spades-2.png", cards_path))
	if err != nil {
		return err
	}

	SPADES_3, _, err = ebitenutil.NewImageFromFile(fmt.Sprintf("%s/spades-3.png", cards_path))
	if err != nil {
		return err
	}

	SPADES_4, _, err = ebitenutil.NewImageFromFile(fmt.Sprintf("%s/spades-4.png", cards_path))
	if err != nil {
		return err
	}

	SPADES_5, _, err = ebitenutil.NewImageFromFile(fmt.Sprintf("%s/spades-5.png", cards_path))
	if err != nil {
		return err
	}

	SPADES_6, _, err = ebitenutil.NewImageFromFile(fmt.Sprintf("%s/spades-6.png", cards_path))
	if err != nil {
		return err
	}

	SPADES_7, _, err = ebitenutil.NewImageFromFile(fmt.Sprintf("%s/spades-7.png", cards_path))
	if err != nil {
		return err
	}

	SPADES_8, _, err = ebitenutil.NewImageFromFile(fmt.Sprintf("%s/spades-8.png", cards_path))
	if err != nil {
		return err
	}

	SPADES_9, _, err = ebitenutil.NewImageFromFile(fmt.Sprintf("%s/spades-9.png", cards_path))
	if err != nil {
		return err
	}

	SPADES_10, _, err = ebitenutil.NewImageFromFile(fmt.Sprintf("%s/spades-10.png", cards_path))
	if err != nil {
		return err
	}

	return err
}
