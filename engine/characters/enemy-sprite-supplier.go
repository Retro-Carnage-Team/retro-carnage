package characters

import (
	"retro-carnage/engine/graphics"
)

type EnemySpriteSupplier interface {
	GetDurationOfEnemyDeathAnimation() int64
	Sprite(elapsedTimeInMs int64) *graphics.SpriteWithOffset
}
