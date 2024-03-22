package characters

import (
	"retro-carnage/assets"
	"retro-carnage/engine/geometry"
	"retro-carnage/engine/graphics"
	"retro-carnage/logging"

	"github.com/faiface/pixel"
)

const (
	downLeftSprite  = "images/environment/gun-turret-down_left.png"
	downRightSprite = "images/environment/gun-turret-down_right.png"
	upLeftSprite    = "images/environment/gun-turret-up_left.png"
	upRightSprite   = "images/environment/gun-turret-up_right.png"
)

type EnemyGunTurretSpriteSupplier struct {
	source string
	sprite *pixel.Sprite
}

func NewEnemyGunTurretSpriteSupplier(direction geometry.Direction) *EnemyGunTurretSpriteSupplier {
	if !direction.IsDiagonal() {
		logging.Error.Fatalf("Gun turrets can have diagonal directions, only. Found %s instead", direction.Name)
	}

	var source string
	switch {
	case direction == geometry.DownLeft:
		source = downLeftSprite
	case direction == geometry.DownRight:
		source = downRightSprite
	case direction == geometry.UpLeft:
		source = upLeftSprite
	case direction == geometry.UpRight:
		source = upRightSprite
	}

	return &EnemyGunTurretSpriteSupplier{
		source: source,
		sprite: assets.SpriteRepository.Get(source),
	}
}

func (supplier *EnemyGunTurretSpriteSupplier) Sprite(int64, ActiveEnemy) *graphics.SpriteWithOffset {
	return &graphics.SpriteWithOffset{
		Offset: geometry.Point{X: 0, Y: 0},
		Source: supplier.source,
		Sprite: supplier.sprite,
	}
}
