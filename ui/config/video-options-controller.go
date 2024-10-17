package config

import (
	"retro-carnage/config"
	"retro-carnage/input"
	"retro-carnage/logging"
	"retro-carnage/ui/common"

	"github.com/Retro-Carnage-Team/pixel2/backends/opengl"
)

const (
	minWindowWidth                   = 1024
	minWindowHeight                  = 768
	optionVideoUsePrimaryMonitor int = iota
	optionVideoUseOtherMonitor
	optionVideoPreviousMonitor
	optionVideoNextMonitor
	optionVideoFullscreen
	optionVideoWindowed
	optionVideoReduceWindowWidth
	optionVideoIncreaseWindowWidth
	optionVideoReduceWindowHeight
	optionVideoIncreaseWindowHeight
	optionVideoSave
	optionVideoBack
)

var (
	optionVideoFocusChanges = []focusChange{
		{movedDown: true, currentSelection: []int{optionVideoUsePrimaryMonitor}, nextSelection: optionVideoUseOtherMonitor},
		{movedUp: true, currentSelection: []int{optionVideoUseOtherMonitor}, nextSelection: optionVideoUsePrimaryMonitor},
		{movedRight: true, currentSelection: []int{optionVideoUseOtherMonitor}, nextSelection: optionVideoPreviousMonitor},
		{movedLeft: true, currentSelection: []int{optionVideoPreviousMonitor}, nextSelection: optionVideoUseOtherMonitor},
		{movedRight: true, currentSelection: []int{optionVideoPreviousMonitor}, nextSelection: optionVideoNextMonitor},
		{movedLeft: true, currentSelection: []int{optionVideoNextMonitor}, nextSelection: optionVideoPreviousMonitor},
		{movedUp: true, currentSelection: []int{optionVideoUseOtherMonitor, optionVideoPreviousMonitor, optionVideoNextMonitor}, nextSelection: optionVideoUsePrimaryMonitor},
		{movedDown: true, currentSelection: []int{optionVideoUseOtherMonitor, optionVideoPreviousMonitor, optionVideoNextMonitor}, nextSelection: optionVideoFullscreen},
		{movedRight: true, currentSelection: []int{optionVideoFullscreen}, nextSelection: optionVideoWindowed},
		{movedLeft: true, currentSelection: []int{optionVideoWindowed}, nextSelection: optionVideoFullscreen},
		{movedUp: true, currentSelection: []int{optionVideoFullscreen, optionVideoWindowed}, nextSelection: optionVideoUseOtherMonitor},
		{movedDown: true, currentSelection: []int{optionVideoFullscreen, optionVideoWindowed}, nextSelection: optionVideoReduceWindowWidth},
		{movedRight: true, currentSelection: []int{optionVideoReduceWindowWidth}, nextSelection: optionVideoIncreaseWindowWidth},
		{movedLeft: true, currentSelection: []int{optionVideoIncreaseWindowWidth}, nextSelection: optionVideoReduceWindowWidth},
		{movedUp: true, currentSelection: []int{optionVideoIncreaseWindowWidth, optionVideoReduceWindowWidth}, nextSelection: optionVideoFullscreen},
		{movedDown: true, currentSelection: []int{optionVideoIncreaseWindowWidth, optionVideoReduceWindowWidth}, nextSelection: optionVideoReduceWindowHeight},
		{movedRight: true, currentSelection: []int{optionVideoReduceWindowHeight}, nextSelection: optionVideoIncreaseWindowHeight},
		{movedLeft: true, currentSelection: []int{optionVideoIncreaseWindowHeight}, nextSelection: optionVideoReduceWindowHeight},
		{movedUp: true, currentSelection: []int{optionVideoIncreaseWindowHeight, optionVideoReduceWindowHeight}, nextSelection: optionVideoReduceWindowWidth},
		{movedDown: true, currentSelection: []int{optionVideoIncreaseWindowHeight, optionVideoReduceWindowHeight}, nextSelection: optionVideoSave},
		{movedRight: true, currentSelection: []int{optionVideoSave}, nextSelection: optionVideoBack},
		{movedLeft: true, currentSelection: []int{optionVideoBack}, nextSelection: optionVideoSave},
		{movedUp: true, currentSelection: []int{optionVideoSave, optionVideoBack}, nextSelection: optionVideoReduceWindowHeight},
	}
)

type videoOptionsController struct {
	inputController      input.InputController
	model                *videoOptionsModel
	screenChangeRequired common.ScreenChangeCallback
}

func newVideoOptionsController(model *videoOptionsModel) *videoOptionsController {
	var result = videoOptionsController{
		model: model,
	}

	result.initializeVideoConfig()
	return &result
}

func (voc *videoOptionsController) setInputController(controller input.InputController) {
	voc.inputController = controller
}

func (voc *videoOptionsController) setScreenChangeCallback(callback common.ScreenChangeCallback) {
	voc.screenChangeRequired = callback
}

func (voc *videoOptionsController) initializeVideoConfig() {
	var configuredMonitorFound = false
	for i, m := range opengl.Monitors() {
		if voc.model.videoConfig.SelectedMonitor != "" && m.Name() == voc.model.videoConfig.SelectedMonitor {
			voc.model.selectedMonitorIndex = i
			var x, y = m.Size()
			if voc.model.videoConfig.Width > int(x) {
				voc.model.videoConfig.Width = int(x)
			}
			if voc.model.videoConfig.Height > int(y) {
				voc.model.videoConfig.Height = int(y)
			}
			configuredMonitorFound = true
		}
	}

	if !configuredMonitorFound {
		voc.model.videoConfig.SelectedMonitor = opengl.Monitors()[0].Name()
		voc.model.selectedMonitorIndex = 0
		var x, y = opengl.Monitors()[0].Size()
		voc.model.videoConfig.Width = int(x)
		voc.model.videoConfig.Height = int(y)
	}
}

func (voc *videoOptionsController) update() {
	voc.processUserInput()
}

func (voc *videoOptionsController) processUserInput() {
	var uiEventState = voc.inputController.GetUiEventStateCombined()
	if nil == uiEventState {
		return
	}

focusHandling:
	for _, fc := range optionVideoFocusChanges {
		if fc.movedLeft == uiEventState.MovedLeft &&
			fc.movedRight == uiEventState.MovedRight &&
			fc.movedDown == uiEventState.MovedDown &&
			fc.movedUp == uiEventState.MovedUp {
			for _, i := range fc.currentSelection {
				if i == voc.model.selectedOption {
					voc.model.selectedOption = fc.nextSelection
					break focusHandling
				}
			}
		}
	}

	if uiEventState.PressedButton {
		voc.processOptionSelected()
	}
}

func (voc *videoOptionsController) processOptionSelected() {
	switch voc.model.selectedOption {
	case optionVideoUsePrimaryMonitor:
		voc.model.videoConfig.UsePrimaryMonitor = true
	case optionVideoUseOtherMonitor:
		voc.model.videoConfig.UsePrimaryMonitor = false
	case optionVideoPreviousMonitor:
		voc.selectPreviousMonitor()
	case optionVideoNextMonitor:
		voc.selectNextMonitor()
	case optionVideoFullscreen:
		voc.model.videoConfig.FullScreen = true
	case optionVideoWindowed:
		voc.model.videoConfig.FullScreen = false
	case optionVideoReduceWindowWidth:
		voc.decreaseScreenWidth()
	case optionVideoIncreaseWindowWidth:
		voc.increaseScreenWidth()
	case optionVideoReduceWindowHeight:
		voc.decreaseScreenHeight()
	case optionVideoIncreaseWindowHeight:
		voc.increaseScreenHeight()
	case optionVideoSave:
		err := config.GetConfigService().SaveVideoConfiguration(voc.model.videoConfig)
		if nil != err {
			logging.Warning.Printf("failed to save video settings: %s", err)
		}
		voc.screenChangeRequired(common.ConfigurationOptions)
	case optionVideoBack:
		voc.screenChangeRequired(common.ConfigurationOptions)
	default:
		logging.Error.Fatal("Unexpected selection in VideoOptionsScreen")
	}
}

func (voc *videoOptionsController) selectPreviousMonitor() {
	voc.model.selectedMonitorIndex = voc.model.selectedMonitorIndex - 1
	if voc.model.selectedMonitorIndex < 0 {
		voc.model.selectedMonitorIndex = len(opengl.Monitors()) - 1
	}
	voc.model.videoConfig.SelectedMonitor = opengl.Monitors()[voc.model.selectedMonitorIndex].Name()
}

func (voc *videoOptionsController) selectNextMonitor() {
	voc.model.selectedMonitorIndex = (voc.model.selectedMonitorIndex + 1) % len(opengl.Monitors())
	voc.model.videoConfig.SelectedMonitor = opengl.Monitors()[voc.model.selectedMonitorIndex].Name()
}

func (voc *videoOptionsController) decreaseScreenWidth() {
	if voc.model.videoConfig.Width > minWindowWidth+10 {
		voc.model.videoConfig.Width = voc.model.videoConfig.Width - 10
	}
}

func (voc *videoOptionsController) increaseScreenWidth() {
	voc.model.videoConfig.Width = voc.model.videoConfig.Width + 10
	var x, _ = opengl.Monitors()[voc.model.selectedMonitorIndex].Size()
	if voc.model.videoConfig.Width > int(x) {
		voc.model.videoConfig.Width = int(x)
	}
}

func (voc *videoOptionsController) decreaseScreenHeight() {
	if voc.model.videoConfig.Height > minWindowHeight+10 {
		voc.model.videoConfig.Height = voc.model.videoConfig.Height - 10
	}
}

func (voc *videoOptionsController) increaseScreenHeight() {
	voc.model.videoConfig.Height = voc.model.videoConfig.Height + 10
	var _, y = opengl.Monitors()[voc.model.selectedMonitorIndex].Size()
	if voc.model.videoConfig.Height > int(y) {
		voc.model.videoConfig.Height = int(y)
	}
}
