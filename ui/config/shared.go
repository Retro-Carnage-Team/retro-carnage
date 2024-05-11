package config

import (
	"retro-carnage/ui/common"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
)

const (
	buttonPadding        = 15
	headlineDistanceLeft = 75
	headlineDistanceTop  = 75
	txtAudioSettings     = "AUDIO SETTINGS"
	txtBack              = "BACK"
	txtHeadlineOptions   = "OPTIONS"
	txtInputSettings     = "INPUT SETTINGS"
	txtVideoSettings     = "VIDEO SETTINGS"
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

func drawTextSelectionRect(w *pixelgl.Window, txt *text.Text) {
	drawSelectionRect(
		w,
		txt.Bounds().Min.X-buttonPadding,
		txt.Bounds().Min.Y-buttonPadding,
		txt.Bounds().Max.X+buttonPadding,
		txt.Bounds().Max.Y+buttonPadding,
	)
}
