package res

import (
	"embed"
	_ "embed"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image"
	_ "image/png"
	"io/fs"
	"log"
	"sync"
)

//go:embed img/* ttf/*
var resources embed.FS

func Image(name string) *ebiten.Image {
	open, err := resources.Open("img/" + name)
	if err != nil {
		log.Fatal(err)
	}
	defer func(open fs.File) { _ = open.Close() }(open)
	img, _, err := image.Decode(open)
	if err != nil {
		log.Fatal(err)
	}
	return ebiten.NewImageFromImage(img)
}

func Font(name string) *opentype.Font {
	open, err := resources.ReadFile("ttf/" + name)
	if err != nil {
		log.Fatal(err)
	}
	parse, err := opentype.Parse(open)
	if err != nil {
		log.Fatal(err)
	}
	return parse
}

var fontCache = make(map[string]font.Face)
var fontCacheMutex = &sync.Mutex{}

func CachedFontFace(name string, size float64) font.Face {
	fontCacheMutex.Lock()
	defer fontCacheMutex.Unlock()
	key := fmt.Sprintf("%s/%0.2f", name, size)
	if face, ok := fontCache[key]; ok {
		return face
	}
	opts := &opentype.FaceOptions{Size: size, DPI: 72, Hinting: font.HintingVertical}
	face, err := opentype.NewFace(Font(name), opts)
	if err != nil {
		log.Fatal(err)
	}
	fontCache[key] = face
	return face
}
