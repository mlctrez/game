package app

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/mlctrez/game/app/compo"
	"os"
)

func New() *Game {
	if os.Getenv("GOARCH") != "wasm" {
		ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
		ebiten.SetWindowSize(1280, 1024)
		ebiten.SetCursorShape(ebiten.CursorShapeCrosshair)
	}
	return &Game{compo: compo.Splash()}
}

type Game struct {
	compo compo.Compo
}

func (g *Game) Update() (err error) {
	if inpututil.IsKeyJustPressed(ebiten.KeyF11) {
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
	}
	g.compo, err = g.compo.Update()
	return err
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.compo.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
