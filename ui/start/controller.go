package start

import (
	"retro-carnage/assets"
	"retro-carnage/logging"
	"retro-carnage/ui/common"
)

const MIN_SCREEN_DURATION = 1_000

type controller struct {
	screenChangeRequired common.ScreenChangeCallback
	screenChangeTimeout  int64
	stereo               *assets.Stereo
	themeLoaded          bool
}

func NewController() *controller {
	var controller = controller{
		screenChangeTimeout: 0,
		stereo:              assets.NewStereo(),
		themeLoaded:         false,
	}

	common.StartScreenInit()
	go controller.stereo.BufferSong(assets.ThemeSong)
	return &controller
}

func (c *controller) SetScreenChangeCallback(callback common.ScreenChangeCallback) {
	c.screenChangeRequired = callback
}

func (c *controller) Update(elapsedTimeInMs int64) {
	c.screenChangeTimeout += elapsedTimeInMs

	if !c.themeLoaded {
		c.themeLoaded = c.stereo.IsSongBuffered(assets.ThemeSong)
		if c.themeLoaded {
			c.stereo.PlaySong(assets.ThemeSong)
		}
	}

	if c.themeLoaded {
		if nil == c.screenChangeRequired {
			logging.Error.Fatalf("No ScreenChangeCallback set in start.controller")
		}

		if c.screenChangeTimeout > MIN_SCREEN_DURATION {
			c.screenChangeRequired(common.Title)
		}
	}
}
