package config

import (
	"fmt"
	"image/color"
	"retro-carnage/engine/geometry"
	"retro-carnage/input"
	"retro-carnage/ui/common"
	"retro-carnage/ui/common/fonts"

	pixel "github.com/Retro-Carnage-Team/pixel2"
	"github.com/Retro-Carnage-Team/pixel2/backends/opengl"
	"github.com/Retro-Carnage-Team/pixel2/ext/imdraw"
	"github.com/Retro-Carnage-Team/pixel2/ext/text"
)

type ControllerOptionsScreen struct {
	controller              *controllerOptionsController
	defaultFontSize         int
	maxControllerNameLength float64
	model                   *controllerOptionsModel
	textDimensions          map[string]*geometry.Point
	window                  *opengl.Window
}

func NewControllerOptionsScreen(playerIdx int) *ControllerOptionsScreen {
	var model = controllerOptionsModel{
		playerIdx:      playerIdx,
		selectedOption: optionControllerPreviousController,
	}
	var controller = newControllerOptionsController(&model)
	var result = ControllerOptionsScreen{
		controller: controller,
		model:      &model,
	}
	return &result
}

func (s *ControllerOptionsScreen) SetUp() {
	s.controller.setUp()

	s.defaultFontSize = fonts.DefaultFontSize()
	s.textDimensions = fonts.GetTextDimensions(
		s.defaultFontSize, txtInputSettingsP1, txtInputSettingsP1, txtSave, txtBack, txtNotConfigured, txtKeyboard,
		txtController, txtActionFire, txtNextWeapon, txtPrevWeapon, txtSelection, txtDecrease, txtIncrease, txtDigitalAxis,
	)
	s.maxControllerNameLength = s.getMaxWidthOfControllerNames()
}

func (s *ControllerOptionsScreen) Update(_ int64) {
	s.controller.update(s.window)

	// draw headline
	var headline = txtInputSettingsP1
	if s.model.playerIdx == 1 {
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

	// Controller label
	var controllerLabelLocationY = headlineLocationY - 5*s.textDimensions[txtController].Y
	s.drawText(txtController, headlineDistanceLeft, controllerLabelLocationY, common.White)

	// Digital axis label
	var axisLabelLocationY = controllerLabelLocationY - 3*s.textDimensions[txtDigitalAxis].Y
	s.drawText(txtDigitalAxis, headlineDistanceLeft, axisLabelLocationY, common.White)

	// Action label
	var actionLabelLocationY = axisLabelLocationY - 3*s.textDimensions[txtActionFire].Y
	txt = s.drawText(txtActionFire, headlineDistanceLeft, actionLabelLocationY, common.White)
	var valueDistanceLeft = headlineDistanceLeft + txt.Bounds().W()/2*3

	// Next weapon label
	var nextWeaponLabelLocationY = actionLabelLocationY - 3*s.textDimensions[txtNextWeapon].Y
	s.drawText(txtNextWeapon, headlineDistanceLeft, nextWeaponLabelLocationY, common.White)

	// Previous weapon label
	var previousWeaponLabelLocationY = nextWeaponLabelLocationY - 3*s.textDimensions[txtNextWeapon].Y
	s.drawText(txtPrevWeapon, headlineDistanceLeft, previousWeaponLabelLocationY, common.White)

	// Values
	s.renderControllerSelectionValues(valueDistanceLeft, controllerLabelLocationY)
	s.renderDigitalAxisValue(valueDistanceLeft, axisLabelLocationY)
	s.renderActionValue(valueDistanceLeft, actionLabelLocationY)
	s.renderNextWeaponValue(valueDistanceLeft, nextWeaponLabelLocationY)
	s.renderPreviousWeaponValue(valueDistanceLeft, previousWeaponLabelLocationY)

	// Action buttons
	var lineLocationX = s.window.Bounds().W() - headlineDistanceTop - s.textDimensions[txtBack].X
	var lineLocationY = headlineDistanceTop
	txt = s.drawText(txtBack, lineLocationX, float64(lineLocationY), common.White)
	if s.model.selectedOption == optionControllerBack {
		drawTextSelectionRect(s.window, txt.Bounds())
	}

	txt = s.drawText(txtSave, lineLocationX-s.textDimensions[txtSave].X-3*buttonPadding, float64(lineLocationY), common.White)
	if s.model.selectedOption == optionControllerSave {
		drawTextSelectionRect(s.window, txt.Bounds())
	}
}

func (s *ControllerOptionsScreen) renderControllerSelectionValues(valueDistanceLeft float64, controllerLabelLocationY float64) {
	var txt = s.drawText(txtDecrease, valueDistanceLeft, controllerLabelLocationY, common.White)
	if s.model.selectedOption == optionControllerPreviousController {
		drawTextSelectionRect(s.window, txt.Bounds())
	} else {
		drawPossibleSelectionRect(s.window, txt.Bounds())
	}

	var distanceLeftControllerName = txt.Bounds().Max.X + buttonPadding*3
	s.drawText(s.model.inputConfig.DeviceName, distanceLeftControllerName, controllerLabelLocationY, common.White)

	txt = s.drawText(txtIncrease, distanceLeftControllerName+s.maxControllerNameLength+buttonPadding*3, controllerLabelLocationY, common.White)
	if s.model.selectedOption == optionControllerNextController {
		drawTextSelectionRect(s.window, txt.Bounds())
	} else {
		drawPossibleSelectionRect(s.window, txt.Bounds())
	}
}

func (s *ControllerOptionsScreen) renderDigitalAxisValue(valueDistanceLeft float64, axisLabelLocationY float64) {
	var txt *text.Text
	if s.model.inputConfig.GamepadConfiguration.HasDigitalAxis {
		txt = s.drawText(txtSelection, valueDistanceLeft, axisLabelLocationY, common.White)
	} else {
		txt = s.drawText(txtSelection, valueDistanceLeft, axisLabelLocationY, common.Black)
	}

	if s.model.selectedOption == optionControllerDigitalAxis {
		drawTextSelectionRect(s.window, txt.Bounds())
	} else {
		drawPossibleSelectionRect(s.window, txt.Bounds())
	}
}

func (s *ControllerOptionsScreen) renderActionValue(valueDistanceLeft float64, actionLabelLocationY float64) {
	var output string
	if s.model.selectedOption == optionControllerAction && s.model.waitForInput {
		output = txtPressButton
	} else {
		output = s.controller.getDisplayTextForValue(s.model.inputConfig.InputFire)
	}

	var txt = s.drawText(output, valueDistanceLeft, actionLabelLocationY, common.White)
	if s.model.selectedOption == optionControllerAction {
		drawTextSelectionRect(s.window, txt.Bounds())
	} else {
		drawPossibleSelectionRect(s.window, txt.Bounds())
	}
}

func (s *ControllerOptionsScreen) renderNextWeaponValue(valueDistanceLeft float64, nextWeaponLabelLocationY float64) {
	var output string
	if s.model.selectedOption == optionControllerNextWeapon && s.model.waitForInput {
		output = txtPressButton
	} else {
		output = s.controller.getDisplayTextForValue(s.model.inputConfig.InputNextWeapon)
	}

	var txt = s.drawText(output, valueDistanceLeft, nextWeaponLabelLocationY, common.White)
	if s.model.selectedOption == optionControllerNextWeapon {
		drawTextSelectionRect(s.window, txt.Bounds())
	} else {
		drawPossibleSelectionRect(s.window, txt.Bounds())
	}
}

func (s *ControllerOptionsScreen) renderPreviousWeaponValue(valueDistanceLeft float64, previousWeaponLabelLocationY float64) {
	var output string
	if s.model.selectedOption == optionControllerPreviousWeapon && s.model.waitForInput {
		output = txtPressButton
	} else {
		output = s.controller.getDisplayTextForValue(s.model.inputConfig.InputPreviousWeapon)
	}

	var txt = s.drawText(output, valueDistanceLeft, previousWeaponLabelLocationY, common.White)
	if s.model.selectedOption == optionControllerPreviousWeapon {
		drawTextSelectionRect(s.window, txt.Bounds())
	} else {
		drawPossibleSelectionRect(s.window, txt.Bounds())
	}
}

func (s *ControllerOptionsScreen) TearDown() {
	// no tear down action required
}

func (s *ControllerOptionsScreen) SetInputController(controller input.InputController) {
	s.controller.setInputController(controller)
}

func (s *ControllerOptionsScreen) SetScreenChangeCallback(callback common.ScreenChangeCallback) {
	s.controller.setScreenChangeCallback(callback)
}

func (s *ControllerOptionsScreen) SetWindow(window *opengl.Window) {
	s.window = window
}

func (s *ControllerOptionsScreen) String() string {
	if s.model.playerIdx == 0 {
		return string(common.ConfigurationControlsP1)
	}
	return string(common.ConfigurationControlsP2)
}

func (s *ControllerOptionsScreen) drawText(output string, x float64, y float64, col color.RGBA) *text.Text {
	var txt = text.New(pixel.V(x, y), fonts.SizeToFontAtlas[s.defaultFontSize])
	txt.Color = col
	_, _ = fmt.Fprint(txt, output)
	txt.Draw(s.window, pixel.IM)
	return txt
}

func (s *ControllerOptionsScreen) getMaxWidthOfControllerNames() float64 {
	var result = 0.0
	for _, m := range s.model.availableDevices {
		var width = fonts.GetTextDimension(s.defaultFontSize, m.DeviceName).X
		if width > result {
			result = width
		}
	}
	return result
}
