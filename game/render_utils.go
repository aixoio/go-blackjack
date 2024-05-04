package game

import "github.com/hajimehoshi/ebiten/v2"

func DrawImageAt(x, y int, screen, img *ebiten.Image) {
	locationGeoM := ebiten.GeoM{}
	locationGeoM.Translate(float64(x), float64(y))

	screen.DrawImage(img, &ebiten.DrawImageOptions{
		GeoM: locationGeoM,
	})
}
