package game

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"retro-carnage/assets"
	"retro-carnage/engine/characters"
	"retro-carnage/ui/common"
	"retro-carnage/ui/common/fonts"
)

const (
	backgroundFadeDelay    int64 = 1000
	backgroundFadeDuration int64 = 250
	bonusIncrementDuration int64 = 165
	bonusIncrement         int64 = 500
	missionBonusDelay      int64 = 1500
	// victorySongDelay    int64 = 500
)

var (
	completedTextLines = []string{"CONGRATULATIONS!", "YOU HAVE COMPLETED YOUR MISSION"}
)

type missionWonAnimation struct {
	backgroundCanvas     *pixelgl.Canvas
	backgroundColorMask  pixel.RGBA
	completedTextVisible bool
	duration             int64
	finished             bool
	mission              *assets.Mission
	missionBonus         int64
	missionBonusDuration int64
	playerBonus          []int64
	playerResultLines    []string
	stereo               *assets.Stereo
	window               *pixelgl.Window
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
		backgroundCanvas:     bgCanvas,
		backgroundColorMask:  pixel.RGBA{A: 0.0},
		completedTextVisible: false,
		duration:             0,
		finished:             false,
		mission:              mission,
		missionBonus:         0,
		missionBonusDuration: 0,
		playerBonus:          []int64{0, 0},
		playerResultLines:    []string{"", ""},
		stereo:               assets.NewStereo(),
		window:               window,
	}
}

func (mwa *missionWonAnimation) update(elapsedTimeInMs int64) {
	if 0 == mwa.duration {
		mwa.initialActions()
	}

	if (mwa.duration > backgroundFadeDelay) && (mwa.duration <= backgroundFadeDelay+backgroundFadeDuration) {
		var elapsed = float64(mwa.duration - backgroundFadeDelay)
		var total = float64(backgroundFadeDuration)
		var alpha = 0.3 * (elapsed / total)
		mwa.backgroundColorMask = pixel.RGBA{A: alpha}
	}

	mwa.completedTextVisible = mwa.duration > backgroundFadeDelay+backgroundFadeDuration/2
	if (mwa.duration >= missionBonusDelay) &&
		(mwa.missionBonus <= int64(mwa.mission.Reward)) &&
		(mwa.missionBonusDuration < 2*bonusIncrementDuration) {

		mwa.missionBonusDuration += elapsedTimeInMs
		if (mwa.missionBonus < int64(mwa.mission.Reward)) && (mwa.missionBonusDuration > bonusIncrementDuration) {
			mwa.missionBonusDuration = 0
			mwa.missionBonus += bonusIncrement
			mwa.playerBonus[0] = mwa.missionBonus
			mwa.playerBonus[1] = mwa.missionBonus
			mwa.stereo.PlayFx(assets.FxPistol1)
		}

		for _, player := range characters.PlayerController.RemainingPlayers() {
			mwa.playerResultLines[player.Index()] = fmt.Sprintf(
				"PLAYER %d - MISSION BONUS: $%7d", 1+player.Index(), mwa.playerBonus[player.Index()])
		}
	}

	if (mwa.duration >= missionBonusDelay) && (mwa.missionBonus == int64(mwa.mission.Reward)) &&
		(mwa.missionBonusDuration > 2*bonusIncrementDuration) {
		mwa.playerResultLines[0] = "Done"
		mwa.playerResultLines[1] = "Done"
	}

	// TODO: Set the finished flag only when the score calculation has been shown (or user cancelled the animation)
	// if mwa.duration > 7500 {
	//  	mwa.finished = true
	// }

	mwa.duration += elapsedTimeInMs
}

func (mwa *missionWonAnimation) drawToScreen() {
	var matrix = pixel.IM.Moved(mwa.window.Bounds().Center())
	mwa.backgroundCanvas.DrawColorMask(mwa.window, matrix, mwa.backgroundColorMask)

	if mwa.completedTextVisible {
		mwa.showTexts()
	}
}

func (mwa *missionWonAnimation) initialActions() {
	mwa.stereo.StopSong(mwa.mission.Music)
	// TODO: start victory music
}

func (mwa *missionWonAnimation) showTexts() {
	var offset = 6.0
	var renderer = fonts.TextRenderer{Window: mwa.window}
	for _, line := range completedTextLines {
		renderer.DrawLineToScreenCenter(line, offset, common.White)
		offset -= 1.5
	}

	offset -= 2
	for _, line := range mwa.playerResultLines {
		renderer.DrawLineToScreenCenter(line, offset, common.White)
		offset -= 1.5
	}
}
