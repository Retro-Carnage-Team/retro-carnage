package geometry

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateMovementVectorRight(t *testing.T) {
	var result = CalculateMovementVector(100, Right, 1.5)
	assert.InDelta(t, 150.0, result.X, 0.00001)
	assert.InDelta(t, 0.0, result.Y, 0.00001)
}

func TestCalculateMovementVectorDown(t *testing.T) {
	var result = CalculateMovementVector(100, Down, 1.5)
	assert.InDelta(t, 0.0, result.X, 0.00001)
	assert.InDelta(t, 150.0, result.Y, 0.00001)
}

func TestCalculateMovementVectorDownLeft(t *testing.T) {
	var result = CalculateMovementVector(100, DownLeft, 1.5)
	assert.InDelta(t, -106.066017, result.X, 0.00001)
	assert.InDelta(t, 106.066017, result.Y, 0.00001)
}

func TestCalculateMovementVectorDownRight(t *testing.T) {
	var result = CalculateMovementVector(100, DownRight, 1.5)
	assert.InDelta(t, 106.066017, result.X, 0.00001)
	assert.InDelta(t, 106.066017, result.Y, 0.00001)
}

func TestMoveLeftNoDeviationNoMax(t *testing.T) {
	var rect = Rectangle{X: 0, Y: 0, Width: 1, Height: 1}
	Move(&rect, 100, Left, 0, 1.5, nil)

	assert.InDelta(t, -150, rect.X, 0.00001)
	assert.InDelta(t, 0, rect.Y, 0.00001)

	assert.InDelta(t, 1, rect.Width, 0.00001)
	assert.InDelta(t, 1, rect.Height, 0.00001)
}

func TestMoveLeftWithDeviationNoMax(t *testing.T) {
	var rect = Rectangle{X: 0, Y: 0, Width: 1, Height: 1}
	// deviation is π/4. So direction becomes DownLeft instead of Left
	Move(&rect, 100, Left, math.Pi/4.0, 1.5, nil)

	assert.InDelta(t, -106.066017, rect.X, 0.00001)
	assert.InDelta(t, 106.066017, rect.Y, 0.00001)

	assert.InDelta(t, 1, rect.Width, 0.00001)
	assert.InDelta(t, 1, rect.Height, 0.00001)
}

func TestMoveLeftWithDeviationAndMax(t *testing.T) {
	var rect = Rectangle{X: 0, Y: 0, Width: 1, Height: 1}
	// max reduces the distance from 150 to 75 so that we can reuse the expected values
	var max float64 = 75.0
	// deviation is π/4. So direction becomes DownLeft instead of Left
	Move(&rect, 100, Left, math.Pi/4.0, 1.5, &max)

	assert.InDelta(t, -106.066017/2, rect.X, 0.00001)
	assert.InDelta(t, 106.066017/2, rect.Y, 0.00001)

	assert.InDelta(t, 1, rect.Width, 0.00001)
	assert.InDelta(t, 1, rect.Height, 0.00001)
}
