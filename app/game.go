package app

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/mlctrez/game/app/splash"
	"os"
)

func New() *Game {
	if os.Getenv("GOARCH") != "wasm" {
		ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
		ebiten.SetWindowSize(1280, 1024)
		ebiten.SetCursorShape(ebiten.CursorShapeCrosshair)
	}
	return &Game{splash: splash.New()}
}

type Game struct {
	splash *splash.Splash
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyF11) {
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		g.splash.Visible(false)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.splash.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
