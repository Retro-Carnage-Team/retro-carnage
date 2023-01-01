package engine

import (
	"retro-carnage/assets"
	"retro-carnage/engine/characters"
	"retro-carnage/engine/geometry"
)

// ExplosiveRpg is a RPG (rocket propelled grenade) flying across the screen.
type ExplosiveRpg struct {
	*Explosive
}

// NewExplosiveRpg builds a new ExplosiveRpg.
func NewExplosiveRpg(
	playerIdx int,
	playerPosition *geometry.Rectangle,
	direction geometry.Direction,
	selectedRpg *assets.Weapon,
) *ExplosiveRpg {
	var offset = characters.SkinForPlayer(playerIdx).BulletOffsets[direction.Name]
	return &ExplosiveRpg{
		Explosive: &Explosive{
			distanceMoved:    0,
			distanceToTarget: float64(selectedRpg.BulletRange),
			direction:        direction,
			FiredByPlayer:    true,
			FiredByPlayerIdx: playerIdx,
			position: &geometry.Rectangle{
				X:      playerPosition.X + offset.X,
				Y:      playerPosition.Y + offset.Y,
				Width:  RpgWidth,
				Height: RpgHeight,
			},
			speed:             selectedRpg.BulletSpeed,
			SpriteSupplier:    NewRpgSpriteSupplier(direction),
			ExplodesOnContact: true,
		},
	}
}
