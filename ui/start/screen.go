// Package start contains the second screen shown to the user. The screen displays a copyright notice and the dedication
// lines. Once the screen has been loaded, the theme song gets buffered. The next screen gets displayed when the theme
// song has been fully buffered.
package start

import (
	"retro-carnage/assets"
	"retro-carnage/engine/input"
	"retro-carnage/ui/common"
	"retro-carnage/ui/common/fonts"

	"github.com/faiface/pixel/pixelgl"
)

const (
	txtFirstLine  = "RETRO CARNAGE"
	txtSecondLine = "(C) 2020 THOMAS WERNER"
	txtThirdLine  = "Dedicated to Emma & Jonathan Werner"
)

type Screen struct {
	screenChangeRequired common.ScreenChangeCallback
	screenChangeTimeout  int64
	stereo               *assets.Stereo
	themeLoaded          bool
	window               *pixelgl.Window
}

func (s *Screen) SetInputController(_ input.InputController) {
	// This screen doesn't handle user input. Therefor no implementation required.
}

func (s *Screen) SetScreenChangeCallback(callback common.ScreenChangeCallback) {
	s.screenChangeRequired = callback
}

func (s *Screen) SetWindow(window *pixelgl.Window) {
	s.window = window
}

func (s *Screen) SetUp() {
	s.screenChangeTimeout = 0
	s.stereo = assets.NewStereo()
	s.themeLoaded = false

	common.StartScreenInit()
}

func (s *Screen) Update(elapsedTimeInMs int64) {
	s.screenChangeTimeout += elapsedTimeInMs
	if s.themeLoaded {
		s.screenChangeRequired(common.Title)
	}
	s.renderScreen()
	if !s.themeLoaded && (s.screenChangeTimeout > 100) {
		s.stereo.PlaySong(assets.ThemeSong)
		s.themeLoaded = true
	}
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
