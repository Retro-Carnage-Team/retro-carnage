package game

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"math"
	"retro-carnage/assets"
	"retro-carnage/ui/common"
	"retro-carnage/ui/common/fonts"
)

type gameWonAnimation struct {
	backgroundSprite *pixel.Sprite
	backgroundMatrix pixel.Matrix
	duration         int64
	finished         bool
	mission          *assets.Mission
	showText         bool
	stereo           *assets.Stereo
	window           *pixelgl.Window
}

const (
	gameWonBackground = "images/other/sunset-soldiers.jpg"
	postGameWonDelay  = 170_000
)

var gameWonTextLines = []string{
	"CONGRATULATIONS",
	"",
	"YOU HAVE MADE IT! THE EVIL OF THE WORLD HAS",
	"SUFFERED A MAJOR SETBACK. FRIGHTENED VILLAINS",
	"ARE HIDING TREMBLING IN CAVES, DICTATORS ARE",
	"ASKING THEIR PEOPLES FOR FORGIVENESS,",
	"RETRIBUTION FOR INNOCENT VICTIMS HAS BEEN TAKEN.",
	"",
	"BUT IT IS NOT YET TIME FOR A VACATION. THERE ARE",
	"OTHER PLANETS FULL OF INJUSTICE. IT IS TIME FOR...",
	"",
	"RETRO-CARNAGE IN SPACE",
}

func createGameWonAnimation(
	mission *assets.Mission,
	window *pixelgl.Window,
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
	if 0 == gwa.duration {
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
