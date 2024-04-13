package lets_begin

import (
	"retro-carnage/assets"
	"retro-carnage/engine/input"
	"retro-carnage/ui/common"
	"retro-carnage/ui/common/fonts"

	"github.com/faiface/pixel/pixelgl"
)

const (
	displayText             = "LET THE MISSION BEGIN"
	timeAfterLastChar       = 500
	timeBetweenChars        = 120
	timeBetweenVolumeChange = 150
)

type Screen struct {
	characterTimer       int64
	screenChangeRequired common.ScreenChangeCallback
	stereo               *assets.Stereo
	text                 string
	textLength           int
	volumeTimer          int64
	window               *pixelgl.Window
}

func (s *Screen) SetInputController(_ input.InputController) {
	// Screen doesn't process user input. So no implementation required.
}

func (s *Screen) SetScreenChangeCallback(callback common.ScreenChangeCallback) {
	s.screenChangeRequired = callback
}

func (s *Screen) SetWindow(window *pixelgl.Window) {
	s.window = window
}

func (s *Screen) SetUp() {
	s.stereo = assets.NewStereo()
}

func (s *Screen) Update(elapsedTimeInMs int64) {
	s.characterTimer += elapsedTimeInMs
	s.volumeTimer += elapsedTimeInMs
	if s.textLength < len(displayText) {
		if s.characterTimer >= timeBetweenChars {
			s.textLength++
			s.text = displayText[:s.textLength]
			s.characterTimer = 0
		}
		if s.volumeTimer >= timeBetweenVolumeChange {
			s.stereo.DecreaseVolume(assets.ThemeSong)
			s.volumeTimer = 0
		}
	} else if s.characterTimer >= timeAfterLastChar {
		s.screenChangeRequired(common.Game)
		s.stereo.StopSong(assets.ThemeSong)
	}
	var renderer = fonts.TextRenderer{Window: s.window}
	renderer.DrawLineToScreenCenter(s.text, 0, common.White)
}

func (s *Screen) TearDown() {
	// No tear down action required
}

func (s *Screen) String() string {
	return string(common.LetTheMissionBegin)
}
