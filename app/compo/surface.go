package compo

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func Surface() Compo {
	g := &surface{}
	return g
}

type surface struct {
	bi         *ebiten.Image
	imageSize  float64
	colorScale float64
	//dmMono     font.Face
}

func (s *surface) Update() (Compo, error) {
	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		return Splash(), nil
	}

	return s, nil
}

func (s *surface) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "surface")
}
