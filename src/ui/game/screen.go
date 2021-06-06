package game

import (
	"github.com/faiface/pixel/pixelgl"
	"retro-carnage/assets"
	"retro-carnage/engine"
	"retro-carnage/engine/input"
	"retro-carnage/logging"
	"retro-carnage/ui/common"
)

type Screen struct {
	engine               *engine.GameEngine
	inputController      input.Controller
	mission              *assets.Mission
	renderer             *engine.Renderer
	screenChangeRequired common.ScreenChangeCallback
	stereo               *assets.Stereo
	window               *pixelgl.Window
}

func (s *Screen) SetInputController(ctrl input.Controller) {
	s.inputController = ctrl
}

func (s *Screen) SetScreenChangeCallback(callback common.ScreenChangeCallback) {
	s.screenChangeRequired = callback
}

func (s *Screen) SetWindow(window *pixelgl.Window) {
	s.window = window
}

func (s *Screen) SetUp() {
	s.stereo = assets.NewStereo()
	s.mission = engine.MissionController.CurrentMission()
	if nil != s.mission {
		s.stereo.PlaySong(s.mission.Music)
		s.engine = engine.NewGameEngine(s.mission)
		s.renderer = engine.NewRenderer(s.engine, s.window)
	}
}

func (s *Screen) Update(elapsedTimeInMs int64) {
	if nil != s.engine && nil != s.renderer {
		s.engine.UpdateGameState(elapsedTimeInMs)
		s.renderer.Render(elapsedTimeInMs)
		if s.engine.Lost {
			s.onGameLost()
		} else if s.engine.Won {
			s.onMissionWon()
		}
	}
}

func (s *Screen) TearDown() {
	if nil != s.mission {
		s.stereo.StopSong(s.mission.Music)
	}
}

func (s *Screen) onGameLost() {
	// TODO: show high score screen
	s.screenChangeRequired(common.Title)
}

func (s *Screen) onMissionWon() {
	// TODO: Show level end animation / score calculation
	var mission = engine.MissionController.CurrentMission()
	if nil != mission {
		engine.MissionController.MarkMissionFinished(mission)
		var remainingMissions, err = engine.MissionController.RemainingMissions()
		if nil != err {
			logging.Error.Fatalf("Error on game screen: Level has been won when none have been initialized")
		}
		if 0 == len(remainingMissions) {
			// TODO: show high score screen
		} else {
			s.screenChangeRequired(common.Mission)
		}
	}
}

func (s *Screen) String() string {
	return string(common.Game)
}
