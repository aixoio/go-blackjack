package fonts

import (
	_ "embed"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

//go:embed fonts/Poppins-Regular.ttf
var poppinsRegular []byte

var PoppinsRegular *text.GoTextFaceSource
