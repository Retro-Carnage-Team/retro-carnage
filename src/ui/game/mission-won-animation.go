package game

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"retro-carnage/assets"
)

const (
	victorySongDelay       int64 = 500
	backgroundFadeDelay    int64 = 1750
	backgroundFadeDuration int64 = 250
)

type missionWonAnimation struct {
	backgroundCanvas    *pixelgl.Canvas
	backgroundColorMask pixel.RGBA
	duration            int64
	finished            bool
	mission             *assets.Mission
	stereo              *assets.Stereo
	window              *pixelgl.Window
}

func createMissionWonAnimation(
	playerInfos []*playerInfo,
	gameCanvas *pixelgl.Canvas,
	mission *assets.Mission,
	window *pixelgl.Window,
) *missionWonAnimation {
	var bgCanvas = pixelgl.NewCanvas(window.Bounds())
	for _, playerInfo := range playerInfos {
		playerInfo.draw(bgCanvas)
	}
	gameCanvas.Draw(bgCanvas, pixel.IM.Moved(gameCanvas.Bounds().Center()))

	return &missionWonAnimation{
		backgroundCanvas:    bgCanvas,
		backgroundColorMask: pixel.RGBA{A: 0.0},
		duration:            0,
		finished:            false,
		mission:             mission,
		stereo:              assets.NewStereo(),
		window:              window,
	}
}

func (mwa *missionWonAnimation) update(elapsedTimeInMs int64) {
	if 0 == mwa.duration {
		mwa.initialActions()
	}
	if mwa.duration > backgroundFadeDelay && mwa.duration <= backgroundFadeDelay+backgroundFadeDuration {
		var elapsed = float64(mwa.duration - backgroundFadeDelay)
		var total = float64(backgroundFadeDuration)
		var alpha = 0.3 * (elapsed / total)
		mwa.backgroundColorMask = pixel.RGBA{A: alpha}
	}

	// TODO: Set the finished flag only when the score calculation has been shown (or user cancelled the animation)
	if mwa.duration > 7500 {
		mwa.finished = true
	}

	mwa.duration += elapsedTimeInMs
}

func (mwa *missionWonAnimation) drawToScreen() {
	var matrix = pixel.IM.Moved(mwa.window.Bounds().Center())
	mwa.backgroundCanvas.DrawColorMask(mwa.window, matrix, mwa.backgroundColorMask)
}

func (mwa *missionWonAnimation) initialActions() {
	mwa.stereo.StopSong(mwa.mission.Music)
	// TODO: start victory music
}
