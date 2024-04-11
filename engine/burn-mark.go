package engine

import "retro-carnage/engine/geometry"

// BurnMark is a mark on the ground caused by an explosion. They don't disappear.
type BurnMark struct {
	position       *geometry.Rectangle
	SpriteSupplier *BurnMarkSpriteSupplier
}

func (bm *BurnMark) Position() *geometry.Rectangle {
	return bm.position
}
