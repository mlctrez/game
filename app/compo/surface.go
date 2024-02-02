package compo

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/mlctrez/game/app/res"
	"golang.org/x/image/font"
	"image/color"
)

func Surface() Compo {
	g := &surface{}
	g.dmMono = res.CachedFontFace("DMMono-Medium.ttf", 8)
	return g
}

type surface struct {
	bi         *ebiten.Image
	imageSize  float64
	colorScale float64
	dmMono     font.Face
}

func (s *surface) Update() (Compo, error) {
	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		return Splash(), nil
	}
	return s, nil
}

func (s *surface) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "surface")
	dx := screen.Bounds().Dx()
	dy := screen.Bounds().Dy()

	for y := 0; y < dy; y = y + 100 {
		vector.StrokeLine(screen, 0, float32(y), float32(dx), float32(y), 1, color.White, false)
	}
	for x := 0; x < dx; x = x + 100 {
		vector.StrokeLine(screen, float32(x), 0, float32(x), float32(dy), 1, color.White, false)
	}

}
