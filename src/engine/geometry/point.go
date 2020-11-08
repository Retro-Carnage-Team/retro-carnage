package geometry

import "fmt"

type Point struct {
	X float32
	Y float32
}

func (p *Point) String() string {
	return fmt.Sprintf("Point[X: %.5f, Y: %.5f]", p.X, p.Y)
}
