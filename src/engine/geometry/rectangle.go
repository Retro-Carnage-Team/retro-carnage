package geometry

import (
	"fmt"
)

type Rectangle struct {
	X      float32
	Y      float32
	Width  float32
	Height float32
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

func (r *Rectangle) GetIntersection(other *Rectangle) *Rectangle {
	var leftX = max(r.X, other.X)
	var rightX = min(r.X+r.Width, other.X+other.Width)
	var topY = max(r.Y, other.Y)
	var bottomY = min(r.Y+r.Height, other.Y+other.Height)
	if leftX < rightX && topY < bottomY {
		var result Rectangle
		result.X = leftX
		result.Y = topY
		result.Width = rightX - leftX
		result.Height = bottomY - topY
		return &result
	}
	return nil
}

func (r *Rectangle) GetLeftBorder() *Line {
	return &Line{Start: &Point{X: r.X, Y: r.Y}, End: &Point{X: r.X, Y: r.Y + r.Height}}
}

func (r *Rectangle) GetRightBorder() *Line {
	return &Line{Start: &Point{X: r.X + r.Width, Y: r.Y}, End: &Point{X: r.X + r.Width, Y: r.Y + r.Height}}
}

func (r *Rectangle) GetTopBorder() *Line {
	return &Line{Start: &Point{X: r.X, Y: r.Y}, End: &Point{X: r.X + r.Width, Y: r.Y}}
}

func (r *Rectangle) GetBottomBorder() *Line {
	return &Line{Start: &Point{X: r.X, Y: r.Y + r.Height}, End: &Point{X: r.X + r.Width, Y: r.Y + r.Height}}
}

func (r *Rectangle) String() string {
	return fmt.Sprintf("Rectangle[x: %.5f, y: %.5f, width: %.5f, height: %.5f]", r.X, r.Y, r.Width, r.Height)
}
