package engine

import (
	"math"
	"retro-carnage/engine/geometry"
)

const (
	durationOfExplosion = DurationOfExplosionFrame * NumberOfExplosionSprites
)

// SomethingThatExplodes includes basically anything that has a position. BOOM!
type SomethingThatExplodes interface {
	Position() *geometry.Rectangle
}

// Explosion is a representation of an Explosive that did it's job.
type Explosion struct {
	causedByPlayer bool
	duration       int64
	hasMark        bool
	playerIdx      int
	Position       *geometry.Rectangle
	SpriteSupplier *ExplosionSpriteSupplier
}

// NewExplosion creates and initializes a new Explosion.
func NewExplosion(causedByPlayer bool, playerIdx int, explosive SomethingThatExplodes) *Explosion {
	return &Explosion{
		causedByPlayer: causedByPlayer,
		duration:       0,
		hasMark:        false,
		playerIdx:      playerIdx,
		Position: &geometry.Rectangle{
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
		Position:       e.Position.Clone(),
		SpriteSupplier: &BurnMarkSpriteSupplier{},
	}
	e.hasMark = true
	return &result
}
