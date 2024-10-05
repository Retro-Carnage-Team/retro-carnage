// Package buy_your_weapons just shows a basic text animation.
package buy_your_weapons

import (
	"fmt"
	"retro-carnage/input"
	"retro-carnage/ui/common"
	"retro-carnage/ui/common/fonts"

	pixel "github.com/Retro-Carnage-Team/pixel2"
	"github.com/Retro-Carnage-Team/pixel2/backends/opengl"
	"github.com/Retro-Carnage-Team/pixel2/ext/text"
)

const (
	timeAfterLastChar = 500
	timeBetweenChars  = 120
)

type Screen struct {
	millisecondsPassed   int64
	PlayerIdx            int
	screenChangeRequired common.ScreenChangeCallback
	text                 string
	textLength           int
	window               *opengl.Window
}

func (s *Screen) SetInputController(_ input.InputController) {
	// This screen doesn't process user input.
}

func (s *Screen) SetScreenChangeCallback(callback common.ScreenChangeCallback) {
	s.screenChangeRequired = callback
}

func (s *Screen) SetWindow(window *opengl.Window) {
	s.window = window
}

func (s *Screen) SetUp() {
	// no setup action required
}

func (s *Screen) Update(elapsedTimeInMs int64) {
	s.millisecondsPassed += elapsedTimeInMs
	if s.textLength < 25 {
		if s.millisecondsPassed >= timeBetweenChars {
			s.textLength++
			s.text = s.getFullText()[:s.textLength]
			s.millisecondsPassed = 0
		}
	} else if s.millisecondsPassed >= timeAfterLastChar {
		if s.PlayerIdx == 0 {
			s.screenChangeRequired(common.ShopP1)
		} else {
			s.screenChangeRequired(common.ShopP2)
		}
	}
	s.drawText()
}

func (s *Screen) TearDown() {
	// no tear down action required
}

func (s *Screen) String() string {
	if s.PlayerIdx == 0 {
		return string(common.BuyYourWeaponsP1)
	}
	return string(common.BuyYourWeaponsP2)
}

func (s *Screen) drawText() {
	var defaultFontSize = fonts.DefaultFontSize()
	var lineDimensions = fonts.GetTextDimension(defaultFontSize, s.getFullText())

	var vertCenter = s.window.Bounds().Max.Y / 2
	var lineX = (s.window.Bounds().Max.X - lineDimensions.X) / 2

	var txt = text.New(pixel.V(lineX, vertCenter), fonts.SizeToFontAtlas[defaultFontSize])
	txt.Color = common.White
	_, _ = fmt.Fprint(txt, s.text)
	txt.Draw(s.window, pixel.IM)
}

func (s *Screen) getFullText() string {
	return fmt.Sprintf("BUY YOUR WEAPONS PLAYER %d", s.PlayerIdx+1)
}
