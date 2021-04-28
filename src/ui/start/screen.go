// Package start contains the second screen shown to the user. The screen displays a copyright notice and the dedication
// lines. Once the screen has been loaded, the theme song gets buffered. The next screen gets displayed when the theme
// song has been fully buffered.
package start

import (
	"github.com/faiface/pixel/pixelgl"
	"retro-carnage/assets"
	"retro-carnage/engine/input"
	"retro-carnage/ui/common"
	"retro-carnage/ui/common/fonts"
)

const txtFirstLine = "RETRO CARNAGE"
const txtSecondLine = "(C) 2020 THOMAS WERNER"
const txtThirdLine = "Dedicated to Jonathan Werner"
const txtFourthLine = "Inspired by 'DOGS OF WAR'"
const txtFifthLine = "(C) 1989 by Elite Systems Ltd."

type Screen struct {
	screenChangeRequired common.ScreenChangeCallback
	screenChangeTimeout  int64
	stereo               *assets.Stereo
	themeLoaded          bool
	window               *pixelgl.Window
}

func (s *Screen) SetInputController(_ input.Controller) {}

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
		// TODO: Buffer theme song asynchronously
		// This next call will buffer the song if it's not buffered already. That blocks the main thread for a couple of
		// seconds. It would be much cooler to start the buffer process asynchronously. Access to the map used to store
		// the songs could be protected by a Mutex. Then the buffering would be triggered once. Then we'd check if the
		// buffering has finished by accessing the map to see if the song it there.
		s.stereo.PlaySong(assets.ThemeSong)
		s.themeLoaded = true
	}
}

func (s *Screen) TearDown() {}

func (s *Screen) String() string {
	return string(common.Start)
}

func (s *Screen) renderScreen() {
	var renderer = fonts.TextRenderer{Window: s.window}
	renderer.DrawLineToScreenCenter(txtFirstLine, 4, common.Red)
	renderer.DrawLineToScreenCenter(txtSecondLine, 2.8, common.Yellow)
	renderer.DrawLineToScreenCenter(txtThirdLine, 0, common.Green)
	renderer.DrawLineToScreenCenter(txtFourthLine, -2.5, common.Green)
	renderer.DrawLineToScreenCenter(txtFifthLine, -3.7, common.Green)
}
