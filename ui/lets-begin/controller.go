package lets_begin

import (
	"retro-carnage/assets"
	"retro-carnage/engine"
	"retro-carnage/ui/common"
)

const (
	timeAfterLastChar       = 500
	timeBetweenChars        = 120
	timeBetweenVolumeChange = 150
)

type controller struct {
	characterTimer       int64
	model                model
	screenChangeRequired common.ScreenChangeCallback
	stereo               *assets.Stereo
	textLength           int
	volumeTimer          int64
}

func newController() *controller {
	var result = controller{
		stereo: assets.NewStereo(),
	}
	return &result
}

func (c *controller) setScreenChangeCallback(callback common.ScreenChangeCallback) {
	c.screenChangeRequired = callback
}

func (c *controller) update(elapsedTimeInMs int64) {
	c.characterTimer += elapsedTimeInMs
	c.volumeTimer += elapsedTimeInMs
	if c.textLength < len(displayText) {
		// text has not been fully typed
		if c.characterTimer >= timeBetweenChars {
			c.textLength++
			c.model.text = displayText[:c.textLength]
			c.characterTimer = 0
		}
		if c.volumeTimer >= timeBetweenVolumeChange {
			c.stereo.DecreaseVolume(assets.ThemeSong)
			c.volumeTimer = 0
		}
	} else if c.isMissionInitialized() {
		// text has been typed and initialization is completed
		c.screenChangeRequired(common.Game)
		c.stereo.StopSong(assets.ThemeSong)
	} else {
		// text has been typed - but initialization is not completed
		c.textLength = c.textLength - 3
		c.model.text = displayText[:c.textLength]
		c.characterTimer = 0
	}
}

func (c *controller) isMissionInitialized() bool {
	var music = engine.MissionController.CurrentMission().Music
	return c.stereo.IsSongBuffered(music)
}
