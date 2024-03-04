package phys

import "math"

// Vec2 is a struct that represents a 2D vector
type Vec2 struct {
	X float64
	Y float64
}

func NewVec2(x, y float64) *Vec2 {
	return &Vec2{X: x, Y: y}
}

// Add adds two vectors together
func (v Vec2) Add(v2 *Vec2) *Vec2 {
	return &Vec2{X: v.X + v2.X, Y: v.Y + v2.Y}
}

// Sub subtracts two vectors
func (v Vec2) Sub(v2 *Vec2) *Vec2 {
	return &Vec2{X: v.X - v2.X, Y: v.Y - v2.Y}
}

// Mul multiplies a vector by a scalar
func (v Vec2) Mul(s float64) *Vec2 {
	return &Vec2{X: v.X * s, Y: v.Y * s}
}

// Div divides a vector by a scalar
func (v Vec2) Div(s float64) *Vec2 {
	return &Vec2{X: v.X / s, Y: v.Y / s}
}

// Mag returns the magnitude of a vector
func (v Vec2) Mag() float64 {
	return Sqrt(v.X*v.X + v.Y*v.Y)
}

// Normalize returns the normalized vector
func (v Vec2) Normalize() *Vec2 {
	return v.Div(v.Mag())
}

// Dot returns the dot product of two vectors
func (v Vec2) Dot(v2 *Vec2) float64 {
	return v.X*v2.X + v.Y*v2.Y
}

// Cross returns the cross product of two vectors
func (v Vec2) Cross(v2 *Vec2) float64 {
	return v.X*v2.Y - v.Y*v2.X
}

// Lerp returns the linear interpolation of two vectors
func (v Vec2) Lerp(v2 Vec2, t float64) *Vec2 {
	return v.Add(v2.Sub(&v).Mul(t))
}

// Angle returns the angle between two vectors
func (v Vec2) Angle(v2 *Vec2) float64 {
	return math.Acos(v.Dot(v2) / (v.Mag() * v2.Mag()))
}

// Rotate returns the vector rotated by an angle in radians
func (v Vec2) Rotate(a float64) Vec2 {
	return Vec2{
		X: v.X*math.Cos(a) - v.Y*math.Sin(a),
		Y: v.X*math.Sin(a) + v.Y*math.Cos(a),
	}
}

// Dist returns the distance between two vectors
func (v Vec2) Dist(v2 *Vec2) float64 {
	return v.Sub(v2).Mag()
}

// SqrMag returns the squared magnitude of a vector
func (v Vec2) SqrMag() float64 {
	return v.X*v.X + v.Y*v.Y
}

// Zero returns a zero vector
func Zero() *Vec2 {
	return &Vec2{X: 0, Y: 0}
}

// UnitX returns a unit vector in the x direction
func UnitX() *Vec2 {
	return &Vec2{X: 1, Y: 0}
}

// UnitY returns a unit vector in the y direction
func UnitY() *Vec2 {
	return &Vec2{X: 0, Y: 1}
}

// Sqrt returns the square root of a number
func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}

// Abs returns the absolute value of a number
func Abs(x float64) float64 {
	return math.Abs(x)
}
