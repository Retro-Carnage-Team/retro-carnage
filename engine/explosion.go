package engine

import (
	"math"
	"retro-carnage/engine/geometry"
)

const (
	durationOfExplosion = DurationOfExplosionFrame * NumberOfExplosionSprites
)

// Explosion is a representation of an Explosive that did it's job.
type Explosion struct {
	causedByPlayer bool
	duration       int64
	hasMark        bool
	playerIdx      int
	position       *geometry.Rectangle
	SpriteSupplier *ExplosionSpriteSupplier
}

// NewExplosion creates and initializes a new Explosion.
func NewExplosion(causedByPlayer bool, playerIdx int, explosive geometry.Positioned) *Explosion {
	return &Explosion{
		causedByPlayer: causedByPlayer,
		duration:       0,
		hasMark:        false,
		playerIdx:      playerIdx,
		position: &geometry.Rectangle{
			X:      math.Round(explosive.Position().X + explosive.Position().Width/2 - ExplosionHitRectWidth/2),
			Y:      math.Round(explosive.Position().Y + explosive.Position().Height - ExplosionHitRectHeight),
			Width:  ExplosionHitRectWidth,
			Height: ExplosionHitRectHeight,
		},
		SpriteSupplier: &ExplosionSpriteSupplier{},
	}
}

// CreatesMark returns true if this Explosion should create a mark when rendered. This happens only once during the life
// span of the Explosion.
func (e *Explosion) CreatesMark() bool {
	return !e.hasMark && (e.duration >= durationOfExplosion/3)
}

// CreateMark creates a BurnMark caused by this explosion - and makes sure this can happen only once.
func (e *Explosion) CreateMark() *BurnMark {
	var result = BurnMark{
		position:       e.position.Clone(),
		SpriteSupplier: &BurnMarkSpriteSupplier{},
	}
	e.hasMark = true
	return &result
}

func (e *Explosion) Position() *geometry.Rectangle {
	return e.position
}
