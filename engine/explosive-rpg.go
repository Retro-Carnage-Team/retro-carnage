package engine

import (
	"retro-carnage/assets"
	"retro-carnage/engine/characters"
	"retro-carnage/engine/geometry"
	"retro-carnage/engine/graphics"
)

const (
	rpgHeight = 10
	rpgWidth  = 10
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
			distanceMoved:     0,
			distanceToTarget:  float64(selectedRpg.BulletRange),
			direction:         direction,
			firedByPlayer:     true,
			playerIdx:         playerIdx,
			position:          &geometry.Rectangle{X: playerPosition.X + offset.X, Y: playerPosition.Y + offset.Y, Width: rpgWidth, Height: rpgHeight},
			speed:             selectedRpg.BulletSpeed,
			SpriteSupplier:    graphics.NewRpgSpriteSupplier(direction),
			ExplodesOnContact: true,
		},
	}
}
