package highscore

import (
	"fmt"
	"math"
	"retro-carnage/assets"
	"retro-carnage/engine/input"
	"retro-carnage/ui/common"
	"retro-carnage/ui/common/fonts"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
)

const (
	title = "HIGH SCORES"
)

// Screen is the High Score table screen.
type Screen struct {
	inputController      input.InputController
	screenChangeRequired common.ScreenChangeCallback
	stereo               *assets.Stereo
	window               *pixelgl.Window
}

// SetInputController passes the input controller to the screen.
func (s *Screen) SetInputController(inputCtrl input.InputController) {
	s.inputController = inputCtrl
}

// SetScreenChangeCallback passes a callback function that cann be called to switch to another screen.
func (s *Screen) SetScreenChangeCallback(callback common.ScreenChangeCallback) {
	s.screenChangeRequired = callback
}

// SetWindow passes the application window to the Screen.
func (s *Screen) SetWindow(window *pixelgl.Window) {
	s.window = window
}

// SetUp initializes the Screen.
// This method gets called once before the Screen gets shown.
func (s *Screen) SetUp() {
	s.stereo = assets.NewStereo()
}

// Update gets called once during each rendering cycle.
// It can be used to draw the content of the Screen.
func (s *Screen) Update(_ int64) {
	if s.window.JustPressed(pixelgl.KeyEnter) || s.inputController.GetUiEventStateCombined().PressedButton {
		s.screenChangeRequired(common.Title)
	} else {
		s.drawTitle()
		s.drawTable()
	}
}

func (s *Screen) drawTitle() {
	var defaultFontSize = fonts.DefaultFontSize()
	var dimensions = fonts.GetTextDimension(defaultFontSize, title)
	var position = pixel.V((s.window.Bounds().W()-dimensions.X)/2, s.window.Bounds().H()-(3*dimensions.Y))
	var txt = text.New(position, fonts.SizeToFontAtlas[defaultFontSize])
	txt.Color = common.Green
	_, _ = fmt.Fprint(txt, title)
	txt.Draw(s.window, pixel.IM)
}

func (s *Screen) drawTable() {
	var defaultFontSize = fonts.DefaultFontSize()
	var maxLineWidth = 0.0
	var maxLineHeight = 0.0
	for i, entry := range EntryControllerInstance.entries {
		var lineDimensions = fonts.GetTextDimension(defaultFontSize, entry.ToString(i+1))
		maxLineHeight = math.Max(maxLineHeight, lineDimensions.Y)
		maxLineWidth = math.Max(maxLineWidth, lineDimensions.X)
	}

	var offset = s.window.Bounds().H() - (maxLineHeight * 6.5)
	var posLeft = (s.window.Bounds().W() - maxLineWidth) / 2
	for i, entry := range EntryControllerInstance.entries {
		var txt = text.New(pixel.V(posLeft, offset), fonts.SizeToFontAtlas[defaultFontSize])
		txt.Color = common.White
		_, _ = fmt.Fprint(txt, entry.ToString(i+1))
		txt.Draw(s.window, pixel.IM)
		offset -= maxLineHeight * 1.5
	}
}

// TearDown can be used as a life-cycle hook to release resources that a Screen blocked.
// It will be called once after the last Update.
func (s *Screen) TearDown() {
	s.stereo.StopSong(assets.GameOverSong)
	s.stereo.StopSong(assets.GameWonSong)
	s.stereo.PlaySong(assets.ThemeSong)
}

// String should return the ScreenName of the Screen
func (s *Screen) String() string {
	return string(common.HighScore)
}
