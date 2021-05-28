package engine

import (
	"retro-carnage/assets"
	"retro-carnage/engine/geometry"
	"retro-carnage/logging"
)

type ExplosiveGrenade struct {
	*Explosive
}

// NewExplosiveGrenadeByPlayer builds an ExplosiveGrenade objects for Grenades thrown by players.
func NewExplosiveGrenadeByPlayer(
	playerIdx int,
	playerPosition *geometry.Rectangle,
	direction geometry.Direction,
	selectedWeapon *assets.Grenade,
) *ExplosiveGrenade {
	var expGrenade = newExplosiveGrenade(playerPosition, direction, selectedWeapon)
	expGrenade.FiredByPlayer = true
	expGrenade.FiredByPlayerIdx = playerIdx
	return expGrenade
}

// NewExplosiveGrenadeByEnemy builds an ExplosiveGrenade objects for Grenades thrown by enemies.
func NewExplosiveGrenadeByEnemy(
	enemyPosition *geometry.Rectangle,
	direction geometry.Direction,
) *ExplosiveGrenade {
	return newExplosiveGrenade(enemyPosition, direction, assets.GrenadeCrate.GetByName("DM41"))
}

func newExplosiveGrenade(
	attackerPosition *geometry.Rectangle,
	direction geometry.Direction,
	selectedWeapon *assets.Grenade,
) *ExplosiveGrenade {
	var offset = grenadeOffsets(direction)
	return &ExplosiveGrenade{
		Explosive: &Explosive{
			distanceMoved:    0,
			distanceToTarget: float64(selectedWeapon.MovementDistance),
			direction:        direction,
			FiredByPlayer:    false,
			FiredByPlayerIdx: -1,
			position: &geometry.Rectangle{
				X:      attackerPosition.X + offset.X,
				Y:      attackerPosition.Y + offset.Y,
				Width:  GrenadeWidth,
				Height: GrenadeHeight,
			},
			speed:             selectedWeapon.MovementSpeed,
			SpriteSupplier:    &GrenadeSpriteSupplier{},
			ExplodesOnContact: false,
		},
	}
}

func grenadeOffsets(direction geometry.Direction) geometry.Point {
	switch {
	case direction == geometry.Up:
		return geometry.Point{X: 45, Y: -GrenadeHeight}
	case direction == geometry.UpRight:
		return geometry.Point{X: 45, Y: -GrenadeHeight}
	case direction == geometry.Right:
		return geometry.Point{X: 90, Y: 100}
	case direction == geometry.DownRight:
		return geometry.Point{X: 90, Y: 100}
	case direction == geometry.Down:
		return geometry.Point{X: 45, Y: 200}
	case direction == geometry.DownLeft:
		return geometry.Point{X: -GrenadeWidth, Y: 100}
	case direction == geometry.Left:
		return geometry.Point{X: -GrenadeWidth, Y: 100}
	case direction == geometry.UpLeft:
		return geometry.Point{X: 0, Y: -GrenadeHeight}
	default:
		logging.Error.Fatalf("no grenadeOffset for direction: %s", direction.Name)
		return geometry.Point{}
	}
}
