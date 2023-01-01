// Package buy_your_weapons just shows a basic text animation.
package buy_your_weapons

import (
	"fmt"
	"github.com/faiface/pixel/pixelgl"
	"retro-carnage/engine/input"
	"retro-carnage/ui/common"
	"retro-carnage/ui/common/fonts"
)

const timeAfterLastChar = 500
const timeBetweenChars = 120

type Screen struct {
	millisecondsPassed   int64
	PlayerIdx            int
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
	if s.textLength < 25 {
		if s.millisecondsPassed >= timeBetweenChars {
			s.textLength++
			s.text = s.getFullText()[:s.textLength]
			s.millisecondsPassed = 0
		}
	} else if s.millisecondsPassed >= timeAfterLastChar {
		if 0 == s.PlayerIdx {
			s.screenChangeRequired(common.ShopP1)
		} else {
			s.screenChangeRequired(common.ShopP2)
		}
	}
	var renderer = fonts.TextRenderer{Window: s.window}
	renderer.DrawLineToScreenCenter(s.text, 0, common.White)
}

func (s *Screen) TearDown() {}

func (s *Screen) String() string {
	if 0 == s.PlayerIdx {
		return string(common.BuyYourWeaponsP1)
	}
	return string(common.BuyYourWeaponsP2)
}

func (s *Screen) getFullText() string {
	return fmt.Sprintf("BUY YOUR WEAPONS PLAYER %d", s.PlayerIdx+1)
}
