package engine

import "retro-carnage/engine/geometry"

const (
	ExplosionHitRectHeight = 200
	ExplosionHitRectWidth  = 200
	PlayerHitRectHeight    = 150
	PlayerHitRectWidth     = 90
	ScreenSize             = 1_500
	ScrollBarrierUp        = 1_000
	ScrollBarrierLeft      = 1_000
	ScrollBarrierRight     = 500
	ScrollMovementPerMs    = 0.3 // Screen.width = 1.500 px / 5.000 milliseconds = 0.2 px / ms
)

var (
	screenRect = &geometry.Rectangle{
		X:      0,
		Y:      0,
		Width:  ScreenSize,
		Height: ScreenSize,
	}
)
