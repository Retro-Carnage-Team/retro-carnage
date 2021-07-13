package engine

import (
	"retro-carnage/assets"
	"retro-carnage/engine/characters"
	"retro-carnage/engine/geometry"
	"retro-carnage/logging"
)

const (
	BulletHeight = 5
	BulletWidth  = 5
)

// Bullet is a projectile that has been fired by a player or enemy.
type Bullet struct {
	distanceMoved    float64
	distanceToTarget float64
	direction        geometry.Direction
	firedByPlayer    bool
	playerIdx        int
	Position         *geometry.Rectangle
	speed            float64
}

// NewBulletFiredByPlayer creates and returns a new instance of Bullet.
func NewBulletFiredByPlayer(
	playerIdx int,
	playerPosition *geometry.Rectangle,
	direction geometry.Direction,
	selectedWeapon *assets.Weapon,
) (result *Bullet) {
	result = &Bullet{
		distanceMoved:    0,
		distanceToTarget: float64(selectedWeapon.BulletRange),
		direction:        direction,
		firedByPlayer:    true,
		playerIdx:        playerIdx,
		Position: &geometry.Rectangle{
			X:      playerPosition.X,
			Y:      playerPosition.Y,
			Width:  BulletWidth,
			Height: BulletHeight,
		},
		speed: selectedWeapon.BulletSpeed,
	}
	result.applyPlayerOffset()
	return
}

func (b *Bullet) applyPlayerOffset() {
	if !b.firedByPlayer || b.playerIdx < 0 || b.playerIdx > 1 {
		logging.Error.Fatal("Failed to apply player offset due to invalid parameters")
	}

	var offsetValue = characters.SkinForPlayer(b.playerIdx).BulletOffsets[b.direction.Name]
	b.Position.Add(&offsetValue)
}

// Move moves the Bullet. Returns true if the Bullet reached it's destination
func (b *Bullet) Move(elapsedTimeInMs int64) bool {
	if b.distanceMoved < b.distanceToTarget {
		var maxDistance = b.distanceToTarget - b.distanceMoved
		b.distanceMoved += geometry.CalculateMovementDistance(elapsedTimeInMs, b.speed, &maxDistance)
		b.Position.X += geometry.CalculateMovementX(elapsedTimeInMs, b.direction, b.speed, &maxDistance)
		b.Position.Y += geometry.CalculateMovementY(elapsedTimeInMs, b.direction, b.speed, &maxDistance)
	}
	return b.distanceMoved >= b.distanceToTarget
}
