package engine

import (
	"retro-carnage/engine/geometry"
)

const (
	GrenadeHeight = 17
	GrenadeWidth  = 32
	RpgHeight     = 10
	RpgWidth      = 10
)

type Explosive struct {
	distanceMoved     float64
	distanceToTarget  float64
	direction         geometry.Direction
	FiredByPlayer     bool
	FiredByPlayerIdx  int
	Position          geometry.Rectangle
	speed             float64
	SpriteSupplier    ExplosiveSpriteSupplier
	ExplodesOnContact bool
}

// Move moves the explosive on screen.
// Returns true if the explosive reached it's destination
func (e *Explosive) Move(elapsedTimeInMs int64) bool {
	if e.distanceMoved < e.distanceToTarget {
		var maxDistance = e.distanceToTarget - e.distanceMoved
		e.distanceMoved += geometry.CalculateMovementDistance(elapsedTimeInMs, e.speed, &maxDistance)
		e.Position.X += geometry.CalculateMovementX(elapsedTimeInMs, e.direction, e.speed, &maxDistance)
		e.Position.Y += geometry.CalculateMovementY(elapsedTimeInMs, e.direction, e.speed, &maxDistance)
	}
	return e.distanceMoved >= e.distanceToTarget
}
