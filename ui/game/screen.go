package game

import (
	"retro-carnage/assets"
	"retro-carnage/engine"
	"retro-carnage/input"
	"retro-carnage/logging"
	"retro-carnage/ui/common"
	"retro-carnage/ui/common/fonts"
	"retro-carnage/ui/highscore"
	"time"

	pixel "github.com/Retro-Carnage-Team/pixel2"
	"github.com/Retro-Carnage-Team/pixel2/backends/opengl"
)

const pausedMessage = "GAME PAUSED"

// Screen in this package is the one that show the actual gameplay.
type Screen struct {
	engine               *engine.GameEngine
	fpsInfo              *fpsInfo
	gameLostAnimation    *gameLostAnimation
	gameWonAnimation     *gameWonAnimation
	inputController      input.InputController
	mission              *assets.Mission
	missionWonAnimation  *missionWonAnimation
	paused               bool
	playerInfos          []*playerInfo
	renderer             *engine.Renderer
	screenChangeRequired common.ScreenChangeCallback
	stereo               *assets.Stereo
	window               *opengl.Window
}

// SetInputController is used to connect Screen with the global input.Controller instance.
func (s *Screen) SetInputController(ctrl input.InputController) {
	s.inputController = ctrl
}

// SetScreenChangeCallback is used to connect Screen with the callback method of ui.MainScreen.
func (s *Screen) SetScreenChangeCallback(callback common.ScreenChangeCallback) {
	s.screenChangeRequired = callback
}

// SetWindow is used to connect Screen with the opengl.Window instance.
func (s *Screen) SetWindow(window *opengl.Window) {
	s.window = window
}

// SetUp is called when the screen got initialized by ui.MainScreen and is about to appear shortly.
func (s *Screen) SetUp() {
	s.paused = false
	s.fpsInfo = &fpsInfo{second: time.NewTicker(time.Second).C}
	s.playerInfos = []*playerInfo{
		newPlayerInfo(0, s.window),
		newPlayerInfo(1, s.window),
	}
	s.stereo = assets.NewStereo()
	s.mission = engine.MissionController.CurrentMission()
	if nil == s.mission {
		logging.Error.Fatalf("No mission selected on game screen")
	}

	s.stereo.PlaySong(s.mission.Music)
	s.engine = engine.NewGameEngine(s.mission)
	s.engine.SetInputController(s.inputController)
	s.renderer = engine.NewRenderer(s.engine, s.window)
}

// Update gets called for every frame that gets displayed.
// Here we update the state of the gameplay based on the time that has elapsed since the last frame.
// Then we render the new game state to the opengl.Window.
func (s *Screen) Update(elapsedTimeInMs int64) {
	if nil == s.engine || nil == s.renderer {
		return
	}

	if !(s.engine.Won || s.engine.Lost) {
		for _, playerInfo := range s.playerInfos {
			playerInfo.draw(s.window)
		}
		if s.window.JustReleased(pixel.KeyPause) {
			s.paused = !s.paused
		}
		if s.paused {
			s.renderGamePausedScreen()
		} else {
			s.updateGameInProgress(elapsedTimeInMs)
			s.renderer.Render(elapsedTimeInMs)
		}
	}

	if s.engine.Won {
		s.updateGameWon(elapsedTimeInMs)
	}

	if s.engine.Lost {
		s.updateGameLost(elapsedTimeInMs)
	}

	s.fpsInfo.update()
	s.fpsInfo.drawToScreen(s.window)
}

func (s *Screen) updateGameInProgress(elapsedTimeInMs int64) {
	s.engine.UpdateGameState(elapsedTimeInMs)
	var gameCanvas = s.renderer.Render(elapsedTimeInMs)
	if s.engine.Lost {
		s.gameLostAnimation = createGameLostAnimation(s.playerInfos, gameCanvas, s.mission, s.window)
	} else if s.engine.Won {
		s.missionWonAnimation = createMissionWonAnimation(
			s.playerInfos,
			gameCanvas,
			s.engine.Kills,
			s.mission,
			s.window,
		)
	}
}

func (s *Screen) updateGameWon(elapsedTimeInMs int64) {
	var uiEventStateCombined = s.inputController.GetUiEventStateCombined()
	var buttonPressed = (nil != uiEventStateCombined) && uiEventStateCombined.PressedButton
	if nil != s.gameWonAnimation {
		s.updateGameWonAnimation(elapsedTimeInMs, buttonPressed)
	} else if nil != s.missionWonAnimation {
		s.updateMissionWonAnimation(elapsedTimeInMs, buttonPressed)
	}
}

func (s *Screen) updateMissionWonAnimation(elapsedTimeInMs int64, buttonPressed bool) {
	for _, playerInfo := range s.playerInfos {
		playerInfo.draw(s.window)
	}
	s.renderer.Render(0)
	s.missionWonAnimation.update(elapsedTimeInMs)
	s.missionWonAnimation.drawToScreen()
	if s.missionWonAnimation.finished || buttonPressed {
		var remainingMissions, _ = engine.MissionController.RemainingMissions()
		// The current mission has not been marked as won, yet. Thus, there is one remaining mission.
		if (len(remainingMissions) == 1) && (remainingMissions[0].Name == s.mission.Name) {
			s.gameWonAnimation = createGameWonAnimation(s.mission, s.window)
		} else {
			s.onMissionWon()
		}
	}
}

func (s *Screen) updateGameWonAnimation(elapsedTimeInMs int64, buttonPressed bool) {
	s.gameWonAnimation.update(elapsedTimeInMs)
	s.gameWonAnimation.drawToScreen()
	if s.gameWonAnimation.finished || buttonPressed {
		s.onMissionWon()
	}
}

func (s *Screen) updateGameLost(elapsedTimeInMs int64) {
	for _, playerInfo := range s.playerInfos {
		playerInfo.draw(s.window)
	}
	s.renderer.Render(0)
	s.gameLostAnimation.update(elapsedTimeInMs)
	s.gameLostAnimation.drawToScreen()
	var uiEventStateCombined = s.inputController.GetUiEventStateCombined()
	var buttonPressed = nil != uiEventStateCombined && uiEventStateCombined.PressedButton
	if s.gameLostAnimation.finished || buttonPressed {
		s.moveToHighScoreScreen()
	}
}

func (s *Screen) renderGamePausedScreen() {
	var gameCanvas = s.renderer.Render(0)
	var matrix = pixel.IM.Moved(s.window.Bounds().Center())
	gameCanvas.DrawColorMask(s.window, matrix, pixel.RGBA{A: 0.5})

	var renderer = fonts.TextRenderer{Window: s.window}
	renderer.DrawLineToScreenCenter(pausedMessage, 0, common.White)
}

// TearDown is called by ui.MainWindow when the Screen has been displayed for the last time.
// Here we clean up and free used resources.
func (s *Screen) TearDown() {
	s.stereo.StopSong(s.mission.Music)
	for _, playerInfo := range s.playerInfos {
		playerInfo.dispose()
	}
	s.window.SetTitle("RETRO CARNAGE")
}

func (s *Screen) onMissionWon() {
	engine.MissionController.MarkMissionFinished(s.mission)
	var remainingMissions, _ = engine.MissionController.RemainingMissions()
	if len(remainingMissions) == 0 {
		s.moveToHighScoreScreen()
	} else {
		s.stereo.PlaySong(assets.ThemeSong)
		s.screenChangeRequired(common.Mission)
	}
}

func (s *Screen) moveToHighScoreScreen() {
	var p1, p2 = highscore.EntryControllerInstance.ReachedHighScore()
	if p1 {
		s.screenChangeRequired(common.EnterNameP1)
	} else if p2 {
		s.screenChangeRequired(common.EnterNameP2)
	} else {
		s.screenChangeRequired(common.HighScore)
	}
}

// String returns the screen name as string
func (s *Screen) String() string {
	return string(common.Game)
}
