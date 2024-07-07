package game

import (
	"math"
	"retro-carnage/assets"
	"retro-carnage/engine/characters"
	"retro-carnage/ui/common"
	"retro-carnage/ui/common/fonts"

	pixel "github.com/Retro-Carnage-Team/pixel2"
	"github.com/Retro-Carnage-Team/pixel2/backends/opengl"
)

type gameWonAnimation struct {
	backgroundSprite *pixel.Sprite
	backgroundMatrix pixel.Matrix
	duration         int64
	finished         bool
	mission          *assets.Mission
	stereo           *assets.Stereo
	window           *opengl.Window
}

const (
	gameWonBackground = "images/other/sunset-soldiers.jpg"
	postGameWonDelay  = 170_000
)

var gameWonTextLines = []string{
	"CONGRATULATIONS",
	"",
	"YOU'VE MADE IT! THE EVIL OF THE WORLD HAS",
	"BEEN DESTROYED! DICTATORS HAVE BEEN",
	"OVERTHROWN, VILLAINS CRUSHED, DESPOTS REMOVED.",
	"",
	"BUT IT'S NOT YET TIME FOR A VACATION. THERE ARE",
	"OTHER PLANETS FULL OF INJUSTICE. IT'S TIME FOR...",
	"",
	"RETRO-CARNAGE IN SPACE !",
}

func createGameWonAnimation(
	mission *assets.Mission,
	window *opengl.Window,
) *gameWonAnimation {
	var backgroundImage = assets.SpriteRepository.Get(gameWonBackground)
	var scaleX = window.Bounds().Max.X / backgroundImage.Picture().Bounds().Max.X
	var scaleY = window.Bounds().Max.Y / backgroundImage.Picture().Bounds().Max.Y
	var scale = math.Max(scaleX, scaleY)
	var matrix = pixel.IM.Scaled(pixel.V(0, 0), scale).Moved(window.Bounds().Center())

	return &gameWonAnimation{
		backgroundMatrix: matrix,
		backgroundSprite: backgroundImage,
		duration:         0,
		finished:         false,
		mission:          mission,
		stereo:           assets.NewStereo(),
		window:           window,
	}
}

func (gwa *gameWonAnimation) update(elapsedTimeInMs int64) {
	if gwa.duration == 0 {
		gwa.initialActions()
	}

	gwa.duration += elapsedTimeInMs
	gwa.finished = gwa.duration > postGameWonDelay
}

func (gwa *gameWonAnimation) drawToScreen() {
	gwa.backgroundSprite.Draw(gwa.window, gwa.backgroundMatrix)
	gwa.showTexts()
}

func (gwa *gameWonAnimation) initialActions() {
	for _, player := range characters.PlayerController.RemainingPlayers() {
		if player.AutomaticWeaponSelected() {
			gwa.stereo.StopFx(player.SelectedWeapon().Sound)
		}
	}
	gwa.stereo.StopSong(gwa.mission.Music)
	gwa.stereo.PlaySong(assets.GameWonSong)
}

func (gwa *gameWonAnimation) showTexts() {
	var renderer = fonts.TextRenderer{Window: gwa.window}
	var offset = 10.0
	for _, line := range gameWonTextLines {
		renderer.DrawLineToScreenCenter(line, offset, common.White)
		offset -= 1.5
	}
}
