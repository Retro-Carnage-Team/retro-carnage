package characters

import (
	"retro-carnage/engine/geometry"
)

type ActiveEnemy struct {
	Dying                   bool
	DyingAnimationCountDown int64
	Movements               []*EnemyMovement
	position                *geometry.Rectangle
	Skin                    EnemySkin
	SpriteSupplier          EnemySpriteSupplier
	Type                    EnemyType
	ViewingDirection        *geometry.Direction
}

func (e *ActiveEnemy) Die() {
	e.Dying = true
	e.DyingAnimationCountDown = 1
	// TODO: Play sound effect when enemy is a person
}

func (e *ActiveEnemy) Position() *geometry.Rectangle {
	return e.position
}

func (e *ActiveEnemy) SetPosition(pos *geometry.Rectangle) {
	e.position = pos
}
