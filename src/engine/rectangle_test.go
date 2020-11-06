package engine

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntersectionNoOverlap(t *testing.T) {
	var r1 = &Rectangle{X: 1, Y: 1, Width: 1, Height: 1}
	var r2 = &Rectangle{X: 5, Y: 5, Width: 1, Height: 1}

	assert.Nil(t, r1.GetIntersection(r2))
}

func TestIntersectionOverlap(t *testing.T) {
	var r1 = &Rectangle{X: 1, Y: 1, Width: 10, Height: 10}
	var r2 = &Rectangle{X: 6, Y: 6, Width: 10, Height: 10}
	var result = r1.GetIntersection(r2)

	assert.InDelta(t, 6, result.X, 0.0001)
	assert.InDelta(t, 6, result.Y, 0.0001)
	assert.InDelta(t, 5, result.Width, 0.0001)
	assert.InDelta(t, 5, result.Height, 0.0001)
}

func TestIntersectionContains(t *testing.T) {
	var r1 = &Rectangle{X: 1, Y: 1, Width: 10, Height: 10}
	var r2 = &Rectangle{X: 3, Y: 3, Width: 3, Height: 3}
	var result = r1.GetIntersection(r2)

	assert.InDelta(t, 3, result.X, 0.0001)
	assert.InDelta(t, 3, result.Y, 0.0001)
	assert.InDelta(t, 3, result.Width, 0.0001)
	assert.InDelta(t, 3, result.Height, 0.0001)
}

func TestAddOffsets(t *testing.T) {
	var r1 = &Rectangle{X: 1, Y: 1, Width: 10, Height: 10}
	result := r1.Add(&Point{2, 3})

	assert.InDelta(t, 3, result.X, 0.0001)
	assert.InDelta(t, 4, result.Y, 0.0001)
	assert.InDelta(t, 10, result.Width, 0.0001)
	assert.InDelta(t, 10, result.Height, 0.0001)
}

func TestSubtractOffsets(t *testing.T) {
	var r1 = &Rectangle{X: 1, Y: 1, Width: 10, Height: 10}
	result := r1.Subtract(&Point{2, 3})

	assert.InDelta(t, -1, result.X, 0.0001)
	assert.InDelta(t, -2, result.Y, 0.0001)
	assert.InDelta(t, 10, result.Width, 0.0001)
	assert.InDelta(t, 10, result.Height, 0.0001)
}

func TestLeftBorder(t *testing.T) {
	var r1 = &Rectangle{X: 3, Y: 3, Width: 2, Height: 2}
	border := r1.GetLeftBorder()
	expected := &Line{Start: &Point{3, 3}, End: &Point{3, 5}}

	assert.True(t, border.Equals(expected))
}

func TestRightBorder(t *testing.T) {
	var r1 = &Rectangle{X: 3, Y: 3, Width: 2, Height: 2}
	border := r1.GetRightBorder()
	expected := &Line{Start: &Point{5, 3}, End: &Point{5, 5}}

	assert.True(t, border.Equals(expected))
}

func TestTopBorder(t *testing.T) {
	var r1 = &Rectangle{X: 3, Y: 3, Width: 2, Height: 2}
	border := r1.GetTopBorder()
	expected := &Line{Start: &Point{3, 3}, End: &Point{5, 3}}

	assert.True(t, border.Equals(expected))
}

func TestBottomBorder(t *testing.T) {
	var r1 = &Rectangle{X: 3, Y: 3, Width: 2, Height: 2}
	border := r1.GetBottomBorder()
	expected := &Line{Start: &Point{3, 5}, End: &Point{5, 5}}

	assert.True(t, border.Equals(expected))
}
