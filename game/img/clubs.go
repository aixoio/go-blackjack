package img

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func loadClubs(cards_path string) error {
	var err error

	CLUBS_ACE, _, err = ebitenutil.NewImageFromFile(fmt.Sprintf("%s/clubs-ace.png", cards_path))
	if err != nil {
		return err
	}

	CLUBS_JACK, _, err = ebitenutil.NewImageFromFile(fmt.Sprintf("%s/clubs-jack.png", cards_path))
	if err != nil {
		return err
	}

	CLUBS_KING, _, err = ebitenutil.NewImageFromFile(fmt.Sprintf("%s/clubs-king.png", cards_path))
	if err != nil {
		return err
	}

	CLUBS_QUEEN, _, err = ebitenutil.NewImageFromFile(fmt.Sprintf("%s/clubs-queen.png", cards_path))
	if err != nil {
		return err
	}

	CLUBS_2, _, err = ebitenutil.NewImageFromFile(fmt.Sprintf("%s/clubs-2.png", cards_path))
	if err != nil {
		return err
	}

	CLUBS_3, _, err = ebitenutil.NewImageFromFile(fmt.Sprintf("%s/clubs-3.png", cards_path))
	if err != nil {
		return err
	}

	CLUBS_4, _, err = ebitenutil.NewImageFromFile(fmt.Sprintf("%s/clubs-4.png", cards_path))
	if err != nil {
		return err
	}

	CLUBS_5, _, err = ebitenutil.NewImageFromFile(fmt.Sprintf("%s/clubs-5.png", cards_path))
	if err != nil {
		return err
	}

	CLUBS_6, _, err = ebitenutil.NewImageFromFile(fmt.Sprintf("%s/clubs-6.png", cards_path))
	if err != nil {
		return err
	}

	CLUBS_7, _, err = ebitenutil.NewImageFromFile(fmt.Sprintf("%s/clubs-7.png", cards_path))
	if err != nil {
		return err
	}

	CLUBS_8, _, err = ebitenutil.NewImageFromFile(fmt.Sprintf("%s/clubs-8.png", cards_path))
	if err != nil {
		return err
	}

	CLUBS_9, _, err = ebitenutil.NewImageFromFile(fmt.Sprintf("%s/clubs-9.png", cards_path))
	if err != nil {
		return err
	}

	CLUBS_10, _, err = ebitenutil.NewImageFromFile(fmt.Sprintf("%s/clubs-10.png", cards_path))
	if err != nil {
		return err
	}

	return err
}
