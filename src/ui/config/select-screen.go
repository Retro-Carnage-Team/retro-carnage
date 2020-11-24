package config

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"retro-carnage/engine/characters"
	"retro-carnage/engine/geometry"
	"retro-carnage/engine/input"
	"retro-carnage/ui/common"
	"retro-carnage/util"
)

const buttonPadding = 15
const txtSelectOnePlayerGame = "START 1 PLAYER GAME"
const txtSelectTwoPlayerGame = "START 2 PLAYER GAME"

type SelectScreen struct {
	inputController      input.Controller
	screenChangeRequired common.ScreenChangeCallback
	selectedOption       int
	textDimensions       map[string]*geometry.Point
	window               *pixelgl.Window
}

func (s *SelectScreen) SetUp() {
	s.selectedOption = 1
	s.textDimensions = common.GetTextDimensions(text.New(pixel.V(0, 0), common.DefaultAtlas),
		txtSelectOnePlayerGame, txtSelectTwoPlayerGame)
}

func (s *SelectScreen) Update(_ int64) {
	s.window.Clear(common.Black)
	s.processUserInput()

	var vertCenter = s.window.Bounds().Max.Y / 2
	var firstLineX = (s.window.Bounds().Max.X - s.textDimensions[txtSelectOnePlayerGame].X) / 2
	var firstLineY = vertCenter + 1.5*s.textDimensions[txtSelectOnePlayerGame].Y

	var txt = text.New(pixel.V(firstLineX, firstLineY), common.DefaultAtlas)
	txt.Color = common.White
	_, _ = fmt.Fprint(txt, txtSelectOnePlayerGame)
	txt.Draw(s.window, pixel.IM)

	var secondLineX = (s.window.Bounds().Max.X - s.textDimensions[txtSelectTwoPlayerGame].X) / 2
	var secondLineY = vertCenter + -1.5*s.textDimensions[txtSelectTwoPlayerGame].Y

	txt = text.New(pixel.V(secondLineX, secondLineY), common.DefaultAtlas)
	txt.Color = common.White
	_, _ = fmt.Fprint(txt, txtSelectTwoPlayerGame)
	txt.Draw(s.window, pixel.IM)

	var bottomFirst = firstLineY - buttonPadding
	var bottomSecond = secondLineY - buttonPadding
	var topFirst = firstLineY + s.textDimensions[txtSelectOnePlayerGame].Y
	var topSecond = secondLineY + s.textDimensions[txtSelectTwoPlayerGame].Y
	var left = util.Min(firstLineX, secondLineX) - buttonPadding
	var right = util.Min(firstLineX, secondLineX) + util.Min(s.textDimensions[txtSelectOnePlayerGame].X, s.textDimensions[txtSelectTwoPlayerGame].X) + buttonPadding

	if 1 == s.selectedOption {
		imd := imdraw.New(nil)
		imd.Color = common.Yellow
		imd.EndShape = imdraw.RoundEndShape
		imd.Push(pixel.V(left, bottomFirst), pixel.V(right, bottomFirst))
		imd.Push(pixel.V(left, bottomFirst), pixel.V(left, topFirst))
		imd.Push(pixel.V(left, topFirst), pixel.V(right, topFirst))
		imd.Push(pixel.V(right, bottomFirst), pixel.V(right, topFirst))
		imd.Line(4)
		imd.Draw(s.window)
	} else if 2 == s.selectedOption {
		imd := imdraw.New(nil)
		imd.Color = common.Yellow
		imd.EndShape = imdraw.RoundEndShape
		imd.Push(pixel.V(left, bottomSecond), pixel.V(right, bottomSecond))
		imd.Push(pixel.V(left, bottomSecond), pixel.V(left, topSecond))
		imd.Push(pixel.V(left, topSecond), pixel.V(right, topSecond))
		imd.Push(pixel.V(right, bottomSecond), pixel.V(right, topSecond))
		imd.Line(4)
		imd.Draw(s.window)
	}
}

func (s *SelectScreen) TearDown() {}

func (s *SelectScreen) SetInputController(controller input.Controller) {
	s.inputController = controller
}

func (s *SelectScreen) SetScreenChangeCallback(callback common.ScreenChangeCallback) {
	s.screenChangeRequired = callback
}

func (s *SelectScreen) SetWindow(window *pixelgl.Window) {
	s.window = window
}

func (s *SelectScreen) String() string {
	return string(common.ConfigurationSelect)
}

func (s *SelectScreen) processUserInput() {
	var uiEventState = s.inputController.GetControllerUiEventStateCombined()
	if nil != uiEventState {
		if uiEventState.PressedButton {
			s.inputController.AssignControllersToPlayers()
			characters.PlayerController.StartNewGame(s.selectedOption)
			s.screenChangeRequired(common.ConfigurationResult)
		} else if uiEventState.MovedUp && 2 == s.selectedOption {
			s.selectedOption = 1
		} else if uiEventState.MovedDown && 1 == s.selectedOption && s.inputController.HasTwoOrMoreDevices() {
			s.selectedOption = 2
		}
	}
}
