package common

import (
	"retro-carnage/assets"
	"retro-carnage/engine/characters"
)

// LoadingScreenInit is called when the loading screen is shown. A good place to start longer running background
// tasks.
func LoadingScreenInit() {
	assets.AmmunitionCrate.Initialize()
	assets.GrenadeCrate.Initialize()
	assets.WeaponCrate.Initialize()
	assets.SpriteRepository.Initialize()
}

// LoadingScreenInitDone returns true when all initialization steps of the loading screen are done.
func LoadingScreenInitDone() bool {
	return assets.AmmunitionCrate.Initialized() &&
		assets.GrenadeCrate.Initialized() &&
		assets.WeaponCrate.Initialized() &&
		assets.SpriteRepository.Initialized()
}

// StartScreenInit is called when the start screen is shown.
func StartScreenInit() {
	go characters.InitEnemySkins("skins")
	go characters.InitPlayerSkins("skins")

	var stereo = assets.NewStereo()
	go stereo.BufferSong(assets.GameOverSong)
	go stereo.BufferSong(assets.GameWonSong)
}

// TitleScreenInit is called when the title screen is shown.
func TitleScreenInit() {
	// currently there's nothing being done at this point
}
