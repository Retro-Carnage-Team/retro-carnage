package config

import (
	"fmt"
	"image/color"
	"retro-carnage/engine/geometry"
	"retro-carnage/engine/input"
	"retro-carnage/logging"
	"retro-carnage/ui/common"
	"retro-carnage/ui/common/fonts"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
)

const (
	// optionInputSelectPlayer1 = 1
	// optionInputSelectPlayer2 = 2
	optionControllerSave = 3
	optionControllerBack = 4
)

var (
	optionControllerFocusChanges = []focusChange{
		// {movedDown: true, currentSelection: []int{optionInputSelectPlayer1}, nextSelection: optionInputSelectPlayer2},
		// {movedUp: true, currentSelection: []int{optionInputSelectPlayer2}, nextSelection: optionInputSelectPlayer1},
		// {movedDown: true, currentSelection: []int{optionInputSelectPlayer2}, nextSelection: optionInputBack},
		// {movedUp: true, currentSelection: []int{optionInputBack}, nextSelection: optionInputSelectPlayer2},
	}
)

type ControllerOptionsScreen struct {
	PlayerIdx            int
	defaultFontSize      int
	inputController      input.InputController
	screenChangeRequired common.ScreenChangeCallback
	selectedOption       int
	textDimensions       map[string]*geometry.Point
	window               *pixelgl.Window
}

func (s *ControllerOptionsScreen) SetUp() {
	s.defaultFontSize = fonts.DefaultFontSize()
	s.selectedOption = optionAudioPlayEffects
	s.textDimensions = fonts.GetTextDimensions(s.defaultFontSize, txtInputSettingsP1, txtInputSettingsP1, txtSave, txtBack, txtNotConfigured)
}

func (s *ControllerOptionsScreen) Update(_ int64) {
	s.processUserInput()

	// draw headline
	var headline = txtInputSettingsP1
	if s.PlayerIdx == 1 {
		headline = txtInputSettingsP2
	}
	var headlineLocationY = s.window.Bounds().H() - headlineDistanceTop
	var txt = s.drawText(headline, headlineDistanceLeft, headlineLocationY, common.White)
	imd := imdraw.New(nil)
	imd.Color = common.White
	imd.EndShape = imdraw.RoundEndShape
	imd.Push(pixel.V(txt.Bounds().Min.X, txt.Bounds().Min.Y), pixel.V(txt.Bounds().Max.X, txt.Bounds().Min.Y))
	imd.Line(4)
	imd.Draw(s.window)

	// Action buttons
	var lineLocationX = s.window.Bounds().W() - headlineDistanceTop - s.textDimensions[txtBack].X
	var lineLocationY = headlineDistanceTop
	txt = s.drawText(txtBack, lineLocationX, float64(lineLocationY), common.White)
	if s.selectedOption == optionAudioBack {
		drawTextSelectionRect(s.window, txt.Bounds())
	}

	txt = s.drawText(txtSave, lineLocationX-s.textDimensions[txtSave].X-3*buttonPadding, float64(lineLocationY), common.White)
	if s.selectedOption == optionAudioSave {
		drawTextSelectionRect(s.window, txt.Bounds())
	}
}

func (s *ControllerOptionsScreen) TearDown() {
	// no tear down action required
}

func (s *ControllerOptionsScreen) SetInputController(controller input.InputController) {
	s.inputController = controller
}

func (s *ControllerOptionsScreen) SetScreenChangeCallback(callback common.ScreenChangeCallback) {
	s.screenChangeRequired = callback
}

func (s *ControllerOptionsScreen) SetWindow(window *pixelgl.Window) {
	s.window = window
}

func (s *ControllerOptionsScreen) String() string {
	return string(common.ConfigurationVideo)
}

func (s *ControllerOptionsScreen) processUserInput() {
	var uiEventState = s.inputController.GetUiEventStateCombined()
	if nil == uiEventState {
		return
	}

focusHandling:
	for _, fc := range optionControllerFocusChanges {
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

func (s *ControllerOptionsScreen) processOptionSelected() {
	switch s.selectedOption {
	case optionControllerSave:
		// TODO: Save controller settings
		s.screenChangeRequired(common.ConfigurationControls)
	case optionControllerBack:
		s.screenChangeRequired(common.ConfigurationControls)
	default:
		logging.Error.Fatal("Unexpected selection in ControllerOptionsScreen")
	}
}

func (s *ControllerOptionsScreen) drawText(output string, x float64, y float64, col color.RGBA) *text.Text {
	var txt = text.New(pixel.V(x, y), fonts.SizeToFontAtlas[s.defaultFontSize])
	txt.Color = col
	_, _ = fmt.Fprint(txt, output)
	txt.Draw(s.window, pixel.IM)
	return txt
}
