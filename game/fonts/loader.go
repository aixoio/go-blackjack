package fonts

import (
	"bytes"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

func LoadFonts() error {
	var err error

	PoppinsRegular, err = text.NewGoTextFaceSource(bytes.NewReader(poppinsRegular))

	return err
}
