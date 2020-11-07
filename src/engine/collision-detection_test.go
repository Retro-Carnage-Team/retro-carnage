package engine

import (
	"github.com/stretchr/testify/assert"
	"retro-carnage.net/engine/geometry"
	"testing"
)

//--- Up -------------------------------------------------------------------------------------------------------------//

func TestShouldFindCollisionForRectMovingUpAgainstCenterOfLargerRect(t *testing.T) {
	var movingRect = &geometry.Rectangle{X: 5, Y: 10, Width: 5, Height: 2}
	var stillRect = &geometry.Rectangle{X: 2, Y: 5, Width: 25, Height: 2}
	var collision = getCollisionForMovementUp(movingRect, stillRect, &geometry.Point{X: 0, Y: -5})

	assert.NotNil(t, collision)
	assert.Equal(t, float32(7), collision.Y)
}

func TestShouldFindCollisionForRectMovingUpAgainstAnotherSmallRectLeft(t *testing.T) {
	var movingRect = &geometry.Rectangle{X: 2, Y: 4, Width: 2, Height: 2}
	var stillRect = &geometry.Rectangle{X: 1, Y: 1, Width: 2, Height: 2}
	var collision = getCollisionForMovementUp(movingRect, stillRect, &geometry.Point{X: 0, Y: -4})

	assert.NotNil(t, collision)
	assert.Equal(t, float32(3), collision.Y)
}

func TestShouldFindCollisionForRectMovingUpAgainstAnotherSmallRectRight(t *testing.T) {
	var movingRect = &geometry.Rectangle{X: 2, Y: 4, Width: 2, Height: 2}
	var stillRect = &geometry.Rectangle{X: 3, Y: 1, Width: 2, Height: 2}
	var collision = getCollisionForMovementUp(movingRect, stillRect, &geometry.Point{X: 0, Y: -4})

	assert.NotNil(t, collision)
	assert.Equal(t, float32(3), collision.Y)
}

func TestShouldStopTheUpMovementOfARectAgainstCenterOfLargerRect(t *testing.T) {
	var movingRect = &geometry.Rectangle{X: 5, Y: 10, Width: 5, Height: 2}
	var stillRect = &geometry.Rectangle{X: 2, Y: 5, Width: 25, Height: 2}
	var result = StopMovementOnCollision(movingRect, stillRect, Up, &geometry.Point{X: 0, Y: -5})

	assert.NotNil(t, result)
	assert.InDelta(t, float32(5), result.X, 0.00001)
	assert.InDelta(t, float32(7), result.Y, 0.00001)
	assert.InDelta(t, float32(5), result.Width, 0.00001)
	assert.InDelta(t, float32(2), result.Height, 0.00001)
}

func TestShouldStopTheUpMovementOfARectAgainstCenterOfSmallerRect(t *testing.T) {
	var movingRect = &geometry.Rectangle{X: 5, Y: 10, Width: 5, Height: 2}
	var stillRect = &geometry.Rectangle{X: 7, Y: 5, Width: 1, Height: 2}
	var result = StopMovementOnCollision(movingRect, stillRect, Up, &geometry.Point{X: 0, Y: -6})

	assert.NotNil(t, result)
	assert.InDelta(t, float32(5), result.X, 0.00001)
	assert.InDelta(t, float32(7), result.Y, 0.00001)
	assert.InDelta(t, float32(5), result.Width, 0.00001)
	assert.InDelta(t, float32(2), result.Height, 0.00001)
}
