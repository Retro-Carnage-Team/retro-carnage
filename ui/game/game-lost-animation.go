package game

import (
	"retro-carnage/assets"
	"retro-carnage/engine/characters"
	"retro-carnage/ui/common"
	"retro-carnage/ui/common/fonts"

	"github.com/Retro-Carnage-Team/pixel"
	"github.com/Retro-Carnage-Team/pixel/pixelgl"
)

type gameLostAnimation struct {
	backgroundCanvas    *pixelgl.Canvas
	backgroundColorMask pixel.RGBA
	duration            int64
	finished            bool
	mission             *assets.Mission
	stereo              *assets.Stereo
	textVisible         bool
	window              *pixelgl.Window
}

const (
	gameLostOpacity   = 0.5
	postGameOverDelay = 60_000
)

var gameOverTextLines = []string{"GAME OVER", "GIVE IT ANOTHER TRY - AND MAKE AN EFFORT THIS TIME!"}

func createGameLostAnimation(
	playerInfos []*playerInfo,
	gameCanvas *pixelgl.Canvas,
	mission *assets.Mission,
	window *pixelgl.Window,
) *gameLostAnimation {
	var bgCanvas = pixelgl.NewCanvas(window.Bounds())
	for _, playerInfo := range playerInfos {
		playerInfo.draw(bgCanvas)
	}
	gameCanvas.Draw(bgCanvas, pixel.IM.Moved(gameCanvas.Bounds().Center()))
	return &gameLostAnimation{
		backgroundCanvas:    bgCanvas,
		backgroundColorMask: pixel.RGBA{A: 0.0},
		duration:            0,
		finished:            false,
		mission:             mission,
		stereo:              assets.NewStereo(),
		textVisible:         false,
		window:              window,
	}
}

func (gla *gameLostAnimation) update(elapsedTimeInMs int64) {
	if 0 == gla.duration {
		gla.initialActions()
	}

	gla.duration += elapsedTimeInMs

	if (gla.duration > backgroundFadeDelay) && (gla.duration <= backgroundFadeDelay+backgroundFadeDuration) {
		var elapsed = float64(gla.duration - backgroundFadeDelay)
		var total = float64(backgroundFadeDuration)
		var alpha = gameLostOpacity * (elapsed / total)
		gla.backgroundColorMask = pixel.RGBA{A: alpha}
	}

	gla.textVisible = gla.duration > backgroundFadeDelay+backgroundFadeDuration/2
	gla.finished = gla.duration > postGameOverDelay
}

func (gla *gameLostAnimation) drawToScreen() {
	var matrix = pixel.IM.Moved(gla.window.Bounds().Center())
	gla.backgroundCanvas.DrawColorMask(gla.window, matrix, gla.backgroundColorMask)
	if gla.textVisible {
		gla.showTexts()
	}
}

func (gla *gameLostAnimation) initialActions() {
	for _, player := range characters.PlayerController.RemainingPlayers() {
		if player.AutomaticWeaponSelected() {
			gla.stereo.StopFx(player.SelectedWeapon().Sound)
		}
	}
	gla.stereo.StopSong(gla.mission.Music)
	gla.stereo.PlaySong(assets.GameOverSong)
}

func (gla *gameLostAnimation) showTexts() {
	var renderer = fonts.TextRenderer{Window: gla.window}
	renderer.DrawLineToScreenCenter(gameOverTextLines[0], 1.5, common.White)
	renderer.DrawLineToScreenCenter(gameOverTextLines[1], -1.5, common.White)
}
