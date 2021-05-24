package engine

import "retro-carnage/engine/graphics"

type ExplosiveSpriteSupplier interface {
	Sprite() *graphics.SpriteWithOffset
}
