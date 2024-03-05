package main

import (
	"sandbox/lib/engine"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	if err := ebiten.RunGame(engine.New(1920, 1080)); err != nil {
		panic(err)
	}
}
