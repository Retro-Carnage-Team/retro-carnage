package graphics

import "retro-carnage/engine/geometry"

// EnemyVisuals encapsulates all props of active enemies that are relevant for their graphical representation.
type EnemyVisuals interface {
	Dying() bool
	Skin() EnemySkin
	ViewingDirection() *geometry.Direction
}
