package config

import (
	"retro-carnage/engine/characters"
	"retro-carnage/input"
	"retro-carnage/ui/common"
	"retro-carnage/ui/common/fonts"

	"github.com/Retro-Carnage-Team/pixel2/backends/opengl"
)

const (
	txtOnePlayerGame = "1 PLAYER GAME"
	txtTwoPlayerGame = "2 PLAYER GAME"
)

type ResultScreen struct {
	controller *resultController
	model      *resultModel
	window     *opengl.Window
}

func NewResultScreen() *ResultScreen {
	var model = resultModel{
		infoTextPlayerOne: "PLAYER 1: ",
		infoTextPlayerTwo: "PLAYER 2: ",
	}
	var controller = newResultController(&model)
	var result = ResultScreen{
		controller: controller,
		model:      &model,
	}
	return &result
}

func (s *ResultScreen) SetUp() {
	s.controller.setUp()
}

func (s *ResultScreen) Update(timeElapsedInMs int64) {
	s.controller.update(timeElapsedInMs)
	s.drawScreen()
}

func (s *ResultScreen) TearDown() {
	// no tear down action required
}

func (s *ResultScreen) SetInputController(controller input.InputController) {
	s.controller.setInputController(controller)
}

func (s *ResultScreen) SetScreenChangeCallback(callback common.ScreenChangeCallback) {
	s.controller.setScreenChangeCallback(callback)
}

func (s *ResultScreen) SetWindow(window *opengl.Window) {
	s.window = window
}

func (s *ResultScreen) String() string {
	return string(common.ConfigurationResult)
}

func (s *ResultScreen) drawScreen() {
	s.window.Clear(common.Black)
	renderer := fonts.TextRenderer{Window: s.window}
	if characters.PlayerController.NumberOfPlayers() == 1 {
		renderer.DrawLineToScreenCenter(txtOnePlayerGame, 2, common.Green)
		renderer.DrawLineToScreenCenter(s.model.infoTextPlayerOne, -1, common.Yellow)
	} else {
		renderer.DrawLineToScreenCenter(txtTwoPlayerGame, 2, common.Green)
		renderer.DrawLineToScreenCenter(s.model.infoTextPlayerOne, -1, common.Yellow)
		renderer.DrawLineToScreenCenter(s.model.infoTextPlayerTwo, -2.5, common.Yellow)
	}
}
