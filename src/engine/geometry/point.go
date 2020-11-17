package geometry

import "fmt"

type Point struct {
	X float64
	Y float64
}

func (p *Point) String() string {
	return fmt.Sprintf("Point[X: %.5f, Y: %.5f]", p.X, p.Y)
}
