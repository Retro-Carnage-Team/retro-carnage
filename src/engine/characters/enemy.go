package characters

import "retro-carnage/engine/geometry"

type Enemy struct {
	ActivationDistance      float64
	Active                  bool
	Dying                   bool
	DyingAnimationCountDown int64
	Movements               []EnemyMovement
	Position                *geometry.Rectangle
	Skin                    EnemySkin
	SpriteSupplier          EnemySpriteSupplier
	Type                    EnemyType
	ViewingDirection        geometry.Direction
}

func (e *Enemy) Activate() {
	e.Active = true

	if Person == e.Type {
		e.SpriteSupplier = &EnemyPersonSpriteSupplier{
			lastDirection:           e.ViewingDirection,
			durationSinceLastSprite: 0,
			lastIndex:               0,
		}
	}

	if Landmine == e.Type {
		e.SpriteSupplier = &EnemyLandmineSpriteSupplier{}
	}
}

func (e *Enemy) Die() {
	e.Dying = true
	e.DyingAnimationCountDown = 1
	// TODO: Play sound effect when enemy is a person
}
