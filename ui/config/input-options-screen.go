package config

import (
	"fmt"
	"image/color"
	"retro-carnage/config"
	"retro-carnage/engine/geometry"
	"retro-carnage/input"
	"retro-carnage/logging"
	"retro-carnage/ui/common"
	"retro-carnage/ui/common/fonts"

	"github.com/Retro-Carnage-Team/pixel"
	"github.com/Retro-Carnage-Team/pixel/imdraw"
	"github.com/Retro-Carnage-Team/pixel/pixelgl"
	"github.com/Retro-Carnage-Team/pixel/text"
)

const (
	optionInputSelectPlayer1 int = iota
	optionInputSelectPlayer2
	optionInputBack
)

var (
	optionInputFocusChanges = []focusChange{
		{movedDown: true, currentSelection: []int{optionInputSelectPlayer1}, nextSelection: optionInputSelectPlayer2},
		{movedUp: true, currentSelection: []int{optionInputSelectPlayer2}, nextSelection: optionInputSelectPlayer1},
		{movedDown: true, currentSelection: []int{optionInputSelectPlayer2}, nextSelection: optionInputBack},
		{movedUp: true, currentSelection: []int{optionInputBack}, nextSelection: optionInputSelectPlayer2},
	}
)

type InputOptionsScreen struct {
	defaultFontSize      int
	inputConfig          []config.InputDeviceConfiguration
	inputController      input.InputController
	screenChangeRequired common.ScreenChangeCallback
	selectedOption       int
	textDimensions       map[string]*geometry.Point
	window               *pixelgl.Window
}

func (s *InputOptionsScreen) SetUp() {
	logging.Info.Println("InputOptionsScreen.Setup")
	s.defaultFontSize = fonts.DefaultFontSize()
	s.selectedOption = optionAudioPlayEffects
	s.textDimensions = fonts.GetTextDimensions(s.defaultFontSize, txtInputSettings, txtBack, txtPlayer1, txtPlayer2, txtNotConfigured)
	s.inputConfig = config.GetConfigService().LoadInputDeviceConfigurations()
}

func (s *InputOptionsScreen) Update(_ int64) {
	s.processUserInput()

	// draw headline
	var headlineLocationY = s.window.Bounds().H() - headlineDistanceTop
	var txt = s.drawText(txtInputSettings, headlineDistanceLeft, headlineLocationY, common.White)
	imd := imdraw.New(nil)
	imd.Color = common.White
	imd.EndShape = imdraw.RoundEndShape
	imd.Push(pixel.V(txt.Bounds().Min.X, txt.Bounds().Min.Y), pixel.V(txt.Bounds().Max.X, txt.Bounds().Min.Y))
	imd.Line(4)
	imd.Draw(s.window)

	// Player 1
	var p1LabelLocationY = headlineLocationY - 5*s.textDimensions[txtPlayer1].Y
	txt = s.drawText(txtPlayer1, headlineDistanceLeft, p1LabelLocationY, common.White)
	var valueDistanceLeft = headlineDistanceLeft + txt.Bounds().W()/2*3

	var p1DeviceName = txtNotConfigured
	if len(s.inputConfig) > 0 {
		p1DeviceName = s.inputConfig[0].DeviceName
	}
	txt = s.drawText(p1DeviceName, valueDistanceLeft, p1LabelLocationY, common.White)
	if s.selectedOption == optionInputSelectPlayer1 {
		drawTextSelectionRect(s.window, txt.Bounds())
	} else {
		drawPossibleSelectionRect(s.window, txt.Bounds())
	}

	// Player 2
	var p2LabelLocationY = p1LabelLocationY - 3.5*s.textDimensions[txtPlayer2].Y
	s.drawText(txtPlayer2, headlineDistanceLeft, p2LabelLocationY, common.White)

	var p2DeviceName = txtNotConfigured
	if len(s.inputConfig) > 1 {
		p2DeviceName = s.inputConfig[1].DeviceName
	}
	txt = s.drawText(p2DeviceName, valueDistanceLeft, p2LabelLocationY, common.White)
	if s.selectedOption == optionInputSelectPlayer2 {
		drawTextSelectionRect(s.window, txt.Bounds())
	} else {
		drawPossibleSelectionRect(s.window, txt.Bounds())
	}

	// Action buttons
	var lineLocationX = s.window.Bounds().W() - headlineDistanceTop - s.textDimensions[txtBack].X
	var lineLocationY = headlineDistanceTop
	txt = s.drawText(txtBack, lineLocationX, float64(lineLocationY), common.White)
	if s.selectedOption == optionInputBack {
		drawTextSelectionRect(s.window, txt.Bounds())
	}
}

func (s *InputOptionsScreen) TearDown() {
	// no tear down action required
}

func (s *InputOptionsScreen) SetInputController(controller input.InputController) {
	s.inputController = controller
}

func (s *InputOptionsScreen) SetScreenChangeCallback(callback common.ScreenChangeCallback) {
	s.screenChangeRequired = callback
}

func (s *InputOptionsScreen) SetWindow(window *pixelgl.Window) {
	s.window = window
}

func (s *InputOptionsScreen) String() string {
	return string(common.ConfigurationControls)
}

func (s *InputOptionsScreen) processUserInput() {
	var uiEventState = s.inputController.GetUiEventStateCombined()
	if nil == uiEventState {
		return
	}

focusHandling:
	for _, fc := range optionInputFocusChanges {
		if fc.movedLeft == uiEventState.MovedLeft &&
			fc.movedRight == uiEventState.MovedRight &&
			fc.movedDown == uiEventState.MovedDown &&
			fc.movedUp == uiEventState.MovedUp {
			for _, i := range fc.currentSelection {
				if i == s.selectedOption {
					s.selectedOption = fc.nextSelection
					break focusHandling
				}
			}
		}
	}

	if uiEventState.PressedButton {
		s.processOptionSelected()
	}
}

func (s *InputOptionsScreen) processOptionSelected() {
	switch s.selectedOption {
	case optionInputSelectPlayer1:
		s.screenChangeRequired(common.ConfigurationControlsP1)
	case optionInputSelectPlayer2:
		s.screenChangeRequired(common.ConfigurationControlsP2)
	case optionInputBack:
		s.screenChangeRequired(common.ConfigurationOptions)
	default:
		logging.Error.Fatal("Unexpected selection in InputOptionsScreen")
	}
}

func (s *InputOptionsScreen) drawText(output string, x float64, y float64, col color.RGBA) *text.Text {
	var txt = text.New(pixel.V(x, y), fonts.SizeToFontAtlas[s.defaultFontSize])
	txt.Color = col
	_, _ = fmt.Fprint(txt, output)
	txt.Draw(s.window, pixel.IM)
	return txt
}
