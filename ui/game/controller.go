package game

import (
	"retro-carnage/assets"
	"retro-carnage/engine"
	"retro-carnage/input"
	"retro-carnage/ui/common"
	"retro-carnage/ui/highscore"

	pixel "github.com/Retro-Carnage-Team/pixel2"
	"github.com/Retro-Carnage-Team/pixel2/backends/opengl"
)

type controller struct {
	engine               *engine.GameEngine
	gameLostAnimation    *gameLostAnimation
	gameWonAnimation     *gameWonAnimation
	missionWonAnimation  *missionWonAnimation
	inputController      input.InputController
	model                *model
	screenChangeRequired common.ScreenChangeCallback
	stereo               *assets.Stereo
}

// SetInputController is used to connect Screen with the global input.Controller instance.
func (c *controller) setInputController(ctrl input.InputController) {
	c.inputController = ctrl

}

// SetScreenChangeCallback is used to connect Screen with the callback method of ui.MainScreen.
func (c *controller) setScreenChangeCallback(callback common.ScreenChangeCallback) {
	c.screenChangeRequired = callback
}

func (c *controller) setUp() {
	c.engine.SetInputController(c.inputController)
	c.stereo.PlaySong(c.model.mission.Music)
}

func (c *controller) tearDown(window *opengl.Window) {
	c.stereo.StopSong(c.model.mission.Music)
	window.SetTitle("RETRO CARNAGE")
}

func (c *controller) update(elapsedTimeInMs int64, window *opengl.Window) {
	var uiEventStateCombined = c.inputController.GetUiEventStateCombined()
	var buttonPressed = (nil != uiEventStateCombined) && uiEventStateCombined.PressedButton

	if c.gameInProgress() {
		c.updateGameInProgress(window, elapsedTimeInMs)
	} else if c.engine.Won {
		c.updateGameWon(elapsedTimeInMs, buttonPressed, window)
	} else if c.engine.Lost {
		c.updateGameLost(elapsedTimeInMs, buttonPressed)
	}
}

func (c *controller) updateGameLost(elapsedTimeInMs int64, buttonPressed bool) {
	if nil != c.gameLostAnimation {
		c.gameLostAnimation.update(elapsedTimeInMs)
		if c.gameLostAnimation.finished || buttonPressed {
			c.moveToHighScoreScreen()
		}
	}
}

func (c *controller) updateGameWon(elapsedTimeInMs int64, buttonPressed bool, window *opengl.Window) {
	if nil != c.gameWonAnimation {
		c.gameWonAnimation.update(elapsedTimeInMs)
		if c.gameWonAnimation.finished || buttonPressed {
			c.onMissionWon()
		}
	} else if nil != c.missionWonAnimation {
		c.missionWonAnimation.update(elapsedTimeInMs)
		if c.missionWonAnimation.finished || buttonPressed {
			var remainingMissions, _ = engine.MissionController.RemainingMissions()
			// The current mission has not been marked as won, yet. Thus, there is one remaining mission.
			if (len(remainingMissions) == 1) && (remainingMissions[0].Name == c.model.mission.Name) {
				c.gameWonAnimation = createGameWonAnimation(c.model.mission, window)
			} else {
				c.onMissionWon()
			}
		}
	}
}

func (c *controller) updateGameInProgress(window *opengl.Window, elapsedTimeInMs int64) {
	if window.JustReleased(pixel.KeyPause) {
		c.model.paused = !c.model.paused
	}
	if !c.model.paused {
		c.engine.UpdateGameState(elapsedTimeInMs)
		c.model.inProgress = c.gameInProgress()
		c.model.lost = c.engine.Lost
		c.model.won = c.engine.Won
	}
}

func (c *controller) gameInProgress() bool {
	return nil != c.engine && !(c.engine.Won || c.engine.Lost)
}

func (c *controller) createGameAnimations(
	playerInfos []*playerInfo,
	gameCanvas *opengl.Canvas,
	window *opengl.Window,
) {
	if c.engine.Lost {
		c.gameLostAnimation = createGameLostAnimation(
			playerInfos,
			gameCanvas,
			c.model.mission,
			window,
		)
	} else if c.engine.Won {
		c.missionWonAnimation = createMissionWonAnimation(
			playerInfos,
			gameCanvas,
			c.engine.Kills,
			c.model.mission,
			window,
		)
	}

}

func (c *controller) onMissionWon() {
	engine.MissionController.MarkMissionFinished(c.model.mission)
	var remainingMissions, _ = engine.MissionController.RemainingMissions()
	if len(remainingMissions) == 0 {
		c.moveToHighScoreScreen()
	} else {
		c.stereo.PlaySong(assets.ThemeSong)
		c.screenChangeRequired(common.Mission)
	}
}

func (c *controller) moveToHighScoreScreen() {
	var p1, p2 = highscore.EntryControllerInstance.ReachedHighScore()
	if p1 {
		c.screenChangeRequired(common.EnterNameP1)
	} else if p2 {
		c.screenChangeRequired(common.EnterNameP2)
	} else {
		c.screenChangeRequired(common.HighScore)
	}
}
