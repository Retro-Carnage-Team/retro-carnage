package engine

import (
	"retro-carnage/assets"
	"retro-carnage/engine/characters"
	"retro-carnage/engine/geometry"
)

const (
	BulletHeight          = 5
	BulletWidth           = 5
	EnemyBulletSpeed      = 1.4
	EnemyBulletRange      = 500
	ExplosiveBulletHeight = 8
	ExplosiveBulletWidth  = 8
)

// Bullet is a projectile that has been fired by a player or enemy.
type Bullet struct {
	distanceMoved    float64
	distanceToTarget float64
	direction        geometry.Direction
	explodes         bool
	firedByPlayer    bool
	playerIdx        int
	position         *geometry.Rectangle
	speed            float64
}

// NewBulletFiredByPlayer creates and returns a new instance of Bullet.
func NewBulletFiredByPlayer(
	playerIdx int,
	playerPosition *geometry.Rectangle,
	direction geometry.Direction,
	selectedWeapon *assets.Weapon,
) (result *Bullet) {

	var bulletHeight float64 = BulletHeight
	var bulletWidth float64 = BulletWidth
	var ammo = assets.AmmunitionCrate.GetByName(selectedWeapon.Ammo)
	if ammo.Explosive {
		bulletHeight = ExplosiveBulletHeight
		bulletWidth = ExplosiveBulletWidth
	}

	result = &Bullet{
		distanceMoved:    0,
		distanceToTarget: float64(selectedWeapon.BulletRange),
		direction:        direction,
		explodes:         ammo.Explosive,
		firedByPlayer:    true,
		playerIdx:        playerIdx,
		position: &geometry.Rectangle{
			X:      playerPosition.X,
			Y:      playerPosition.Y,
			Width:  bulletHeight,
			Height: bulletWidth,
		},
		speed: selectedWeapon.BulletSpeed,
	}

	var offsetValue = characters.SkinForPlayer(playerIdx).BulletOffsets[direction.Name]
	result.position.Add(&offsetValue)
	return result
}

// NewBulletFiredByEnemy creates and returns a new instance of Bullet.
func NewBulletFiredByEnemy(enemy *characters.ActiveEnemy) (result *Bullet) {
	result = &Bullet{
		distanceMoved:    0,
		distanceToTarget: float64(EnemyBulletRange),
		direction:        *enemy.ViewingDirection,
		firedByPlayer:    false,
		position: &geometry.Rectangle{
			X:      enemy.Position().X,
			Y:      enemy.Position().Y,
			Width:  BulletWidth,
			Height: BulletHeight,
		},
		speed: EnemyBulletSpeed,
	}

	var skin = characters.GetEnemySkin(enemy.Skin)
	var offsetValue = skin.BulletOffsets[enemy.ViewingDirection.Name]
	result.position.Add(&offsetValue)
	return result
}

// Move moves the Bullet. Returns true if the Bullet reached it's destination
func (b *Bullet) Move(elapsedTimeInMs int64) bool {
	if b.distanceMoved < b.distanceToTarget {
		var maxDistance = b.distanceToTarget - b.distanceMoved
		b.distanceMoved += geometry.CalculateMovementDistance(elapsedTimeInMs, b.speed, &maxDistance)
		b.position.X += geometry.CalculateMovementX(elapsedTimeInMs, b.direction, b.speed, &maxDistance)
		b.position.Y += geometry.CalculateMovementY(elapsedTimeInMs, b.direction, b.speed, &maxDistance)
	}
	return b.distanceMoved >= b.distanceToTarget
}

func (b *Bullet) Position() *geometry.Rectangle {
	return b.position
}
