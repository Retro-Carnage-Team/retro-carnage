package title

import (
	"retro-carnage/assets"
	"retro-carnage/engine/cheat"
	"retro-carnage/input"
	"retro-carnage/ui/common"

	"github.com/Retro-Carnage-Team/pixel2/backends/opengl"
)

const screenTimeout = 60_000

type controller struct {
	cheatController      *cheat.CheatController
	inputController      input.InputController
	screenChangeRequired common.ScreenChangeCallback
	screenChangeTimeout  int64
	stereo               *assets.Stereo
}

func newController() *controller {
	var result = controller{
		cheatController: cheat.GetCheatController(),
		stereo:          assets.NewStereo(),
	}
	return &result
}

func (c *controller) setInputController(controller input.InputController) {
	c.inputController = controller
}

func (c *controller) setScreenChangeCallback(callback common.ScreenChangeCallback) {
	c.screenChangeRequired = callback
}

func (c *controller) update(elapsedTimeInMs int64, window *opengl.Window) {
	c.screenChangeTimeout += elapsedTimeInMs
	for _, btn := range common.KeyboardButtons {
		if window.JustPressed(btn) {
			if c.cheatController.HandleKeyboardInput(btn) {
				c.stereo.PlayFx(assets.FxCheatSwitch)
			}
			break
		}
	}

	var uiEventState = c.inputController.GetUiEventStateCombined()
	if (nil != uiEventState && uiEventState.PressedButton) || screenTimeout <= c.screenChangeTimeout {
		c.screenChangeRequired(common.ConfigurationSelect)
	}
}
