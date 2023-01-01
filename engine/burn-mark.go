package engine

import "retro-carnage/engine/geometry"

// BurnMark is a mark on the ground caused by an explosion. They don't disappear.
type BurnMark struct {
	Position       *geometry.Rectangle
	SpriteSupplier *BurnMarkSpriteSupplier
}
