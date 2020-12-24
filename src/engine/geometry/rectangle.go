package geometry

import (
	"fmt"
	"math"
)

type Rectangle struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

func (r *Rectangle) Add(offset *Point) *Rectangle {
	r.X += offset.X
	r.Y += offset.Y
	return r
}

func (r *Rectangle) Subtract(offset *Point) *Rectangle {
	r.X -= offset.X
	r.Y -= offset.Y
	return r
}

func (r *Rectangle) Intersection(other *Rectangle) *Rectangle {
	var leftX = math.Max(r.X, other.X)
	var rightX = math.Min(r.X+r.Width, other.X+other.Width)
	var topY = math.Max(r.Y, other.Y)
	var bottomY = math.Min(r.Y+r.Height, other.Y+other.Height)
	if leftX < rightX && topY < bottomY {
		return &Rectangle{X: leftX, Y: topY, Width: rightX - leftX, Height: bottomY - topY}
	}
	return nil
}

func (r *Rectangle) LeftBorder() *Line {
	return &Line{Start: &Point{X: r.X, Y: r.Y}, End: &Point{X: r.X, Y: r.Y + r.Height}}
}

func (r *Rectangle) RightBorder() *Line {
	return &Line{Start: &Point{X: r.X + r.Width, Y: r.Y}, End: &Point{X: r.X + r.Width, Y: r.Y + r.Height}}
}

func (r *Rectangle) TopBorder() *Line {
	return &Line{Start: &Point{X: r.X, Y: r.Y}, End: &Point{X: r.X + r.Width, Y: r.Y}}
}

func (r *Rectangle) BottomBorder() *Line {
	return &Line{Start: &Point{X: r.X, Y: r.Y + r.Height}, End: &Point{X: r.X + r.Width, Y: r.Y + r.Height}}
}

func (r *Rectangle) Center() *Point {
	return &Point{
		X: (r.X + r.X + r.Width) / 2,
		Y: (r.Y + r.Y + r.Height) / 2,
	}
}

func (r *Rectangle) String() string {
	return fmt.Sprintf("Rectangle[x: %.5f, y: %.5f, width: %.5f, height: %.5f]", r.X, r.Y, r.Width, r.Height)
}
