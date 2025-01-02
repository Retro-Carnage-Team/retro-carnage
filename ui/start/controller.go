package start

import (
	"retro-carnage/assets"
	"retro-carnage/engine/graphics"
	"retro-carnage/input"
	"retro-carnage/logging"
	"retro-carnage/ui/common"
)

const MIN_SCREEN_DURATION = 7_500

type controller struct {
	inputController      input.InputController
	screenChangeRequired common.ScreenChangeCallback
	screenChangeTimeout  int64
	stereo               *assets.Stereo
	themeLoaded          bool
}

func newController() *controller {
	var controller = controller{
		screenChangeTimeout: 0,
		stereo:              assets.NewStereo(),
		themeLoaded:         false,
	}

	controller.init()
	return &controller
}

func (c *controller) setInputController(i input.InputController) {
	c.inputController = i
}

func (c *controller) setScreenChangeCallback(callback common.ScreenChangeCallback) {
	c.screenChangeRequired = callback
}

func (c *controller) update(elapsedTimeInMs int64) {
	c.screenChangeTimeout += elapsedTimeInMs

	if !c.themeLoaded {
		c.themeLoaded = c.stereo.IsSongBuffered(assets.ThemeSong)
	}

	if c.themeLoaded {
		if nil == c.screenChangeRequired {
			logging.Error.Fatalf("No ScreenChangeCallback set in start.controller")
		}

		var uiEventState = c.inputController.GetUiEventStateCombined()
		if (nil != uiEventState && uiEventState.PressedButton) || c.screenChangeTimeout > MIN_SCREEN_DURATION {
			c.screenChangeRequired(common.Title)
			if c.themeLoaded {
				c.stereo.PlaySong(assets.ThemeSong)
			}
		}
	}
}

func (c *controller) init() {
	go graphics.InitEnemySkins("skins")
	go graphics.InitPlayerSkins("skins")

	go c.stereo.BufferSong(assets.ThemeSong)
	go c.stereo.BufferSong(assets.GameOverSong)
	go c.stereo.BufferSong(assets.GameWonSong)
}
