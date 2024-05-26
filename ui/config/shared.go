package config

import (
	"image/color"
	"retro-carnage/ui/common"

	pixel "github.com/Retro-Carnage-Team/pixel2"
	"github.com/Retro-Carnage-Team/pixel2/backends/opengl"
	"github.com/Retro-Carnage-Team/pixel2/ext/imdraw"
)

const (
	buttonPadding        = 15
	headlineDistanceLeft = 75
	headlineDistanceTop  = 75
	txtActionFire        = "ACTION / FIRE"
	txtAudioSettings     = "AUDIO SETTINGS"
	txtBack              = "BACK"
	txtController        = "CONTROLLER"
	txtDecrease          = "-"
	txtDigitalAxis       = "DIGITAL AXIS"
	txtEffects           = "SOUND EFFECTS"
	txtFiveDigits        = "55555"
	txtFullScreen        = "FULLSCREEN"
	txtHeadlineOptions   = "OPTIONS"
	txtIncrease          = "+"
	txtInputSettings     = "INPUT SETTINGS"
	txtInputSettingsP1   = "INPUT SETTINGS PLAYER 1"
	txtInputSettingsP2   = "INPUT SETTINGS PLAYER 2"
	txtKeyboard          = "KEYBOARD"
	txtMonitor           = "MONITOR"
	txtMusic             = "MUSIC"
	txtNextWeapon        = "NEXT WEAPON"
	txtNotConfigured     = "NOT CONFIGURED"
	txtOtherMonitor      = "USE OTHER MONITOR: "
	txtPlayer1           = "PLAYER 1"
	txtPlayer2           = "PLAYER 2"
	txtPrevWeapon        = "PREVIOUS WEAPON"
	txtPressButton       = "PRESS BUTTON TO SELECT"
	txtPrimaryMonitor    = "USE PRIMARY MONITOR"
	txtRestartRequired   = "RESTART GAME TO APPLY CHANGED SETTINGS"
	txtSave              = "SAVE"
	txtScreenmode        = "SCREEN MODE"
	txtSelection         = "x"
	txtVideoSettings     = "VIDEO SETTINGS"
	txtWindowed          = "WINDOWED"
	txtWindowSize        = "WINDOW SIZE"
)

func drawRect(w *opengl.Window, left float64, bottom float64, right float64, top float64, col color.RGBA) {
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

func drawPossibleSelectionRect(w *opengl.Window, txtRect pixel.Rect) {
	drawRect(
		w,
		txtRect.Min.X-buttonPadding,
		txtRect.Min.Y,
		txtRect.Max.X+buttonPadding,
		txtRect.Max.Y+buttonPadding,
		common.LightGray,
	)
}

func drawTextSelectionRect(w *opengl.Window, txtRect pixel.Rect) {
	drawRect(
		w,
		txtRect.Min.X-buttonPadding,
		txtRect.Min.Y,
		txtRect.Max.X+buttonPadding,
		txtRect.Max.Y+buttonPadding,
		common.Yellow,
	)
}
