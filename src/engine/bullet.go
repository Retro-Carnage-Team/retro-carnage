package engine

import (
	"retro-carnage/assets"
	"retro-carnage/engine/geometry"
	"retro-carnage/logging"
)

const (
	BulletHeight = 5
	BulletWidth  = 5
)

var (
	bulletOffsetForPlayer0 = buildBulletOffsetForPlayer0()
	bulletOffsetForPlayer1 = buildBulletOffsetForPlayer1()
)

// TODO: Bullet offsets should not be hardcoded.
// They depend on the player / enemy skin. So there should be two factories (player / enemy). These factories should get
// their values from the skin.
func buildBulletOffsetForPlayer0() (result map[geometry.Direction]*geometry.Point) {
	result = make(map[geometry.Direction]*geometry.Point)
	result[geometry.Up] = &geometry.Point{X: 80, Y: -BulletHeight}
	result[geometry.UpRight] = &geometry.Point{X: 105, Y: 0}
	result[geometry.Right] = &geometry.Point{X: 126, Y: 43}
	result[geometry.DownRight] = &geometry.Point{X: 103, Y: 100}
	result[geometry.Down] = &geometry.Point{X: 14, Y: 185}
	result[geometry.DownLeft] = &geometry.Point{X: -20, Y: 70}
	result[geometry.Left] = &geometry.Point{X: -15, Y: 3}
	result[geometry.UpLeft] = &geometry.Point{X: -BulletWidth, Y: -20}
	return
}

func buildBulletOffsetForPlayer1() (result map[geometry.Direction]*geometry.Point) {
	result = make(map[geometry.Direction]*geometry.Point)
	result[geometry.Up] = &geometry.Point{X: 87, Y: -(BulletHeight + 40)}
	result[geometry.UpRight] = &geometry.Point{X: 116, Y: 9}
	result[geometry.Right] = &geometry.Point{X: 135, Y: 52}
	result[geometry.DownRight] = &geometry.Point{X: 98, Y: 100}
	result[geometry.Down] = &geometry.Point{X: 16, Y: 160}
	result[geometry.DownLeft] = &geometry.Point{X: -(BulletWidth + 20), Y: 57}
	result[geometry.Left] = &geometry.Point{X: -BulletWidth, Y: 12}
	result[geometry.UpLeft] = &geometry.Point{X: -BulletWidth, Y: -13}
	return
}

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
	playerPosition geometry.Rectangle,
	direction geometry.Direction,
	selectedWeapon assets.Weapon,
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

	var offsets = bulletOffsetForPlayer0
	if 1 == b.playerIdx {
		offsets = bulletOffsetForPlayer1
	}
	var offsetValue = offsets[b.direction]
	if nil != offsetValue {
		b.Position.Add(offsetValue)
	}
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
