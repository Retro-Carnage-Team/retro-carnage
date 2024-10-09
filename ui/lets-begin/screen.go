package lets_begin

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
	displayText = "LET THE MISSION BEGIN..."
)

type Screen struct {
	controller *controller
	model      *model
	window     *opengl.Window
}

func NewScreen() *Screen {
	var model = model{}
	var result = Screen{
		controller: newController(&model),
		model:      &model,
	}
	return &result
}

func (s *Screen) SetInputController(_ input.InputController) {
	// Screen doesn't process user input. So no implementation required.
}

func (s *Screen) SetScreenChangeCallback(callback common.ScreenChangeCallback) {
	s.controller.setScreenChangeCallback(callback)
}

func (s *Screen) SetWindow(window *opengl.Window) {
	s.window = window
}

func (s *Screen) SetUp() {
	// not required
}

func (s *Screen) Update(elapsedTimeInMs int64) {
	s.controller.update(elapsedTimeInMs)
	s.drawScreen()
}

func (s *Screen) TearDown() {
	// No tear down action required
}

func (s *Screen) String() string {
	return string(common.LetTheMissionBegin)
}

func (s *Screen) drawScreen() {
	var defaultFontSize = fonts.DefaultFontSize()
	var lineDimensions = fonts.GetTextDimension(defaultFontSize, displayText)

	var vertCenter = s.window.Bounds().Max.Y / 2
	var lineX = (s.window.Bounds().Max.X - lineDimensions.X) / 2

	var txt = text.New(pixel.V(lineX, vertCenter), fonts.SizeToFontAtlas[defaultFontSize])
	txt.Color = common.White
	_, _ = fmt.Fprint(txt, s.model.text)
	txt.Draw(s.window, pixel.IM)
}
