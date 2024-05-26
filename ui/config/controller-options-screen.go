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
	optionControllerPreviousController = iota
	optionControllerNextController
	optionControllerDigitalAxis
	optionControllerAction
	optionControllerNextWeapon
	optionControllerPreviousWeapon
	optionControllerSave
	optionControllerBack
)

var (
	gamepadButtonNames = map[pixelgl.GamepadButton]string{
		pixelgl.ButtonA:           "A",
		pixelgl.ButtonB:           "B",
		pixelgl.ButtonX:           "X",
		pixelgl.ButtonY:           "Y",
		pixelgl.ButtonLeftBumper:  "LEFT BUMPER",
		pixelgl.ButtonRightBumper: "RIGHT BUMPER",
		pixelgl.ButtonBack:        "BACK",
		pixelgl.ButtonStart:       "START",
		pixelgl.ButtonGuide:       "GUIDE",
		pixelgl.ButtonLeftThumb:   "LEFT THUMB",
		pixelgl.ButtonRightThumb:  "RIGHT THUMB",
		pixelgl.ButtonDpadUp:      "DPAD UP",
		pixelgl.ButtonDpadRight:   "DPAD RIGHT",
		pixelgl.ButtonDpadDown:    "DPAD DOWN",
		pixelgl.ButtonDpadLeft:    "DPAD LEFT",
	}

	optionControllerFocusChanges = []focusChange{
		{movedRight: true, currentSelection: []int{optionControllerPreviousController}, nextSelection: optionControllerNextController},
		{movedLeft: true, currentSelection: []int{optionControllerNextController}, nextSelection: optionControllerPreviousController},
		{movedDown: true, currentSelection: []int{optionControllerPreviousController, optionControllerNextController}, nextSelection: optionControllerDigitalAxis},
		{movedUp: true, currentSelection: []int{optionControllerDigitalAxis}, nextSelection: optionControllerPreviousController},
		{movedDown: true, currentSelection: []int{optionControllerDigitalAxis}, nextSelection: optionControllerAction},
		{movedUp: true, currentSelection: []int{optionControllerAction}, nextSelection: optionControllerDigitalAxis},
		{movedDown: true, currentSelection: []int{optionControllerAction}, nextSelection: optionControllerNextWeapon},
		{movedUp: true, currentSelection: []int{optionControllerNextWeapon}, nextSelection: optionControllerAction},
		{movedDown: true, currentSelection: []int{optionControllerNextWeapon}, nextSelection: optionControllerPreviousWeapon},
		{movedUp: true, currentSelection: []int{optionControllerPreviousWeapon}, nextSelection: optionControllerNextWeapon},
		{movedDown: true, currentSelection: []int{optionControllerPreviousWeapon}, nextSelection: optionControllerSave},
		{movedUp: true, currentSelection: []int{optionControllerSave, optionControllerBack}, nextSelection: optionControllerPreviousWeapon},
		{movedRight: true, currentSelection: []int{optionControllerSave}, nextSelection: optionControllerBack},
		{movedLeft: true, currentSelection: []int{optionControllerBack}, nextSelection: optionControllerSave},
	}
)

type ControllerOptionsScreen struct {
	PlayerIdx               int
	availableDevices        []input.InputDeviceInfo
	defaultFontSize         int
	inputConfig             config.InputDeviceConfiguration
	inputController         input.InputController
	maxControllerNameLength float64
	screenChangeRequired    common.ScreenChangeCallback
	selectedDeviceIndex     int
	selectedOption          int
	textDimensions          map[string]*geometry.Point
	waitForInput            bool
	window                  *pixelgl.Window
}

func (s *ControllerOptionsScreen) SetUp() {
	s.defaultFontSize = fonts.DefaultFontSize()
	s.selectedOption = optionControllerPreviousController
	s.textDimensions = fonts.GetTextDimensions(
		s.defaultFontSize, txtInputSettingsP1, txtInputSettingsP1, txtSave, txtBack, txtNotConfigured, txtKeyboard,
		txtController, txtActionFire, txtNextWeapon, txtPrevWeapon, txtSelection, txtDecrease, txtIncrease, txtDigitalAxis,
	)
	s.availableDevices = s.inputController.GetInputDeviceInfos()
	s.maxControllerNameLength = s.getMaxWidthOfControllerNames()

	var inputConfigs = config.GetConfigService().LoadInputDeviceConfigurations()
	if len(inputConfigs) > s.PlayerIdx {
		s.inputConfig = inputConfigs[s.PlayerIdx]
		for i, d := range s.availableDevices {
			if d.DeviceName == s.inputConfig.DeviceName {
				s.selectedDeviceIndex = i
				break
			}
		}
	} else {
		s.inputConfig.DeviceName = config.DeviceNameKeyboard
	}
}

func (s *ControllerOptionsScreen) Update(_ int64) {
	if !s.waitForInput {
		s.processUserInput()
	} else {
		s.assignSelectedButton()
	}

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
	if s.selectedOption == optionControllerBack {
		drawTextSelectionRect(s.window, txt.Bounds())
	}

	txt = s.drawText(txtSave, lineLocationX-s.textDimensions[txtSave].X-3*buttonPadding, float64(lineLocationY), common.White)
	if s.selectedOption == optionControllerSave {
		drawTextSelectionRect(s.window, txt.Bounds())
	}
}

func (s *ControllerOptionsScreen) renderControllerSelectionValues(valueDistanceLeft float64, controllerLabelLocationY float64) {
	var txt = s.drawText(txtDecrease, valueDistanceLeft, controllerLabelLocationY, common.White)
	if s.selectedOption == optionControllerPreviousController {
		drawTextSelectionRect(s.window, txt.Bounds())
	} else {
		drawPossibleSelectionRect(s.window, txt.Bounds())
	}

	var distanceLeftControllerName = txt.Bounds().Max.X + buttonPadding*3
	s.drawText(s.inputConfig.DeviceName, distanceLeftControllerName, controllerLabelLocationY, common.White)

	txt = s.drawText(txtIncrease, distanceLeftControllerName+s.maxControllerNameLength+buttonPadding*3, controllerLabelLocationY, common.White)
	if s.selectedOption == optionControllerNextController {
		drawTextSelectionRect(s.window, txt.Bounds())
	} else {
		drawPossibleSelectionRect(s.window, txt.Bounds())
	}
}

func (s *ControllerOptionsScreen) renderDigitalAxisValue(valueDistanceLeft float64, axisLabelLocationY float64) {
	var txt *text.Text
	if s.inputConfig.GamepadConfiguration.HasDigitalAxis {
		txt = s.drawText(txtSelection, valueDistanceLeft, axisLabelLocationY, common.White)
	} else {
		txt = s.drawText(txtSelection, valueDistanceLeft, axisLabelLocationY, common.Black)
	}

	if s.selectedOption == optionControllerDigitalAxis {
		drawTextSelectionRect(s.window, txt.Bounds())
	} else {
		drawPossibleSelectionRect(s.window, txt.Bounds())
	}
}

func (s *ControllerOptionsScreen) renderActionValue(valueDistanceLeft float64, actionLabelLocationY float64) {
	var output string
	if s.selectedOption == optionControllerAction && s.waitForInput {
		output = txtPressButton
	} else {
		output = s.getDisplayTextForValue(s.inputConfig.InputFire)
	}

	var txt = s.drawText(output, valueDistanceLeft, actionLabelLocationY, common.White)
	if s.selectedOption == optionControllerAction {
		drawTextSelectionRect(s.window, txt.Bounds())
	} else {
		drawPossibleSelectionRect(s.window, txt.Bounds())
	}
}

func (s *ControllerOptionsScreen) renderNextWeaponValue(valueDistanceLeft float64, nextWeaponLabelLocationY float64) {
	var output string
	if s.selectedOption == optionControllerNextWeapon && s.waitForInput {
		output = txtPressButton
	} else {
		output = s.getDisplayTextForValue(s.inputConfig.InputNextWeapon)
	}

	var txt = s.drawText(output, valueDistanceLeft, nextWeaponLabelLocationY, common.White)
	if s.selectedOption == optionControllerNextWeapon {
		drawTextSelectionRect(s.window, txt.Bounds())
	} else {
		drawPossibleSelectionRect(s.window, txt.Bounds())
	}
}

func (s *ControllerOptionsScreen) renderPreviousWeaponValue(valueDistanceLeft float64, previousWeaponLabelLocationY float64) {
	var output string
	if s.selectedOption == optionControllerPreviousWeapon && s.waitForInput {
		output = txtPressButton
	} else {
		output = s.getDisplayTextForValue(s.inputConfig.InputPreviousWeapon)
	}

	var txt = s.drawText(output, valueDistanceLeft, previousWeaponLabelLocationY, common.White)
	if s.selectedOption == optionControllerPreviousWeapon {
		drawTextSelectionRect(s.window, txt.Bounds())
	} else {
		drawPossibleSelectionRect(s.window, txt.Bounds())
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
	if s.PlayerIdx == 0 {
		return string(common.ConfigurationControlsP1)
	}
	return string(common.ConfigurationControlsP2)
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
	case optionControllerPreviousController:
		s.selectPreviousController()
	case optionControllerNextController:
		s.selectNextController()
	case optionControllerDigitalAxis:
		s.inputConfig.GamepadConfiguration.HasDigitalAxis = !s.inputConfig.GamepadConfiguration.HasDigitalAxis
	case optionControllerAction, optionControllerNextWeapon, optionControllerPreviousWeapon:
		s.waitForInput = true
	case optionControllerSave:
		var err = config.GetConfigService().SaveInputDeviceConfiguration(s.inputConfig, s.PlayerIdx)
		if nil != err {
			logging.Warning.Printf("failed to save controller settings for player %d: %s", s.PlayerIdx, err)
		}
		s.screenChangeRequired(common.ConfigurationControls)
	case optionControllerBack:
		s.screenChangeRequired(common.ConfigurationControls)
	default:
		logging.Error.Fatal("Unexpected selection in ControllerOptionsScreen")
	}
}

func (s *ControllerOptionsScreen) assignSelectedButton() {
	var selectedValue = -1
	if s.inputConfig.DeviceName == config.DeviceNameKeyboard {
		for _, btn := range pixelgl.KeyboardButtons {
			if s.window.JustPressed(btn) {
				selectedValue = int(btn)
				break
			}
		}
	} else {
		for _, btn := range pixelgl.GamepadButtons {
			if s.window.JoystickJustPressed(pixelgl.Joystick(s.inputConfig.GamepadConfiguration.JoystickIndex), btn) {
				selectedValue = int(btn)
				break
			}
		}
	}

	if selectedValue == -1 {
		return
	}

	switch s.selectedOption {
	case optionControllerAction:
		s.inputConfig.InputFire = selectedValue
		s.waitForInput = false
	case optionControllerNextWeapon:
		s.inputConfig.InputNextWeapon = selectedValue
		s.waitForInput = false
	case optionControllerPreviousWeapon:
		s.inputConfig.InputPreviousWeapon = selectedValue
		s.waitForInput = false
	}
}

func (s *ControllerOptionsScreen) drawText(output string, x float64, y float64, col color.RGBA) *text.Text {
	var txt = text.New(pixel.V(x, y), fonts.SizeToFontAtlas[s.defaultFontSize])
	txt.Color = col
	_, _ = fmt.Fprint(txt, output)
	txt.Draw(s.window, pixel.IM)
	return txt
}

func (s *ControllerOptionsScreen) selectPreviousController() {
	s.selectedDeviceIndex = s.selectedDeviceIndex - 1
	if s.selectedDeviceIndex < 0 {
		s.selectedDeviceIndex = len(s.availableDevices) - 1
	}
	s.inputConfig.DeviceName = s.availableDevices[s.selectedDeviceIndex].DeviceName
	s.inputConfig.JoystickIndex = s.availableDevices[s.selectedDeviceIndex].JoystickIndex
}

func (s *ControllerOptionsScreen) selectNextController() {
	s.selectedDeviceIndex = (s.selectedDeviceIndex + 1) % len(s.availableDevices)
	s.inputConfig.DeviceName = s.availableDevices[s.selectedDeviceIndex].DeviceName
	s.inputConfig.JoystickIndex = s.availableDevices[s.selectedDeviceIndex].JoystickIndex
}

func (s *ControllerOptionsScreen) getMaxWidthOfControllerNames() float64 {
	var result = 0.0
	for _, m := range s.availableDevices {
		var width = fonts.GetTextDimension(s.defaultFontSize, m.DeviceName).X
		if width > result {
			result = width
		}
	}
	return result
}

func (s *ControllerOptionsScreen) getDisplayTextForValue(value int) string {
	if s.inputConfig.DeviceName == config.DeviceNameKeyboard {
		return pixelgl.Button(value).String()
	}

	var gpButton = pixelgl.GamepadButton(value)
	return gamepadButtonNames[gpButton]
}
