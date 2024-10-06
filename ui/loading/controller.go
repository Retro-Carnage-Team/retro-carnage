package loading

import (
	"retro-carnage/assets"
	"retro-carnage/ui/common"
)

const screenTimeout = 8_500

type controller struct {
	screenChangeRequired common.ScreenChangeCallback
	screenChangeTimeout  int64
}

func newController() *controller {
	var result = controller{
		screenChangeTimeout: 0,
	}
	result.init()
	return &result
}

func (c *controller) setScreenChangeCallback(callback common.ScreenChangeCallback) {
	c.screenChangeRequired = callback
}

func (c *controller) update(elapsedTimeInMs int64) {
	c.screenChangeTimeout += elapsedTimeInMs
	if c.screenChangeTimeout >= screenTimeout && c.isInitDone() {
		c.screenChangeRequired(common.Start)
	}
}

func (c *controller) init() {
	var stereo = assets.NewStereo()
	stereo.PlayFx(assets.FxLoading)

	assets.AmmunitionCrate.Initialize()
	assets.GrenadeCrate.Initialize()
	assets.WeaponCrate.Initialize()
	assets.SpriteRepository.Initialize()
}

func (c *controller) isInitDone() bool {
	return assets.AmmunitionCrate.Initialized() &&
		assets.GrenadeCrate.Initialized() &&
		assets.WeaponCrate.Initialized() &&
		assets.SpriteRepository.Initialized()
}
