package engine

import (
	"retro-carnage/assets"
	"retro-carnage/engine/geometry"
	"retro-carnage/engine/graphics"
	"retro-carnage/logging"
)

const (
	grenadeHeight          = 17
	grenadeWidth           = 32
	grenadeTypeUsedByEnemy = "M67"
)

// Something that can explode
type Explosive struct {
	distanceMoved     float64
	distanceToTarget  float64
	direction         geometry.Direction
	firedByPlayer     bool
	playerIdx         int
	position          *geometry.Rectangle
	speed             float64
	SpriteSupplier    graphics.ExplosiveSpriteSupplier
	ExplodesOnContact bool
}

// NewExplosiveGrenadeByPlayer builds an Explosive object for grenades thrown by players.
func NewExplosiveGrenadeByPlayer(
	playerIdx int,
	playerPosition *geometry.Rectangle,
	direction geometry.Direction,
	selectedWeapon *assets.Grenade,
) *Explosive {
	var expGrenade = newExplosiveGrenade(playerPosition, direction, selectedWeapon)
	expGrenade.firedByPlayer = true
	expGrenade.playerIdx = playerIdx
	return expGrenade
}

// NewExplosiveGrenadeByEnemy builds an Explosive object for grenades thrown by enemies.
func NewExplosiveGrenadeByEnemy(
	enemyPosition *geometry.Rectangle,
	direction geometry.Direction,
) *Explosive {
	return newExplosiveGrenade(enemyPosition, direction, assets.GrenadeCrate.GetByName(grenadeTypeUsedByEnemy))
}

// Move updates the position of the explosive on screen.
// Returns true if the explosive reached it's destination
func (e *Explosive) Move(elapsedTimeInMs int64) bool {
	if e.distanceMoved < e.distanceToTarget {
		var maxDistance = e.distanceToTarget - e.distanceMoved
		e.distanceMoved += geometry.Move(e.position, elapsedTimeInMs, e.direction, 0, e.speed, &maxDistance)
	}
	return e.distanceMoved >= e.distanceToTarget
}

func (e *Explosive) Position() *geometry.Rectangle {
	return e.position
}

func grenadeOffsets(direction geometry.Direction) geometry.Point {
	switch {
	case direction == geometry.Up:
		return geometry.Point{X: 45, Y: -grenadeHeight}
	case direction == geometry.UpRight:
		return geometry.Point{X: 45, Y: -grenadeHeight}
	case direction == geometry.Right:
		return geometry.Point{X: 90, Y: 100}
	case direction == geometry.DownRight:
		return geometry.Point{X: 90, Y: 100}
	case direction == geometry.Down:
		return geometry.Point{X: 45, Y: 200}
	case direction == geometry.DownLeft:
		return geometry.Point{X: -grenadeWidth, Y: 100}
	case direction == geometry.Left:
		return geometry.Point{X: -grenadeWidth, Y: 100}
	case direction == geometry.UpLeft:
		return geometry.Point{X: 0, Y: -grenadeHeight}
	default:
		logging.Error.Fatalf("no grenadeOffset for direction: %s", direction.Name)
		return geometry.Point{}
	}
}

func newExplosiveGrenade(
	attackerPosition *geometry.Rectangle,
	direction geometry.Direction,
	selectedWeapon *assets.Grenade,
) *Explosive {
	var offset = grenadeOffsets(direction)
	return &Explosive{
		distanceMoved:     0,
		distanceToTarget:  float64(selectedWeapon.MovementDistance),
		direction:         direction,
		firedByPlayer:     false,
		playerIdx:         -1,
		position:          &geometry.Rectangle{X: attackerPosition.X + offset.X, Y: attackerPosition.Y + offset.Y, Width: grenadeWidth, Height: grenadeHeight},
		speed:             selectedWeapon.MovementSpeed,
		SpriteSupplier:    &graphics.GrenadeSpriteSupplier{},
		ExplodesOnContact: false,
	}
}
