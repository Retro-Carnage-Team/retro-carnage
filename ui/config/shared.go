package config

import (
	"image/color"
	"retro-carnage/ui/common"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

const (
	buttonPadding        = 15
	headlineDistanceLeft = 75
	headlineDistanceTop  = 75
	txtAudioSettings     = "AUDIO SETTINGS"
	txtBack              = "BACK"
	txtDecrease          = "-"
	txtFullScreen        = "FULLSCREEN"
	txtHeadlineOptions   = "OPTIONS"
	txtIncrease          = "+"
	txtInputSettings     = "INPUT SETTINGS"
	txtMonitor           = "MONITOR"
	txtOtherMonitor      = "USE OTHER MONITOR: "
	txtPrimaryMonitor    = "USE PRIMARY MONITOR"
	txtSave              = "SAVE"
	txtScreenmode        = "SCREEN MODE"
	txtSelection         = "x"
	txtVideoSettings     = "VIDEO SETTINGS"
	txtWindowed          = "WINDOWED"
	txtWindowSize        = "WINDOW SIZE"
)

func drawRect(w *pixelgl.Window, left float64, bottom float64, right float64, top float64, col color.RGBA) {
	imd := imdraw.New(nil)
	imd.Color = col
	imd.EndShape = imdraw.RoundEndShape
	imd.Push(pixel.V(left, bottom), pixel.V(right, bottom))
	imd.Push(pixel.V(left, bottom), pixel.V(left, top))
	imd.Push(pixel.V(left, top), pixel.V(right, top))
	imd.Push(pixel.V(right, bottom), pixel.V(right, top))
	imd.Line(4)
	imd.Draw(w)
}

func drawSelectionRect(w *pixelgl.Window, left float64, bottom float64, right float64, top float64) {
	drawRect(w, left, bottom, right, top, common.Yellow)
}

func drawPossibleSelectionRect(w *pixelgl.Window, txtRect pixel.Rect) {
	drawRect(
		w,
		txtRect.Min.X-buttonPadding,
		txtRect.Min.Y-buttonPadding,
		txtRect.Max.X+buttonPadding,
		txtRect.Max.Y+buttonPadding,
		common.LightGray,
	)
}

func drawTextSelectionRect(w *pixelgl.Window, txtRect pixel.Rect) {
	drawRect(
		w,
		txtRect.Min.X-buttonPadding,
		txtRect.Min.Y-buttonPadding,
		txtRect.Max.X+buttonPadding,
		txtRect.Max.Y+buttonPadding,
		common.Yellow,
	)
}
