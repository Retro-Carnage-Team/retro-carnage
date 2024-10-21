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

	pixel "github.com/Retro-Carnage-Team/pixel2"
	"github.com/Retro-Carnage-Team/pixel2/backends/opengl"
	"github.com/Retro-Carnage-Team/pixel2/ext/imdraw"
	"github.com/Retro-Carnage-Team/pixel2/ext/text"
)

type InputOptionsScreen struct {
	controller      *inputOptionsController
	defaultFontSize int
	model           *inputOptionsModel
	textDimensions  map[string]*geometry.Point
	window          *opengl.Window
}

func NewInputOptionsScreen() *InputOptionsScreen {
	var model = inputOptionsModel{
		inputConfig:    config.GetConfigService().LoadInputDeviceConfigurations(),
		selectedOption: optionAudioPlayEffects,
	}
	var controller = newInputOptionsController(&model)
	var result = InputOptionsScreen{
		controller: controller,
		model:      &model,
	}
	return &result
}

func (s *InputOptionsScreen) SetUp() {
	logging.Info.Println("InputOptionsScreen.Setup")
	s.defaultFontSize = fonts.DefaultFontSize()
	s.textDimensions = fonts.GetTextDimensions(s.defaultFontSize, txtInputSettings, txtBack, txtPlayer1, txtPlayer2, txtNotConfigured)
}

func (s *InputOptionsScreen) Update(_ int64) {
	s.controller.update()

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
	if len(s.model.inputConfig) > 0 {
		p1DeviceName = s.model.inputConfig[0].DeviceName
	}
	txt = s.drawText(p1DeviceName, valueDistanceLeft, p1LabelLocationY, common.White)
	if s.model.selectedOption == optionInputSelectPlayer1 {
		drawTextSelectionRect(s.window, txt.Bounds())
	} else {
		drawPossibleSelectionRect(s.window, txt.Bounds())
	}

	// Player 2
	var p2LabelLocationY = p1LabelLocationY - 3.5*s.textDimensions[txtPlayer2].Y
	s.drawText(txtPlayer2, headlineDistanceLeft, p2LabelLocationY, common.White)

	var p2DeviceName = txtNotConfigured
	if len(s.model.inputConfig) > 1 {
		p2DeviceName = s.model.inputConfig[1].DeviceName
	}
	txt = s.drawText(p2DeviceName, valueDistanceLeft, p2LabelLocationY, common.White)
	if s.model.selectedOption == optionInputSelectPlayer2 {
		drawTextSelectionRect(s.window, txt.Bounds())
	} else {
		drawPossibleSelectionRect(s.window, txt.Bounds())
	}

	// Action buttons
	var lineLocationX = s.window.Bounds().W() - headlineDistanceTop - s.textDimensions[txtBack].X
	var lineLocationY = headlineDistanceTop
	txt = s.drawText(txtBack, lineLocationX, float64(lineLocationY), common.White)
	if s.model.selectedOption == optionInputBack {
		drawTextSelectionRect(s.window, txt.Bounds())
	}
}

func (s *InputOptionsScreen) TearDown() {
	// no tear down action required
}

func (s *InputOptionsScreen) SetInputController(controller input.InputController) {
	s.controller.setInputController(controller)
}

func (s *InputOptionsScreen) SetScreenChangeCallback(callback common.ScreenChangeCallback) {
	s.controller.setScreenChangeCallback(callback)
}

func (s *InputOptionsScreen) SetWindow(window *opengl.Window) {
	s.window = window
}

func (s *InputOptionsScreen) String() string {
	return string(common.ConfigurationControls)
}

func (s *InputOptionsScreen) drawText(output string, x float64, y float64, col color.RGBA) *text.Text {
	var txt = text.New(pixel.V(x, y), fonts.SizeToFontAtlas[s.defaultFontSize])
	txt.Color = col
	_, _ = fmt.Fprint(txt, output)
	txt.Draw(s.window, pixel.IM)
	return txt
}
