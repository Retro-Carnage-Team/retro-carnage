package config

import (
	"retro-carnage/config"
	"retro-carnage/input"
	"retro-carnage/logging"
	"retro-carnage/ui/common"

	pixel "github.com/Retro-Carnage-Team/pixel2"
	"github.com/Retro-Carnage-Team/pixel2/backends/opengl"
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
	gamepadButtonNames = map[pixel.GamepadButton]string{
		pixel.GamepadA:           "A",
		pixel.GamepadB:           "B",
		pixel.GamepadX:           "X",
		pixel.GamepadY:           "Y",
		pixel.GamepadLeftBumper:  "LEFT BUMPER",
		pixel.GamepadRightBumper: "RIGHT BUMPER",
		pixel.GamepadBack:        "BACK",
		pixel.GamepadStart:       "START",
		pixel.GamepadGuide:       "GUIDE",
		pixel.GamepadLeftThumb:   "LEFT THUMB",
		pixel.GamepadRightThumb:  "RIGHT THUMB",
		pixel.GamepadDpadUp:      "DPAD UP",
		pixel.GamepadDpadRight:   "DPAD RIGHT",
		pixel.GamepadDpadDown:    "DPAD DOWN",
		pixel.GamepadDpadLeft:    "DPAD LEFT",
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

type controllerOptionsController struct {
	inputController      input.InputController
	model                *controllerOptionsModel
	screenChangeRequired common.ScreenChangeCallback
}

func newControllerOptionsController(model *controllerOptionsModel) *controllerOptionsController {
	var result = controllerOptionsController{
		model: model,
	}
	return &result
}

func (coc *controllerOptionsController) setInputController(controller input.InputController) {
	coc.inputController = controller
}

func (coc *controllerOptionsController) setScreenChangeCallback(callback common.ScreenChangeCallback) {
	coc.screenChangeRequired = callback
}

func (coc *controllerOptionsController) setUp() {
	coc.model.availableDevices = coc.inputController.GetInputDeviceInfos()

	var inputConfigs = config.GetConfigService().LoadInputDeviceConfigurations()
	if len(inputConfigs) > coc.model.playerIdx {
		coc.model.inputConfig = inputConfigs[coc.model.playerIdx]
		for i, d := range coc.model.availableDevices {
			if d.DeviceName == coc.model.inputConfig.DeviceName {
				coc.model.selectedDeviceIndex = i
				break
			}
		}
	} else {
		coc.model.inputConfig.DeviceName = config.DeviceNameKeyboard
	}
}

func (coc *controllerOptionsController) update(window *opengl.Window) {
	if !coc.model.waitForInput {
		coc.processUserInput()
	} else {
		coc.assignSelectedButton(window)
	}
}

func (coc *controllerOptionsController) processUserInput() {
	var uiEventState = coc.inputController.GetUiEventStateCombined()
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
				if i == coc.model.selectedOption {
					coc.model.selectedOption = fc.nextSelection
					break focusHandling
				}
			}
		}
	}

	if uiEventState.PressedButton {
		coc.processOptionSelected()
	}
}

func (coc *controllerOptionsController) processOptionSelected() {
	switch coc.model.selectedOption {
	case optionControllerPreviousController:
		coc.selectPreviousController()
	case optionControllerNextController:
		coc.selectNextController()
	case optionControllerDigitalAxis:
		coc.model.inputConfig.GamepadConfiguration.HasDigitalAxis = !coc.model.inputConfig.GamepadConfiguration.HasDigitalAxis
	case optionControllerAction, optionControllerNextWeapon, optionControllerPreviousWeapon:
		coc.model.waitForInput = true
	case optionControllerSave:
		var err = config.GetConfigService().SaveInputDeviceConfiguration(coc.model.inputConfig, coc.model.playerIdx)
		if nil != err {
			logging.Warning.Printf("failed to save controller settings for player %d: %s", coc.model.playerIdx, err)
		}
		coc.screenChangeRequired(common.ConfigurationControls)
	case optionControllerBack:
		coc.screenChangeRequired(common.ConfigurationControls)
	default:
		logging.Error.Fatal("Unexpected selection in controllerOptionsController")
	}
}

func (coc *controllerOptionsController) assignSelectedButton(window *opengl.Window) {
	var selectedValue = -1
	if coc.model.inputConfig.DeviceName == config.DeviceNameKeyboard {
		for _, btn := range common.KeyboardButtons {
			if window.JustPressed(btn) {
				selectedValue = int(btn)
				break
			}
		}
	} else {
		for _, btn := range common.GamepadButtons {
			if window.JoystickJustPressed(pixel.Joystick(coc.model.inputConfig.GamepadConfiguration.JoystickIndex), btn) {
				selectedValue = int(btn)
				break
			}
		}
	}

	if selectedValue == -1 {
		return
	}

	switch coc.model.selectedOption {
	case optionControllerAction:
		coc.model.inputConfig.InputFire = selectedValue
		coc.model.waitForInput = false
	case optionControllerNextWeapon:
		coc.model.inputConfig.InputNextWeapon = selectedValue
		coc.model.waitForInput = false
	case optionControllerPreviousWeapon:
		coc.model.inputConfig.InputPreviousWeapon = selectedValue
		coc.model.waitForInput = false
	}
}

func (coc *controllerOptionsController) selectPreviousController() {
	coc.model.selectedDeviceIndex = coc.model.selectedDeviceIndex - 1
	if coc.model.selectedDeviceIndex < 0 {
		coc.model.selectedDeviceIndex = len(coc.model.availableDevices) - 1
	}
	coc.model.inputConfig.DeviceName = coc.model.availableDevices[coc.model.selectedDeviceIndex].DeviceName
	coc.model.inputConfig.JoystickIndex = coc.model.availableDevices[coc.model.selectedDeviceIndex].JoystickIndex
}

func (coc *controllerOptionsController) selectNextController() {
	coc.model.selectedDeviceIndex = (coc.model.selectedDeviceIndex + 1) % len(coc.model.availableDevices)
	coc.model.inputConfig.DeviceName = coc.model.availableDevices[coc.model.selectedDeviceIndex].DeviceName
	coc.model.inputConfig.JoystickIndex = coc.model.availableDevices[coc.model.selectedDeviceIndex].JoystickIndex
}

func (coc *controllerOptionsController) getDisplayTextForValue(value int) string {
	if coc.model.inputConfig.DeviceName == config.DeviceNameKeyboard {
		return pixel.Button(value).String()
	}

	var gpButton = pixel.GamepadButton(value)
	return gamepadButtonNames[gpButton]
}
