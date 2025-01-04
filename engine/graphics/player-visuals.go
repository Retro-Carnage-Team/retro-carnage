package graphics

import "retro-carnage/engine/geometry"

// PlayerVisuals encapsulates all props of players that are relevant for their graphical representation.
type PlayerVisuals interface {
	Dying() bool
	Idle() bool
	Invincible() bool
	PlayerIndex() int
	ViewingDirection() *geometry.Direction
}
