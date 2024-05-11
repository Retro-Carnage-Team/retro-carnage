package config

import (
	"retro-carnage/ui/common"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

const (
	buttonPadding        = 15
	headlineDistanceLeft = 75
	headlineDistanceTop  = 75
)

func drawSelectionRect(w *pixelgl.Window, left float64, bottom float64, right float64, top float64) {
	imd := imdraw.New(nil)
	imd.Color = common.Yellow
	imd.EndShape = imdraw.RoundEndShape
	imd.Push(pixel.V(left, bottom), pixel.V(right, bottom))
	imd.Push(pixel.V(left, bottom), pixel.V(left, top))
	imd.Push(pixel.V(left, top), pixel.V(right, top))
	imd.Push(pixel.V(right, bottom), pixel.V(right, top))
	imd.Line(4)
	imd.Draw(w)
}
