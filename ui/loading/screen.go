// Package loading contains the first screen shown to the user. This contains the message that the game is loading.
// Actually there's not much loading going on. We just play the sound of an Amiga 500 diskette drive loading a game.
package loading

import (
	"retro-carnage/engine/geometry"
	"retro-carnage/input"
	"retro-carnage/ui/common"
	"retro-carnage/ui/common/fonts"

	"github.com/Retro-Carnage-Team/pixel2/backends/opengl"
)

const (
	txtFirstLine  = "RETRO CARNAGE"
	txtSecondLine = "(C) 2020 THOMAS WERNER"
	txtThirdLine  = "Dedicated to Emma & Jonathan Werner"
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
	var renderer = fonts.TextRenderer{Window: s.window}
	renderer.DrawLineToScreenCenter(txtFirstLine, 4, common.Red)
	renderer.DrawLineToScreenCenter(txtSecondLine, 2.8, common.Yellow)
	renderer.DrawLineToScreenCenter(txtThirdLine, 0, common.Green)
}
