package engine

import (
	"retro-carnage/assets"
	"retro-carnage/engine/geometry"
)

const (
	MinPlayerDistanceToBorder = 25
	// PlayerMovementPerMs is the player's speed in pixel / second.
	// The value is from this calculation: Screen.width = 1500 / 2.000 milliseconds = 0.75 px / ms
	PlayerMovementPerMs = 0.75
)

func UpdatePlayerMovement(
	elapsedTimeInMs int64,
	direction geometry.Direction,
	oldPosition *geometry.Rectangle,
	obstacles []assets.Obstacle,
) *geometry.Rectangle {
	var movement = getMovementVector(elapsedTimeInMs, direction)
	var updated = false
	var updatedPosition = oldPosition
	for _, obstacle := range obstacles {
		var restrictedMovement = geometry.StopMovementOnCollision(updatedPosition, &obstacle.Rectangle, direction, &movement)
		if nil != restrictedMovement {
			updated = true
			updatedPosition = restrictedMovement
		}
	}

	if updated {
		return limitPlayerMovementToScreenArea(updatedPosition)
	}
	return limitPlayerMovementToScreenArea(oldPosition.Add(&movement))
}

func getMovementVector(elapsedTimeInMs int64, direction geometry.Direction) geometry.Point {
	return geometry.Point{
		X: geometry.CalculateMovementX(elapsedTimeInMs, direction, PlayerMovementPerMs, nil),
		Y: geometry.CalculateMovementY(elapsedTimeInMs, direction, PlayerMovementPerMs, nil),
	}
}

func limitPlayerMovementToScreenArea(position *geometry.Rectangle) *geometry.Rectangle {
	if position.X < MinPlayerDistanceToBorder {
		position.X = MinPlayerDistanceToBorder
	}
	if position.X > ScreenSize-MinPlayerDistanceToBorder-PlayerHitRectWidth {
		position.X = ScreenSize - MinPlayerDistanceToBorder - PlayerHitRectWidth
	}
	if position.Y < MinPlayerDistanceToBorder {
		position.Y = MinPlayerDistanceToBorder
	}
	if position.Y > ScreenSize-MinPlayerDistanceToBorder-PlayerHitRectHeight {
		position.Y = ScreenSize - MinPlayerDistanceToBorder - PlayerHitRectHeight
	}
	return position
}
