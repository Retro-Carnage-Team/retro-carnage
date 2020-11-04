package engine

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntersectionNoOverlap(t *testing.T) {
	var r1 = NewRectangle(1, 1, 1, 1)
	var r2 = NewRectangle(5, 5, 1, 1)

	assert.Nil(t, r1.GetIntersection(r2))
}

func TestIntersectionOverlap(t *testing.T) {
	var r1 = NewRectangle(1, 1, 10, 10)
	var r2 = NewRectangle(6, 6, 10, 10)
	var result = r1.GetIntersection(r2)

	assert.InDelta(t, 6, result.X, 0.0001)
	assert.InDelta(t, 6, result.Y, 0.0001)
	assert.InDelta(t, 5, result.Width, 0.0001)
	assert.InDelta(t, 5, result.Height, 0.0001)
}

func TestIntersectionContains(t *testing.T) {
	var r1 = NewRectangle(1, 1, 10, 10)
	var r2 = NewRectangle(3, 3, 3, 3)
	var result = r1.GetIntersection(r2)

	assert.InDelta(t, 3, result.X, 0.0001)
	assert.InDelta(t, 3, result.Y, 0.0001)
	assert.InDelta(t, 3, result.Width, 0.0001)
	assert.InDelta(t, 3, result.Height, 0.0001)
}

func TestAddOffsets(t *testing.T) {
	var r1 = NewRectangle(1, 1, 10, 10)
	result := r1.Add(NewPoint(2, 3))

	assert.InDelta(t, 3, result.X, 0.0001)
	assert.InDelta(t, 4, result.Y, 0.0001)
	assert.InDelta(t, 10, result.Width, 0.0001)
	assert.InDelta(t, 10, result.Height, 0.0001)
}

func TestSubtractOffsets(t *testing.T) {
	var r1 = NewRectangle(1, 1, 10, 10)
	result := r1.Subtract(NewPoint(2, 3))

	assert.InDelta(t, -1, result.X, 0.0001)
	assert.InDelta(t, -2, result.Y, 0.0001)
	assert.InDelta(t, 10, result.Width, 0.0001)
	assert.InDelta(t, 10, result.Height, 0.0001)
}

func TestLeftBorder(t *testing.T) {
	var r1 = NewRectangle(3, 3, 2, 2)
	border := r1.GetLeftBorder()
	expected := NewLine(NewPoint(3, 3), NewPoint(3, 5))

	assert.True(t, border.Equals(expected))
}

func TestRightBorder(t *testing.T) {
	var r1 = NewRectangle(3, 3, 2, 2)
	border := r1.GetRightBorder()
	expected := NewLine(NewPoint(5, 3), NewPoint(5, 5))

	assert.True(t, border.Equals(expected))
}

func TestTopBorder(t *testing.T) {
	var r1 = NewRectangle(3, 3, 2, 2)
	border := r1.GetTopBorder()
	expected := NewLine(NewPoint(3, 3), NewPoint(5, 3))

	assert.True(t, border.Equals(expected))
}

func TestBottomBorder(t *testing.T) {
	var r1 = NewRectangle(3, 3, 2, 2)
	border := r1.GetBottomBorder()
	expected := NewLine(NewPoint(3, 5), NewPoint(5, 5))

	assert.True(t, border.Equals(expected))
}
