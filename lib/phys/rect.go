package phys

import "math"

// Rect is a struct that represents a rectangle
type Rect struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

// NewRect creates a new rectangle
func NewRect(x, y, w, h float64) *Rect {
	return &Rect{
		X:      x,
		Y:      y,
		Width:  w,
		Height: h,
	}
}

// Contains returns whether a point is inside the rectangle
func (r Rect) Contains(x, y float64) bool {
	return x >= r.X && x <= r.X+r.Width && y >= r.Y && y <= r.Y+r.Height
}

// Intersects returns whether two rectangles intersect
func (r Rect) Intersects(r2 *Rect) bool {
	return r.X < r2.X+r2.Width && r.X+r.Width > r2.X && r.Y < r2.Y+r2.Height && r.Y+r.Height > r2.Y
}

// Center returns the center of the rectangle
func (r Rect) Center() *Vec2 {
	return NewVec2(r.X+r.Width/2, r.Y+r.Height/2)
}

// TopLeft returns the top left corner of the rectangle
func (r Rect) TopLeft() *Vec2 {
	return NewVec2(r.X, r.Y)
}

// TopRight returns the top right corner of the rectangle
func (r Rect) TopRight() *Vec2 {
	return NewVec2(r.X+r.Width, r.Y)
}

// BottomLeft returns the bottom left corner of the rectangle
func (r Rect) BottomLeft() *Vec2 {
	return NewVec2(r.X, r.Y+r.Height)
}

// BottomRight returns the bottom right corner of the rectangle
func (r Rect) BottomRight() *Vec2 {
	return NewVec2(r.X+r.Width, r.Y+r.Height)
}

// Bounds returns the bounds of the rectangle
func (r Rect) Bounds() *Rect {
	return NewRect(r.X, r.Y, r.Width, r.Height)
}

// Area returns the area of the rectangle
func (r Rect) Area() float64 {
	return r.Width * r.Height
}

// Perimeter returns the perimeter of the rectangle
func (r Rect) Perimeter() float64 {
	return 2*r.Width + 2*r.Height
}

// Expand expands the rectangle by a certain amount
func (r Rect) Expand(amount float64) *Rect {
	return NewRect(r.X-amount, r.Y-amount, r.Width+amount*2, r.Height+amount*2)
}

// Translate translates the rectangle by a certain amount
func (r Rect) Translate(x, y float64) *Rect {
	return NewRect(r.X+x, r.Y+y, r.Width, r.Height)
}

// Scale scales the rectangle by a certain amount
func (r Rect) Scale(s float64) *Rect {
	return NewRect(r.X, r.Y, r.Width*s, r.Height*s)
}

// Union returns the union of two rectangles
func (r Rect) Union(r2 *Rect) *Rect {
	x := math.Min(r.X, r2.X)
	y := math.Min(r.Y, r2.Y)
	w := math.Max(r.X+r.Width, r2.X+r2.Width) - x
	h := math.Max(r.Y+r.Height, r2.Y+r2.Height) - y
	return NewRect(x, y, w, h)
}

// Intersection returns the intersection of two rectangles
func (r Rect) Intersection(r2 *Rect) *Rect {
	x := math.Max(r.X, r2.X)
	y := math.Max(r.Y, r2.Y)
	w := math.Min(r.X+r.Width, r2.X+r2.Width) - x
	h := math.Min(r.Y+r.Height, r2.Y+r2.Height) - y
	if w < 0 || h < 0 {
		return NewRect(0, 0, 0, 0)
	}
	return NewRect(x, y, w, h)
}

// Overlaps returns whether two rectangles overlap
func (r Rect) Overlaps(r2 *Rect) bool {
	return r.Intersects(r2)
}
