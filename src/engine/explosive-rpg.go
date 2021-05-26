package engine

import (
	"retro-carnage/assets"
	"retro-carnage/engine/geometry"
	"retro-carnage/logging"
)

type ExplosiveRpg struct {
	Explosive
}

// NewExplosiveRpg builds an ExplosiveRpg.
func NewExplosiveRpg(
	playerIdx int,
	playerPosition geometry.Rectangle,
	direction geometry.Direction,
	selectedRpg *assets.Grenade,
) *ExplosiveRpg {
	var offset = rpgOffset(playerIdx, direction)
	return &ExplosiveRpg{
		Explosive: Explosive{
			distanceMoved:    0,
			distanceToTarget: float64(selectedRpg.MovementDistance),
			direction:        direction,
			FiredByPlayer:    true,
			FiredByPlayerIdx: playerIdx,
			position: &geometry.Rectangle{
				X:      playerPosition.X + offset.X,
				Y:      playerPosition.Y + offset.Y,
				Width:  RpgWidth,
				Height: RpgHeight,
			},
			speed:             selectedRpg.MovementSpeed,
			SpriteSupplier:    NewRpgSpriteSupplier(direction),
			ExplodesOnContact: true,
		},
	}
}

func rpgOffset(playerIdx int, direction geometry.Direction) geometry.Point {
	if 0 == playerIdx {
		return offsetForPlayer0(direction)
	}
	return offsetForPlayer1(direction)
}

func offsetForPlayer0(direction geometry.Direction) geometry.Point {
	switch {
	case direction == geometry.Up:
		return geometry.Point{X: 80, Y: -RpgHeight}
	case direction == geometry.UpRight:
		return geometry.Point{X: 110, Y: -45}
	case direction == geometry.Right:
		return geometry.Point{X: 131, Y: 43}
	case direction == geometry.DownRight:
		return geometry.Point{X: 113, Y: 110}
	case direction == geometry.Down:
		return geometry.Point{X: 14, Y: 195}
	case direction == geometry.DownLeft:
		return geometry.Point{X: -120, Y: 80}
	case direction == geometry.Left:
		return geometry.Point{X: -25, Y: 3}
	case direction == geometry.UpLeft:
		return geometry.Point{X: -RpgWidth, Y: -20}
	default:
		logging.Error.Fatalf("no such offset for player 0 / direction: %s", direction.Name)
		return geometry.Point{}
	}
}

func offsetForPlayer1(direction geometry.Direction) geometry.Point {
	switch {
	case direction == geometry.Up:
		return geometry.Point{X: 87, Y: -(RpgHeight + 40)}
	case direction == geometry.UpRight:
		return geometry.Point{X: 126, Y: 9}
	case direction == geometry.Right:
		return geometry.Point{X: 145, Y: 52}
	case direction == geometry.DownRight:
		return geometry.Point{X: 108, Y: 110}
	case direction == geometry.Down:
		return geometry.Point{X: 16, Y: 170}
	case direction == geometry.DownLeft:
		return geometry.Point{X: -(RpgWidth + 20), Y: 67}
	case direction == geometry.Left:
		return geometry.Point{X: -RpgWidth, Y: 12}
	case direction == geometry.UpLeft:
		return geometry.Point{X: -RpgWidth, Y: -13}
	default:
		logging.Error.Fatalf("no such offset for player 0 / direction: %s", direction.Name)
		return geometry.Point{}
	}
}
