package splash

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/mlctrez/game/app/res"
	"golang.org/x/image/font"
	"image/color"
)

func New() *Splash {
	g := &Splash{}
	g.visible = true
	g.bi = res.Image("splash.png")
	g.imageSize = float64(g.bi.Bounds().Dx())
	g.colorScale = 0.0
	g.dmMono = res.CachedFontFace("DMMono-Medium.ttf", 24)

	return g
}

type Splash struct {
	visible    bool
	bi         *ebiten.Image
	imageSize  float64
	colorScale float64
	dmMono     font.Face
}

func (s *Splash) FadeIn() {
	s.colorScale = 0.0
}

func (s *Splash) Draw(screen *ebiten.Image) {
	if !s.visible {
		return
	}

	tps := ebiten.ActualTPS()
	if tps > 0 {
		msg := fmt.Sprintf("TPS: %2.2f ", tps)
		//text.Draw(screen, msg, s.dmMono, 3, 21, color.RGBA{B: 0xff, A: 0xff})
		text.Draw(screen, msg, s.dmMono, 0, 26, color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff})
	}

	//sx := float64(screen.Bounds().Dx())
	//sy := float64(screen.Bounds().Dy())
	//
	//scale := (math.Min(sx, sy) / s.imageSize) * 0.95
	//
	//op := &ebiten.DrawImageOptions{}
	//op.GeoM.Scale(scale, scale)
	//op.GeoM.Translate(sx/2-(s.imageSize*scale)/2, sy/2-(s.imageSize*scale)/2)
	//if s.colorScale < 1 {
	//	s.colorScale += 0.02
	//	colorScale := ebiten.ColorScale{}
	//	colorScale.ScaleAlpha(float32(s.colorScale))
	//	op.ColorScale = colorScale
	//}
	//screen.DrawImage(s.bi, op)
}

func (s *Splash) Visible(v bool) {
	s.visible = v
}
