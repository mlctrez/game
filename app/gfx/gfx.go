package gfx

import (
	"embed"
	_ "embed"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
	"log"
)

//go:embed *.png
var resources embed.FS

func Load(name string) *ebiten.Image {
	open, err := resources.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	img, _, err := image.Decode(open)
	if err != nil {
		log.Fatal(err)
	}
	return ebiten.NewImageFromImage(img)
}
