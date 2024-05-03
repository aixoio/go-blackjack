package img

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func loadDiamonds(cards_path string) error {
	var err error

	DIAMONDS_ACE, _, err = ebitenutil.NewImageFromFile(fmt.Sprintf("%s/diamonds-ace.png", cards_path))
	if err != nil {
		return err
	}

	DIAMONDS_JACK, _, err = ebitenutil.NewImageFromFile(fmt.Sprintf("%s/diamonds-jack.png", cards_path))
	if err != nil {
		return err
	}

	DIAMONDS_KING, _, err = ebitenutil.NewImageFromFile(fmt.Sprintf("%s/diamonds-king.png", cards_path))
	if err != nil {
		return err
	}

	DIAMONDS_QUEEN, _, err = ebitenutil.NewImageFromFile(fmt.Sprintf("%s/diamonds-queen.png", cards_path))
	if err != nil {
		return err
	}

	DIAMONDS_2, _, err = ebitenutil.NewImageFromFile(fmt.Sprintf("%s/diamonds-2.png", cards_path))
	if err != nil {
		return err
	}

	DIAMONDS_3, _, err = ebitenutil.NewImageFromFile(fmt.Sprintf("%s/diamonds-3.png", cards_path))
	if err != nil {
		return err
	}

	DIAMONDS_4, _, err = ebitenutil.NewImageFromFile(fmt.Sprintf("%s/diamonds-4.png", cards_path))
	if err != nil {
		return err
	}

	DIAMONDS_5, _, err = ebitenutil.NewImageFromFile(fmt.Sprintf("%s/diamonds-5.png", cards_path))
	if err != nil {
		return err
	}

	DIAMONDS_6, _, err = ebitenutil.NewImageFromFile(fmt.Sprintf("%s/diamonds-6.png", cards_path))
	if err != nil {
		return err
	}

	DIAMONDS_7, _, err = ebitenutil.NewImageFromFile(fmt.Sprintf("%s/diamonds-7.png", cards_path))
	if err != nil {
		return err
	}

	DIAMONDS_8, _, err = ebitenutil.NewImageFromFile(fmt.Sprintf("%s/diamonds-8.png", cards_path))
	if err != nil {
		return err
	}

	DIAMONDS_9, _, err = ebitenutil.NewImageFromFile(fmt.Sprintf("%s/diamonds-9.png", cards_path))
	if err != nil {
		return err
	}

	DIAMONDS_10, _, err = ebitenutil.NewImageFromFile(fmt.Sprintf("%s/diamonds-10.png", cards_path))
	if err != nil {
		return err
	}

	return err
}
