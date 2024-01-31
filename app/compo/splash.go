package compo

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/mlctrez/game/app/res"
	"math"
)

func Splash() *splash {
	g := &splash{}
	g.bi = res.Image("splash.png")
	g.imageSize = float64(g.bi.Bounds().Dx())
	g.colorScale = 0.0
	return g
}

type splash struct {
	bi         *ebiten.Image
	imageSize  float64
	colorScale float64
}

func (s *splash) Update() (Compo, error) {
	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		return Surface(), nil
	}
	return s, nil
}

func (s *splash) Draw(screen *ebiten.Image) {
	sx := float64(screen.Bounds().Dx())
	sy := float64(screen.Bounds().Dy())

	scale := (math.Min(sx, sy) / s.imageSize) * 0.95

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(sx/2-(s.imageSize*scale)/2, sy/2-(s.imageSize*scale)/2)
	if s.colorScale < 1 {
		s.colorScale += 0.02
		colorScale := ebiten.ColorScale{}
		colorScale.ScaleAlpha(float32(s.colorScale))
		op.ColorScale = colorScale
	}
	screen.DrawImage(s.bi, op)
}
