package game

import (
	"github.com/faiface/pixel/pixelgl"
	"retro-carnage/assets"
	"retro-carnage/engine"
	"retro-carnage/engine/input"
	"retro-carnage/logging"
	"retro-carnage/ui/common"
	"retro-carnage/ui/highscore"
	"time"
)

// Screen in this package is the one that show the actual gameplay.
type Screen struct {
	engine               *engine.GameEngine
	fpsInfo              *fpsInfo
	gameLostAnimation    *gameLostAnimation
	gameWonAnimation     *gameWonAnimation
	inputController      input.Controller
	mission              *assets.Mission
	missionWonAnimation  *missionWonAnimation
	playerInfos          []*playerInfo
	renderer             *engine.Renderer
	screenChangeRequired common.ScreenChangeCallback
	stereo               *assets.Stereo
	window               *pixelgl.Window
}

// SetInputController is used to connect Screen with the global input.Controller instance.
func (s *Screen) SetInputController(ctrl input.Controller) {
	s.inputController = ctrl
}

// SetScreenChangeCallback is used to connect Screen with the callback method of ui.MainScreen.
func (s *Screen) SetScreenChangeCallback(callback common.ScreenChangeCallback) {
	s.screenChangeRequired = callback
}

// SetWindow is used to connect Screen with the pixelgl.Window instance.
func (s *Screen) SetWindow(window *pixelgl.Window) {
	s.window = window
}

// SetUp is called when the screen got initialized by ui.MainScreen and is about to appear shortly.
func (s *Screen) SetUp() {
	s.fpsInfo = &fpsInfo{second: time.Tick(time.Second)}
	s.playerInfos = []*playerInfo{
		newPlayerInfo(0, s.window),
		newPlayerInfo(1, s.window),
	}
	s.stereo = assets.NewStereo()
	s.mission = engine.MissionController.CurrentMission()
	if nil == s.mission {
		logging.Error.Fatalf("No missing selected on game screen")
	}

	s.stereo.PlaySong(s.mission.Music)
	s.engine = engine.NewGameEngine(s.mission)
	s.engine.SetInputController(s.inputController)
	s.renderer = engine.NewRenderer(s.engine, s.window)
}

// Update gets called for every frame that gets displayed.
// Here we update the state of the gameplay based on the time that has elapsed since the last frame.
// Then we render the new game state to the pixelgl.Window.
func (s *Screen) Update(elapsedTimeInMs int64) {
	if nil != s.engine && nil != s.renderer {
		if !(s.engine.Won || s.engine.Lost) {
			for _, playerInfo := range s.playerInfos {
				playerInfo.draw(s.window)
			}
			s.updateGameInProgress(elapsedTimeInMs)
			s.renderer.Render(elapsedTimeInMs)
		}

		if s.engine.Won {
			s.updateGameWon(elapsedTimeInMs)
		}

		if s.engine.Lost {
			s.updateGameLost(elapsedTimeInMs)
		}
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
	if nil != s.gameWonAnimation {
		s.gameWonAnimation.update(elapsedTimeInMs)
		s.gameWonAnimation.drawToScreen()
		if s.gameWonAnimation.finished || s.inputController.ControllerUiEventStateCombined().PressedButton {
			s.onMissionWon()
		}
	} else if nil != s.missionWonAnimation {
		for _, playerInfo := range s.playerInfos {
			playerInfo.draw(s.window)
		}
		s.renderer.Render(0)
		s.missionWonAnimation.update(elapsedTimeInMs)
		s.missionWonAnimation.drawToScreen()
		if s.missionWonAnimation.finished || s.inputController.ControllerUiEventStateCombined().PressedButton {
			var remainingMissions, _ = engine.MissionController.RemainingMissions()
			// The current mission has not been marked as won, yet. Thus there is one remaining mission.
			if (1 == len(remainingMissions)) && (remainingMissions[0].Name == s.mission.Name) {
				s.gameWonAnimation = createGameWonAnimation(s.mission, s.window)
			} else {
				s.onMissionWon()
			}
		}
	}
}

func (s *Screen) updateGameLost(elapsedTimeInMs int64) {
	for _, playerInfo := range s.playerInfos {
		playerInfo.draw(s.window)
	}
	s.renderer.Render(0)
	s.gameLostAnimation.update(elapsedTimeInMs)
	s.gameLostAnimation.drawToScreen()
	if s.gameLostAnimation.finished || s.inputController.ControllerUiEventStateCombined().PressedButton {
		s.moveToHighScoreScreen()
	}
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
	if 0 == len(remainingMissions) {
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
