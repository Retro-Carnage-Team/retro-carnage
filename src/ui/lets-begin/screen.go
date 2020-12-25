package lets_begin

import (
	"github.com/faiface/pixel/pixelgl"
	"retro-carnage/engine/input"
	"retro-carnage/ui/common"
	"retro-carnage/ui/common/fonts"
)

const displayText = "LET THE MISSION BEGIN"
const timeAfterLastChar = 500
const timeBetweenChars = 120

type Screen struct {
	millisecondsPassed   int64
	screenChangeRequired common.ScreenChangeCallback
	text                 string
	textLength           int
	window               *pixelgl.Window
}

func (s *Screen) SetInputController(_ input.Controller) {}

func (s *Screen) SetScreenChangeCallback(callback common.ScreenChangeCallback) {
	s.screenChangeRequired = callback
}

func (s *Screen) SetWindow(window *pixelgl.Window) {
	s.window = window
}

func (s *Screen) SetUp() {}

func (s *Screen) Update(elapsedTimeInMs int64) {
	s.millisecondsPassed += elapsedTimeInMs
	if s.textLength < len(displayText) {
		if s.millisecondsPassed >= timeBetweenChars {
			s.textLength++
			s.text = displayText[:s.textLength]
			s.millisecondsPassed = 0
		}
	} else if s.millisecondsPassed >= timeAfterLastChar {
		s.screenChangeRequired(common.Game)
	}
	var renderer = fonts.TextRenderer{Window: s.window}
	renderer.DrawLineToScreenCenter(s.text, 0, common.White)
}

func (s *Screen) TearDown() {}

func (s *Screen) String() string {
	return string(common.LetTheMissionBegin)
}
