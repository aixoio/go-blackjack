package img

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func loadHearts(cards_path string) error {
	var err error

	HEARTS_ACE, _, err = ebitenutil.NewImageFromFile(fmt.Sprintf("%s/hearts-ace.png", cards_path))
	if err != nil {
		return err
	}

	HEARTS_JACK, _, err = ebitenutil.NewImageFromFile(fmt.Sprintf("%s/hearts-jack.png", cards_path))
	if err != nil {
		return err
	}

	HEARTS_KING, _, err = ebitenutil.NewImageFromFile(fmt.Sprintf("%s/hearts-king.png", cards_path))
	if err != nil {
		return err
	}

	HEARTS_QUEEN, _, err = ebitenutil.NewImageFromFile(fmt.Sprintf("%s/hearts-queen.png", cards_path))
	if err != nil {
		return err
	}

	HEARTS_2, _, err = ebitenutil.NewImageFromFile(fmt.Sprintf("%s/hearts-2.png", cards_path))
	if err != nil {
		return err
	}

	HEARTS_3, _, err = ebitenutil.NewImageFromFile(fmt.Sprintf("%s/hearts-3.png", cards_path))
	if err != nil {
		return err
	}

	HEARTS_4, _, err = ebitenutil.NewImageFromFile(fmt.Sprintf("%s/hearts-4.png", cards_path))
	if err != nil {
		return err
	}

	HEARTS_5, _, err = ebitenutil.NewImageFromFile(fmt.Sprintf("%s/hearts-5.png", cards_path))
	if err != nil {
		return err
	}

	HEARTS_6, _, err = ebitenutil.NewImageFromFile(fmt.Sprintf("%s/hearts-6.png", cards_path))
	if err != nil {
		return err
	}

	HEARTS_7, _, err = ebitenutil.NewImageFromFile(fmt.Sprintf("%s/hearts-7.png", cards_path))
	if err != nil {
		return err
	}

	HEARTS_8, _, err = ebitenutil.NewImageFromFile(fmt.Sprintf("%s/hearts-8.png", cards_path))
	if err != nil {
		return err
	}

	HEARTS_9, _, err = ebitenutil.NewImageFromFile(fmt.Sprintf("%s/hearts-9.png", cards_path))
	if err != nil {
		return err
	}

	HEARTS_10, _, err = ebitenutil.NewImageFromFile(fmt.Sprintf("%s/hearts-10.png", cards_path))
	if err != nil {
		return err
	}

	return err
}
