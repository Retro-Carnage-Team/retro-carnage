package geometry

import "fmt"

type Point struct {
	X float32
	Y float32
}

func (p *Point) String() string {
	return fmt.Sprintf("%.5f/%.5f", p.X, p.Y)
}
