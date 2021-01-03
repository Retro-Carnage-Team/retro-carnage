package geometry

import (
	"fmt"
	"github.com/faiface/pixel"
)

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func (p *Point) String() string {
	return fmt.Sprintf("Point[X: %.5f, Y: %.5f]", p.X, p.Y)
}

func (p *Point) Add(other *Point) *Point {
	return &Point{
		X: p.X + other.X,
		Y: p.Y + other.Y,
	}
}

func (p *Point) ToVec() pixel.Vec {
	return pixel.V(p.X, p.Y)
}
