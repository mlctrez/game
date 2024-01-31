package compo

import "github.com/hajimehoshi/ebiten/v2"

type Compo interface {
	Update() (Compo, error)
	Draw(screen *ebiten.Image)
	//Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int)

}
