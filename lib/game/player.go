package game

import (
	"sandbox/lib/phys"

	"github.com/hajimehoshi/ebiten/v2"
)

// Player is a struct that represents a player in the game
type Player struct {
	Image     *ebiten.Image
	Position  *phys.Vec2
	Velocity  *phys.Vec2
	Direction *phys.Vec2
	Speed     float64
	IsMoving  bool
}

// NewPlayer creates a new player
func NewPlayer(image *ebiten.Image, position *phys.Vec2, speed float64) *Player {
	return &Player{
		Image:     image,
		Position:  position,
		Velocity:  phys.NewVec2(0, 0),
		Direction: phys.NewVec2(0, 1),
		Speed:     speed,
		IsMoving:  false,
	}
}

// Update updates the player
func (p *Player) Update() {
	p.Velocity = phys.NewVec2(0, 0)

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		p.Velocity = p.Direction.Mul(-1).Mul(p.Speed)
		p.Position = p.Position.Add(p.Velocity)
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		p.Velocity = p.Direction.Mul(1).Mul(p.Speed)
		p.Position = p.Position.Add(p.Velocity)
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.Direction = p.Direction.Rotate(-0.1).Mul(p.Speed)
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.Direction = p.Direction.Rotate(0.1).Mul(p.Speed)
	}

	if p.Velocity.Mag() > 0 {
		p.IsMoving = true
	} else {
		p.IsMoving = false
	}

	// p.Position = p.Position.Add(p.Velocity.Mul(p.Speed))

	// p.Velocity = phys.NewVec2(0, 0)

	// if ebiten.IsKeyPressed(ebiten.KeyW) {
	// 	p.Velocity.Y -= 1
	// }
	// if ebiten.IsKeyPressed(ebiten.KeyS) {
	// 	p.Velocity.Y += 1
	// }
	// if ebiten.IsKeyPressed(ebiten.KeyA) {
	// 	p.Velocity.X -= 1
	// }
	// if ebiten.IsKeyPressed(ebiten.KeyD) {
	// 	p.Velocity.X += 1
	// }

	// if p.Velocity.Mag() > 0 {
	// 	p.Velocity = p.Velocity.Normalize()
	// }

	// p.Position = p.Position.Add(p.Velocity.Mul(p.Speed))

	// if p.Velocity.Mag() > 0 {
	// 	p.Direction = p.Velocity.Normalize()
	// 	p.IsMoving = true
	// } else {
	// 	p.IsMoving = false
	// }
}

// Draw draws the player
func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.Position.X, p.Position.Y)
	screen.DrawImage(p.Image, op)
}

// Width returns the width of the player
func (p *Player) Width() float64 {
	return float64(p.Image.Bounds().Dx())
}

// Height returns the height of the player
func (p *Player) Height() float64 {
	return float64(p.Image.Bounds().Dy())
}

// Center returns the center of the player
func (p *Player) Center() *phys.Vec2 {
	return phys.NewVec2(p.Position.X+p.Width()/2, p.Position.Y+p.Height()/2)
}

// Bounds returns the bounds of the player
func (p *Player) Bounds() *phys.Rect {
	return phys.NewRect(p.Position.X, p.Position.Y, p.Width(), p.Height())
}
