package engine

import (
	"retro-carnage/assets"
	"retro-carnage/engine/characters"
	"retro-carnage/engine/geometry"
	"retro-carnage/engine/graphics"

	"math/rand"
)

const (
	BulletHeight          = 5
	BulletWidth           = 5
	EnemyBulletSpeed      = 1.4
	EnemyBulletRange      = 500
	ExplosiveBulletHeight = 8
	ExplosiveBulletWidth  = 8
	ShotHeight            = 3
	ShotWidth             = 3
)

var (
	bulletPatternsScattering [][]float64 = [][]float64{
		{-0.11071, -0.09508, +0.01545, +0.01715, +0.03318},
		{-0.11265, -0.05904, -0.03388, +0.00859, +0.11847},
		{-0.16909, -0.10517, -0.09701, +0.02297, +0.17806},
		{-0.15711, -0.10899, +0.00952, +0.06921, +0.12005},
		{-0.14312, -0.09084, -0.04467, +0.01118, +0.06204},
	}
)

// Bullet is a projectile that has been fired by a player or enemy.
type Bullet struct {
	distanceMoved               float64
	distanceToTarget            float64
	direction                   geometry.Direction
	directionDeviationInRadians float64
	explodes                    bool
	firedByPlayer               bool
	playerIdx                   int
	position                    *geometry.Rectangle
	speed                       float64
}

// NewBulletFiredByPlayer creates and returns a slice with one or more instances of Bullet.
func NewBulletFiredByPlayer(
	playerIdx int,
	playerPosition *geometry.Rectangle,
	direction geometry.Direction,
	selectedWeapon *assets.Weapon,
) (result []*Bullet) {
	result = make([]*Bullet, 0)

	var ammo = assets.AmmunitionCrate.GetByName(selectedWeapon.Ammo)
	bulletPattern := getBulletPattern(ammo)
	bulletWidth, bulletHeight := getBulletDimension(ammo)

	for i := 0; i < len(bulletPattern); i++ {
		var bullet = &Bullet{
			distanceMoved:               0,
			distanceToTarget:            float64(selectedWeapon.BulletRange),
			direction:                   direction,
			directionDeviationInRadians: bulletPattern[i],
			explodes:                    ammo.Explosive,
			firedByPlayer:               true,
			playerIdx:                   playerIdx,
			position:                    &geometry.Rectangle{X: playerPosition.X, Y: playerPosition.Y, Width: bulletHeight, Height: bulletWidth},
			speed:                       selectedWeapon.BulletSpeed,
		}
		var offsetValue = graphics.SkinForPlayer(playerIdx).BulletOffsets[direction.Name]
		bullet.position.Add(&offsetValue)
		result = append(result, bullet)
	}

	return result
}

func getBulletPattern(ammo *assets.Ammunition) []float64 {
	var bulletPattern []float64 = []float64{0}
	if ammo.Scattering {
		// get a random scattering pattern (for shotguns)
		bulletPattern = bulletPatternsScattering[rand.Intn(len(bulletPatternsScattering))]
	}
	return bulletPattern
}

func getBulletDimension(ammo *assets.Ammunition) (float64, float64) {
	if ammo.Explosive {
		return ExplosiveBulletWidth, ExplosiveBulletHeight
	}
	if ammo.Scattering {
		return ShotWidth, ShotHeight
	}
	return BulletWidth, BulletHeight
}

// NewBulletFiredByEnemy creates and returns a new instance of Bullet.
func NewBulletFiredByEnemy(enemy *characters.ActiveEnemy) (result *Bullet) {
	result = &Bullet{
		distanceMoved:               0,
		distanceToTarget:            float64(EnemyBulletRange),
		direction:                   *enemy.ViewingDirection,
		directionDeviationInRadians: 0,
		firedByPlayer:               false,
		position:                    &geometry.Rectangle{X: enemy.Position().X, Y: enemy.Position().Y, Width: BulletWidth, Height: BulletHeight},
		speed:                       EnemyBulletSpeed,
	}

	var skin = graphics.GetEnemySkin(enemy.Skin)
	var offsetValue = skin.BulletOffsets[enemy.ViewingDirection.Name]
	result.position.Add(&offsetValue)
	return result
}

// Move moves the Bullet. Returns true if the Bullet reached it's destination
func (b *Bullet) Move(elapsedTimeInMs int64) bool {
	if b.distanceMoved < b.distanceToTarget {
		var maxDistance = b.distanceToTarget - b.distanceMoved
		b.distanceMoved += geometry.Move(b.position, elapsedTimeInMs, b.direction, b.directionDeviationInRadians, b.speed, &maxDistance)
	}
	return b.distanceMoved >= b.distanceToTarget
}

func (b *Bullet) Position() *geometry.Rectangle {
	return b.position
}
