// Package mission contains the screen that displays the world map to the user. The user can then use the controller to
// select his next mission.
package mission

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"retro-carnage/assets"
	"retro-carnage/engine"
	"retro-carnage/engine/geometry"
	"retro-carnage/engine/input"
	"retro-carnage/logging"
	"retro-carnage/ui/common"
	"retro-carnage/util"
)

const worldMapImagePath = "./images/backgrounds/world-map.jpg"

type Screen struct {
	availableMissions    []*assets.Mission
	inputController      input.Controller
	missionsInitialized  bool
	screenChangeRequired common.ScreenChangeCallback
	selectedMission      *assets.Mission
	textDimensions       map[string]*geometry.Point
	window               *pixelgl.Window
	worldMapSprite       *pixel.Sprite
}

func (s *Screen) SetInputController(inputCtrl input.Controller) {
	s.inputController = inputCtrl
}

func (s *Screen) SetScreenChangeCallback(callback common.ScreenChangeCallback) {
	s.screenChangeRequired = callback
}

func (s *Screen) SetWindow(window *pixelgl.Window) {
	s.window = window
}

func (s *Screen) SetUp() {
	if !assets.MissionRepository.Initialized() {
		assets.MissionRepository.Initialize()
	} else {
		s.initializeMissions()
	}
	s.worldMapSprite = common.LoadSprite(worldMapImagePath)
}

func (s *Screen) initializeMissions() {
	remainingMissions, err := engine.MissionController.RemainingMissions()
	if nil != err {
		logging.Error.Fatalf("Failed to retrieve list of remaining missions: %v", err)
	}
	if 0 == len(remainingMissions) {
		logging.Error.Fatalf("List of remaining missions is empty. Game should have ended!")
	}

	s.availableMissions = remainingMissions
	s.selectedMission = remainingMissions[0]
}

func (s *Screen) Update(_ int64) {
	if !s.missionsInitialized && assets.MissionRepository.Initialized() {
		s.missionsInitialized = true
		s.initializeMissions()
	}

	s.drawWorldMap()

	if s.missionsInitialized {
		// TODO: Get input
		// TODO: Update position (if necessary)
		// TODO: Draw mission locations
	}

	// TODO: Draw image of client
	// TODO: Draw mission briefing
}

func (s *Screen) TearDown() {}

func (s *Screen) String() string {
	return string(common.Mission)
}

func (s *Screen) drawWorldMap() {
	var factorX = (s.window.Bounds().Max.X - 100) / s.worldMapSprite.Picture().Bounds().Max.X
	var factorY = (s.window.Bounds().Max.X * 3 / 4) / s.worldMapSprite.Picture().Bounds().Max.X
	var factor = util.Min(factorX, factorY)

	var mapCenter = s.window.Bounds().Center()
	mapCenter.Y -= s.window.Bounds().Max.Y / 8
	s.worldMapSprite.Draw(s.window, pixel.IM.Scaled(pixel.Vec{X: 0, Y: 0}, factor).Moved(mapCenter))
}
