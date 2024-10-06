// Package loading contains the first screen shown to the user. This contains the message that the game is loading.
// Actually there's not much loading going on. We just play the sound of an Amiga 500 diskette drive loading a game.
package loading

import (
	"fmt"
	"retro-carnage/engine/geometry"
	"retro-carnage/input"
	"retro-carnage/ui/common"
	"retro-carnage/ui/common/fonts"

	pixel "github.com/Retro-Carnage-Team/pixel2"
	"github.com/Retro-Carnage-Team/pixel2/backends/opengl"
	"github.com/Retro-Carnage-Team/pixel2/ext/text"
	"golang.org/x/image/colornames"
)

const (
	txtFirstLine  = "RETRO CARNAGE"
	txtSecondLine = "IS LOADING"
)

type Screen struct {
	controller     *controller
	textDimensions map[string]*geometry.Point
	window         *opengl.Window
}

func NewScreen() *Screen {
	var result = Screen{
		controller: newController(),
	}
	return &result
}

func (s *Screen) SetInputController(_ input.InputController) {
	// The screen doesn't require user input. Therefor no implementation of this method.
}

func (s *Screen) SetScreenChangeCallback(callback common.ScreenChangeCallback) {
	s.controller.setScreenChangeCallback(callback)
}

func (s *Screen) SetWindow(window *opengl.Window) {
	s.window = window
}

func (s *Screen) SetUp() {
	s.textDimensions = fonts.GetTextDimensions(fonts.DefaultFontSize(), txtFirstLine, txtSecondLine)
}

func (s *Screen) Update(elapsedTimeInMs int64) {
	s.controller.update(elapsedTimeInMs)
	s.drawScreen()
}

func (s *Screen) TearDown() {
	// No tear down action required
}

func (s *Screen) String() string {
	return string(common.Loading)
}

func (s *Screen) drawScreen() {
	var firstLineDimensions = s.textDimensions[txtFirstLine]
	var firstLineX = (s.window.Bounds().Max.X - firstLineDimensions.X) / 2
	var firstLineY = (s.window.Bounds().Max.Y-(3*firstLineDimensions.Y))/2 + firstLineDimensions.Y*1.5

	var secondLineDimensions = s.textDimensions[txtSecondLine]
	var secondLineX = (s.window.Bounds().Max.X - secondLineDimensions.X) / 2
	var secondLineY = (s.window.Bounds().Max.Y - (3 * secondLineDimensions.Y)) / 2

	var defaultFontSize = fonts.DefaultFontSize()
	var txt = text.New(pixel.V(firstLineX, firstLineY), fonts.SizeToFontAtlas[defaultFontSize])
	_, _ = fmt.Fprint(txt, txtFirstLine)
	txt.Color = colornames.Red
	txt.Draw(s.window, pixel.IM)

	txt = text.New(pixel.V(secondLineX, secondLineY), fonts.SizeToFontAtlas[defaultFontSize])
	_, _ = fmt.Fprint(txt, txtSecondLine)
	txt.Color = colornames.Red
	txt.Draw(s.window, pixel.IM)
}
