// Package start contains the second screen shown to the user. The screen displays a copyright notice and the dedication
// lines. Once the screen has been loaded, the theme song gets buffered. The next screen gets displayed when the theme
// song has been fully buffered.
package start

import (
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
	controller *controller
	window     *opengl.Window
}

func NewScreen() *Screen {
	var result = Screen{
		controller: NewController(),
	}
	return &result
}

func (s *Screen) SetInputController(_ input.InputController) {
	// This screen doesn't handle user input. Therefor no implementation required.
}

func (s *Screen) SetScreenChangeCallback(callback common.ScreenChangeCallback) {
	s.controller.SetScreenChangeCallback(callback)
}

func (s *Screen) SetWindow(window *opengl.Window) {
	s.window = window
}

func (s *Screen) SetUp() {
	// No set up action required
}

func (s *Screen) Update(elapsedTimeInMs int64) {
	s.controller.Update(elapsedTimeInMs)
	s.renderScreen()
}

func (s *Screen) TearDown() {
	// No tear down action required
}

func (s *Screen) String() string {
	return string(common.Start)
}

func (s *Screen) renderScreen() {
	var renderer = fonts.TextRenderer{Window: s.window}
	renderer.DrawLineToScreenCenter(txtFirstLine, 4, common.Red)
	renderer.DrawLineToScreenCenter(txtSecondLine, 2.8, common.Yellow)
	renderer.DrawLineToScreenCenter(txtThirdLine, 0, common.Green)
}
