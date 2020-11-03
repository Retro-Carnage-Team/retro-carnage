package engine

import "fmt"

type Point struct {
	X float32
	Y float32
}

func (p *Point) String() string {
	return fmt.Sprintf("%.5f/%.5f", p.X, p.Y)
}

func NewPoint(x float32, y float32) Point {
	var p Point
	p.X = x
	p.Y = y
	return p
}
