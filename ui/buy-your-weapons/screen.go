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

type Screen struct {
	controller *controller
	window     *opengl.Window
}

func NewScreen(playerIdx int) *Screen {
	var result = Screen{
		controller: newController(playerIdx),
	}
	return &result
}

func (s *Screen) SetInputController(_ input.InputController) {
	// This screen doesn't process user input.
}

func (s *Screen) SetScreenChangeCallback(callback common.ScreenChangeCallback) {
	s.controller.setScreenChangeCallback(callback)
}

func (s *Screen) SetWindow(window *opengl.Window) {
	s.window = window
}

func (s *Screen) SetUp() {
	// no setup action required
}

func (s *Screen) Update(elapsedTimeInMs int64) {
	s.controller.update(elapsedTimeInMs)
	s.drawScreen()
}

func (s *Screen) TearDown() {
	// no tear down action required
}

func (s *Screen) String() string {
	if s.controller.playerIdx == 0 {
		return string(common.BuyYourWeaponsP1)
	}
	return string(common.BuyYourWeaponsP2)
}

func (s *Screen) drawScreen() {
	var defaultFontSize = fonts.DefaultFontSize()
	var lineDimensions = fonts.GetTextDimension(defaultFontSize, s.controller.getFullText())

	var vertCenter = s.window.Bounds().Max.Y / 2
	var lineX = (s.window.Bounds().Max.X - lineDimensions.X) / 2

	var txt = text.New(pixel.V(lineX, vertCenter), fonts.SizeToFontAtlas[defaultFontSize])
	txt.Color = common.White
	_, _ = fmt.Fprint(txt, s.controller.text)
	txt.Draw(s.window, pixel.IM)
}
