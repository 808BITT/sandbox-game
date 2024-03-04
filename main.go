package main

import (
	"sandbox/lib/engine"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	if err := ebiten.RunGame(engine.New(1360/2, 768/2)); err != nil {
		panic(err)
	}
}
