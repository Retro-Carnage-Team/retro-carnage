package game

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"retro-carnage/assets"
	"retro-carnage/engine"
	"retro-carnage/engine/geometry"
	"retro-carnage/engine/input"
	"retro-carnage/logging"
	"retro-carnage/ui/common"
)

type Screen struct {
	engine               *engine.GameEngine
	inputController      input.Controller
	mission              *assets.Mission
	playerInfoAreas      []*geometry.Rectangle
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
	var player0InfoArea = &geometry.Rectangle{
		X:      0,
		Y:      0,
		Width:  (s.window.Bounds().W() - s.window.Bounds().H()) / 2,
		Height: s.window.Bounds().H(),
	}
	var player1InfoArea = &geometry.Rectangle{
		X:      s.window.Bounds().W() - player0InfoArea.Width,
		Y:      0,
		Width:  player0InfoArea.Width,
		Height: player0InfoArea.Height,
	}

	s.playerInfoAreas = []*geometry.Rectangle{player0InfoArea, player1InfoArea}
	s.stereo = assets.NewStereo()
	s.mission = engine.MissionController.CurrentMission()
	if nil != s.mission {
		s.stereo.PlaySong(s.mission.Music)
		s.engine = engine.NewGameEngine(s.mission)
		s.engine.SetInputController(s.inputController)
		s.renderer = engine.NewRenderer(s.engine, s.window)
	}
}

func (s *Screen) Update(elapsedTimeInMs int64) {
	s.drawPlayerInfo(0, s.playerInfoAreas[0])
	s.drawPlayerInfo(1, s.playerInfoAreas[1])
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

func (s *Screen) drawPlayerInfo(playerIdx int, rect *geometry.Rectangle) {
	imd := imdraw.New(nil)
	imd.Color = common.Red

	imd.Push(pixel.V(rect.X, rect.Y))
	imd.Push(pixel.V(rect.X+rect.Width, rect.Y))
	imd.Push(pixel.V(rect.X, rect.Y+rect.Height))
	imd.Push(pixel.V(rect.X+rect.Width, rect.Y+rect.Height))
	imd.Rectangle(1)
	imd.Draw(s.window)
}
