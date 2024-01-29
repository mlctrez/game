package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mlctrez/game/app"
	"log"
)

func main() {
	if err := ebiten.RunGame(app.New()); err != nil {
		log.Fatal(err)
	}
}
