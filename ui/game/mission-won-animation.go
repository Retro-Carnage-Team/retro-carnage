package game

import (
	"fmt"
	"retro-carnage/assets"
	"retro-carnage/engine/characters"
	"retro-carnage/ui/common"
	"retro-carnage/ui/common/fonts"

	pixel "github.com/Retro-Carnage-Team/pixel2"
	"github.com/Retro-Carnage-Team/pixel2/backends/opengl"
)

type missionWonAnimationSection int

const (
	animationSectionOne      missionWonAnimationSection = 0
	animationSectionTwo      missionWonAnimationSection = 1
	animationSectionThree    missionWonAnimationSection = 2
	backgroundFadeDelay      int64                      = 1_000
	backgroundFadeDuration   int64                      = 250
	bonusIncrementDuration   int64                      = 165
	bonusIncrement           int64                      = 1_000
	missionBonusDelay        int64                      = 2_500
	missionWonOpacity        float64                    = 0.5
	postAnimationDelay       int64                      = 5_000
	postMissionDelay         int64                      = 2_500
	revengeBonusPerKill      int64                      = 10
	revengeIncrementDuration int64                      = 100
)

var (
	completedTextLines = []string{"CONGRATULATIONS!", "YOU HAVE COMPLETED YOUR MISSION"}
)

type missionWonAnimation struct {
	animationSection     missionWonAnimationSection
	backgroundCanvas     *opengl.Canvas
	backgroundColorMask  pixel.RGBA
	completedTextVisible bool
	duration             int64
	finished             bool
	kills                []int
	killCounter          []int
	killCounterDuration  int64
	mission              *assets.Mission
	missionBonus         int64
	missionBonusDuration int64
	playerBonus          []int64
	playerResultLines    []string
	stereo               *assets.Stereo
	window               *opengl.Window
}

func createMissionWonAnimation(
	playerInfos []*playerInfo,
	gameCanvas *opengl.Canvas,
	kills []int,
	mission *assets.Mission,
	window *opengl.Window,
) *missionWonAnimation {
	var bgCanvas = opengl.NewCanvas(window.Bounds())
	for _, playerInfo := range playerInfos {
		playerInfo.draw(bgCanvas)
	}
	gameCanvas.Draw(bgCanvas, pixel.IM.Moved(gameCanvas.Bounds().Center()))

	return &missionWonAnimation{
		animationSection:     animationSectionOne,
		backgroundCanvas:     bgCanvas,
		backgroundColorMask:  pixel.RGBA{A: 0.0},
		completedTextVisible: false,
		duration:             0,
		finished:             false,
		kills:                kills,
		killCounter:          []int{0, 0},
		killCounterDuration:  0,
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
	if mwa.duration == 0 {
		mwa.initialActions()
	}

	switch mwa.animationSection {
	case animationSectionOne:
		mwa.runAnimationSectionOne()
	case animationSectionTwo:
		mwa.runAnimationSectionTwo(elapsedTimeInMs)
	case animationSectionThree:
		mwa.runAnimationSectionThree(elapsedTimeInMs)
	}

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
	for _, player := range characters.PlayerController.RemainingPlayers() {
		if player.AutomaticWeaponSelected() {
			mwa.stereo.StopFx(player.SelectedWeapon().Sound)
		}
	}
	mwa.stereo.StopSong(mwa.mission.Music)
}

func (mwa *missionWonAnimation) showTexts() {
	if mwa.killCounterDuration < postAnimationDelay/2 {
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
}

func (mwa *missionWonAnimation) isKillCounterDone() bool {
	var done = true
	for _, player := range characters.PlayerController.RemainingPlayers() {
		var playerIndex = player.Index()
		done = done && (mwa.kills[playerIndex] == mwa.killCounter[playerIndex])
	}
	return done
}

func (mwa *missionWonAnimation) runAnimationSectionOne() {
	if (mwa.duration > backgroundFadeDelay) && (mwa.duration <= backgroundFadeDelay+backgroundFadeDuration) {
		var elapsed = float64(mwa.duration - backgroundFadeDelay)
		var total = float64(backgroundFadeDuration)
		var alpha = missionWonOpacity * (elapsed / total)
		mwa.backgroundColorMask = pixel.RGBA{A: alpha}
	}
	mwa.completedTextVisible = mwa.duration > backgroundFadeDelay+backgroundFadeDuration/2

	if mwa.duration >= missionBonusDelay {
		mwa.animationSection = animationSectionTwo
	}
}

func (mwa *missionWonAnimation) runAnimationSectionTwo(elapsedTimeInMs int64) {
	mwa.missionBonusDuration += elapsedTimeInMs

	if (mwa.duration >= missionBonusDelay) &&
		(mwa.missionBonus <= int64(mwa.mission.Reward)) &&
		(mwa.missionBonusDuration < 2*bonusIncrementDuration) {

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
		(mwa.missionBonusDuration > postMissionDelay) {
		mwa.animationSection = animationSectionThree
	}
}

func (mwa *missionWonAnimation) runAnimationSectionThree(elapsedTimeInMs int64) {
	mwa.killCounterDuration += elapsedTimeInMs
	if mwa.isKillCounterDone() {
		mwa.stereo.StopFx(assets.FxAr10)
		var progress = float64(mwa.killCounterDuration) / float64(postAnimationDelay)
		var bgAlpha = missionWonOpacity + (1-missionWonOpacity)*progress
		mwa.backgroundColorMask = pixel.RGBA{A: bgAlpha}
		if mwa.killCounterDuration > postAnimationDelay/2 {
			mwa.playerResultLines[0] = ""
			mwa.playerResultLines[1] = ""
		}
		mwa.finished = mwa.killCounterDuration >= postAnimationDelay
	} else if mwa.killCounterDuration > revengeIncrementDuration {
		if (mwa.killCounter[0] == 0) && (mwa.killCounter[1] == 0) {
			mwa.stereo.PlayFx(assets.FxAr10)
		}
		mwa.killCounterDuration = 0
		for _, player := range characters.PlayerController.RemainingPlayers() {
			if mwa.killCounter[player.Index()] < mwa.kills[player.Index()] {
				mwa.killCounter[player.Index()] += 1
				mwa.playerBonus[player.Index()] = mwa.playerBonus[player.Index()] + revengeBonusPerKill
			}
			mwa.playerResultLines[player.Index()] = fmt.Sprintf(
				"PLAYER %d - REVENGE BONUS x %d: $%7d",
				1+player.Index(),
				mwa.kills[player.Index()]-mwa.killCounter[player.Index()],
				mwa.playerBonus[player.Index()],
			)
		}
	}
}
