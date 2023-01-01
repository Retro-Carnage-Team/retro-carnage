package geometry

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateMovementDistance(t *testing.T) {
	var result = CalculateMovementDistance(int64(10), 1.3, nil)
	assert.InDelta(t, 13.0, result, 0.00001)
}

func TestCalculateMovementDistanceWithMaximum(t *testing.T) {
	var max = 11.0
	var result = CalculateMovementDistance(10, 1.3, &max)
	assert.InDelta(t, 11, result, 0.00001)
}

func TestCalculateMovementXLeft(t *testing.T) {
	var result = CalculateMovementX(100, Left, 0.75, nil)
	assert.InDelta(t, -75.0, result, 0.00001)
}

func TestCalculateMovementXUp(t *testing.T) {
	var result = CalculateMovementX(100, Up, 0.75, nil)
	assert.InDelta(t, 0.0, result, 0.00001)
}

func TestCalculateMovementXUpRight(t *testing.T) {
	var result = CalculateMovementX(48, UpRight, 0.75, nil)
	assert.InDelta(t, 25.0, result, 0.00001)
}

func TestCalculateMovementYLeft(t *testing.T) {
	var result = CalculateMovementY(100, Left, 0.75, nil)
	assert.InDelta(t, 0.0, result, 0.00001)
}

func TestCalculateMovementYUp(t *testing.T) {
	var result = CalculateMovementY(100, Up, 0.75, nil)
	assert.InDelta(t, -75.0, result, 0.00001)
}

func TestCalculateMovementYUpRight(t *testing.T) {
	var result = CalculateMovementY(48, UpRight, 0.75, nil)
	assert.InDelta(t, -25.0, result, 0.00001)
}
