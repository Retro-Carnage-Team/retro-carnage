package engine

import "retro-carnage/engine/graphics"

// ExplosiveSpriteSupplier is an interface common to all explosives.
type ExplosiveSpriteSupplier interface {
	Sprite(int64) *graphics.SpriteWithOffset
}
