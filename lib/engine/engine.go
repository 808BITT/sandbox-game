package engine

import (
	"errors"
	"sandbox/lib/assets"
	"sandbox/lib/game"
	"sandbox/lib/phys"

	"github.com/hajimehoshi/ebiten/v2"
)

type Engine struct {
	W      int
	H      int
	Player *game.Player
	X      int
	Y      int
}

func New(w, h int) *Engine {
	ebiten.SetFullscreen(true)

	return &Engine{
		W:      w,
		H:      h,
		Player: game.NewPlayer(assets.LoadImage("img/player.png"), phys.NewVec2(300, 300), 1),
		X:      1,
		Y:      1,
	}
}

func (g *Engine) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return errors.New("game closed")
	}

	// update the player based on input
	g.Player.Update()

	// handle player moving off screen
	if g.Player.Position.X < 0 {
		if g.X > 0 {
			g.X--
			g.Player.Position.X = g.Width() - g.Player.Width()
		} else {
			g.Player.Position.X = 0
		}
	} else if g.Player.Position.X > g.Width()-g.Player.Width() {
		if g.X < 3 {
			g.X++
			g.Player.Position.X = 0
		} else {
			g.Player.Position.X = g.Width() - g.Player.Width()
		}
	}
	if g.Player.Position.Y < 0 {
		if g.Y > 0 {
			g.Y--
			g.Player.Position.Y = g.Height() - g.Player.Height()
		} else {
			g.Player.Position.Y = 0
		}
	} else if g.Player.Position.Y > g.Height()-g.Player.Height() {
		if g.Y < 3 {
			g.Y++
			g.Player.Position.Y = 0
		} else {
			g.Player.Position.Y = g.Height() - g.Player.Height()
		}
	}

	return nil
}

func (g *Engine) Draw(screen *ebiten.Image) {
	// make a single color image the size of the screen
	// this is the background

	g.Player.Draw(screen)
}

func (g *Engine) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.W, g.H
}

func (g *Engine) Width() float64 {
	return float64(g.W)
}

func (g *Engine) Height() float64 {
	return float64(g.H)
}
