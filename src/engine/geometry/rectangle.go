package geometry

import (
	"fmt"
	"retro-carnage.net/util"
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

func (r *Rectangle) GetIntersection(other *Rectangle) *Rectangle {
	var mu = util.MathUtil{}

	var leftX = mu.Max(r.X, other.X)
	var rightX = mu.Min(r.X+r.Width, other.X+other.Width)
	var topY = mu.Max(r.Y, other.Y)
	var bottomY = mu.Min(r.Y+r.Height, other.Y+other.Height)
	if leftX < rightX && topY < bottomY {
		return &Rectangle{X: leftX, Y: topY, Width: rightX - leftX, Height: bottomY - topY}
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
