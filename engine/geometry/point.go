package geometry

import (
	"fmt"
	"math"

	pixel "github.com/Retro-Carnage-Team/pixel2"
)

// Point is a specific spot on a 2d plane
type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// String returns a textual representation of this Point
func (p *Point) String() string {
	return fmt.Sprintf("Point[X: %.5f, Y: %.5f]", p.X, p.Y)
}

// Add returns a new Point that represents a vector addition of this and the given spot.
// This instance will not be modified.
func (p *Point) Add(other *Point) *Point {
	return &Point{
		X: p.X + other.X,
		Y: p.Y + other.Y,
	}
}

// Multiply returns the result of a scalar vector multiplication of this point and the given scalar value.
// This instance will not be modified.
func (p *Point) Multiply(factor float64) *Point {
	return &Point{
		X: p.X * factor,
		Y: p.Y * factor,
	}
}

// ToVec creates a new pixel.Vec with the values of this Point.
func (p *Point) ToVec() pixel.Vec {
	return pixel.V(p.X, p.Y)
}

// Zero returns true if the spot is exactly on the center coordinate of the 2d plane.
func (p *Point) Zero() bool {
	return (math.Abs(p.X) < 0.00001) && (math.Abs(p.Y) < 0.00001)
}
