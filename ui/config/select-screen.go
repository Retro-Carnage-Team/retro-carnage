package config

import (
	"fmt"
	"retro-carnage/engine/geometry"
	"retro-carnage/input"
	"retro-carnage/ui/common"
	"retro-carnage/ui/common/fonts"

	pixel "github.com/Retro-Carnage-Team/pixel2"
	"github.com/Retro-Carnage-Team/pixel2/backends/opengl"
	"github.com/Retro-Carnage-Team/pixel2/ext/text"
)

const (
	txtSelectOnePlayerGame = "START 1 PLAYER GAME"
	txtSelectTwoPlayerGame = "START 2 PLAYER GAME"
	txtSelectOptions       = "OPTIONS"
)

type SelectScreen struct {
	controller      *selectController
	defaultFontSize int
	model           *selectModel
	textDimensions  map[string]*geometry.Point
	window          *opengl.Window
}

func NewSelectScreen() *SelectScreen {
	var model = selectModel{
		// Multiplayer has to be tested thoroughly first
		// multiplayerPossible: len(s.inputController.GetInputDeviceInfos()) > 1
		multiplayerPossible: false,
		selectedOption:      optionOnePlayer,
	}
	var controller = newSelectController(&model)
	var result = SelectScreen{
		controller: controller,
		model:      &model,
	}
	return &result
}

func (s *SelectScreen) SetUp() {
	s.defaultFontSize = fonts.DefaultFontSize()
	s.textDimensions = fonts.GetTextDimensions(s.defaultFontSize, txtSelectOnePlayerGame, txtSelectTwoPlayerGame, txtSelectOptions)
}

func (s *SelectScreen) Update(_ int64) {
	s.controller.update()
	s.drawScreen()
}

func (s *SelectScreen) TearDown() {
	// no tear down action required
}

func (s *SelectScreen) SetInputController(controller input.InputController) {
	s.controller.setInputController(controller)
}

func (s *SelectScreen) SetScreenChangeCallback(callback common.ScreenChangeCallback) {
	s.controller.setScreenChangeCallback(callback)
}

func (s *SelectScreen) SetWindow(window *opengl.Window) {
	s.window = window
}

func (s *SelectScreen) String() string {
	return string(common.ConfigurationSelect)
}

func (s *SelectScreen) drawScreen() {
	var vertCenter = s.window.Bounds().Max.Y / 2

	var firstLineX = (s.window.Bounds().Max.X - s.textDimensions[txtSelectOnePlayerGame].X) / 2
	var firstLineY = vertCenter + 1.5*s.textDimensions[txtSelectOnePlayerGame].Y

	var txt = text.New(pixel.V(firstLineX, firstLineY), fonts.SizeToFontAtlas[s.defaultFontSize])
	txt.Color = common.White
	_, _ = fmt.Fprint(txt, txtSelectOnePlayerGame)
	txt.Draw(s.window, pixel.IM)

	if s.model.selectedOption == optionOnePlayer {
		drawTextSelectionRect(s.window, txt.Bounds())
	}

	var secondLineX = (s.window.Bounds().Max.X - s.textDimensions[txtSelectTwoPlayerGame].X) / 2
	var secondLineY = vertCenter + -1.5*s.textDimensions[txtSelectTwoPlayerGame].Y

	var startLine3 = -1.5
	if s.model.multiplayerPossible {
		txt = text.New(pixel.V(secondLineX, secondLineY), fonts.SizeToFontAtlas[s.defaultFontSize])
		txt.Color = common.White
		_, _ = fmt.Fprint(txt, txtSelectTwoPlayerGame)
		txt.Draw(s.window, pixel.IM)
		startLine3 = -4.5

		if s.model.selectedOption == optionTwoPlayers {
			drawTextSelectionRect(s.window, txt.Bounds())
		}
	}

	var thirdLineX = (s.window.Bounds().Max.X - s.textDimensions[txtSelectOptions].X) / 2
	var thirdLineY = vertCenter + startLine3*s.textDimensions[txtSelectOptions].Y

	txt = text.New(pixel.V(thirdLineX, thirdLineY), fonts.SizeToFontAtlas[s.defaultFontSize])
	txt.Color = common.White
	_, _ = fmt.Fprint(txt, txtSelectOptions)
	txt.Draw(s.window, pixel.IM)

	if s.model.selectedOption == optionOptions {
		drawTextSelectionRect(s.window, txt.Bounds())
	}
}
