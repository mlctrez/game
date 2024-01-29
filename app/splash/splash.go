package splash

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mlctrez/game/app/gfx"
	"math"
)

func New() *Splash {
	g := &Splash{}
	g.visible = true
	g.bi = gfx.Load("splash.png")
	g.imageSize = float64(g.bi.Bounds().Dx())
	g.colorScale = 0.0
	return g
}

type Splash struct {
	visible    bool
	bi         *ebiten.Image
	imageSize  float64
	colorScale float64
}

func (s *Splash) FadeIn() {
	s.colorScale = 0.0
}

func (s *Splash) Draw(screen *ebiten.Image) {
	if !s.visible {
		return
	}
	sx := float64(screen.Bounds().Dx())
	sy := float64(screen.Bounds().Dy())

	scale := (math.Min(sx, sy) / s.imageSize) * 0.95

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(sx/2-(s.imageSize*scale)/2, sy/2-(s.imageSize*scale)/2)
	if s.colorScale < 1 {
		s.colorScale += 0.01
		colorScale := ebiten.ColorScale{}
		colorScale.ScaleAlpha(float32(s.colorScale))
		op.ColorScale = colorScale
	}
	screen.DrawImage(s.bi, op)
}

func (s *Splash) Visible(v bool) {
	s.visible = v
}
