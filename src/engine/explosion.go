package engine

import (
	"math"
	"retro-carnage/engine/geometry"
)

type SomethingThatExplodes interface {
	Position() geometry.Rectangle
}

type Explosion struct {
	duration       int64
	causedByPlayer bool
	playerIdx      int
	position       geometry.Rectangle
	spriteSupplier *ExplosionSpriteSupplier
}

func NewExplosion(causedByPlayer bool, playerIdx int, explosive SomethingThatExplodes) *Explosion {
	return &Explosion{
		duration:       0,
		causedByPlayer: causedByPlayer,
		playerIdx:      playerIdx,
		position: geometry.Rectangle{
			X:      math.Round(explosive.Position().X + explosive.Position().Width/2 - ExplosionHitRectWidth/2),
			Y:      math.Round(explosive.Position().Y + explosive.Position().Height - ExplosionHitRectHeight),
			Width:  ExplosionHitRectWidth,
			Height: ExplosionHitRectHeight,
		},
		spriteSupplier: &ExplosionSpriteSupplier{},
	}
}
