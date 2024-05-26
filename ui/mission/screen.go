// Package mission contains the screen that displays the world map to the user. The user can then use the controller to
// select his next mission.
package mission

import (
	"math"
	"retro-carnage/assets"
	"retro-carnage/engine"
	"retro-carnage/engine/geometry"
	"retro-carnage/input"
	"retro-carnage/logging"
	"retro-carnage/ui/common"
	"retro-carnage/ui/common/fonts"

	pixel "github.com/Retro-Carnage-Team/pixel2"
	"github.com/Retro-Carnage-Team/pixel2/backends/opengl"
	"github.com/Retro-Carnage-Team/pixel2/ext/imdraw"
)

const (
	crossHairImagePath   = "images/other/crosshair.png"
	locationMarkerRadius = 10
	worldMapImagePath    = "images/other/world-map.jpg"
	worldMapWidth        = 1280
	worldMapHeight       = 783
)

var (
	locationMarkerColor = common.ParseHexColor("#fea400")
)

type Screen struct {
	availableMissions         []*assets.Mission
	briefingFontSize          int
	crossHairSprite           *pixel.Sprite
	inputController           input.InputController
	missionsInitialized       bool
	missionNameToClientSprite map[string]*pixel.Sprite
	screenChangeRequired      common.ScreenChangeCallback
	selectedMission           *assets.Mission
	window                    *opengl.Window
	worldMapSprite            *pixel.Sprite
}

func (s *Screen) SetInputController(inputCtrl input.InputController) {
	s.inputController = inputCtrl
}

func (s *Screen) SetScreenChangeCallback(callback common.ScreenChangeCallback) {
	s.screenChangeRequired = callback
}

func (s *Screen) SetWindow(window *opengl.Window) {
	s.window = window
}

func (s *Screen) SetUp() {
	if !assets.MissionRepository.Initialized() {
		assets.MissionRepository.Initialize()
	} else {
		s.initializeMissions()
	}
	s.briefingFontSize = fonts.DefaultFontSize() - 2
	s.crossHairSprite = assets.SpriteRepository.Get(crossHairImagePath)
	s.worldMapSprite = assets.SpriteRepository.Get(worldMapImagePath)
}

func (s *Screen) initializeMissions() {
	remainingMissions, err := engine.MissionController.RemainingMissions()
	if nil != err {
		logging.Error.Fatalf("Failed to retrieve list of remaining missions: %v", err)
	}
	if len(remainingMissions) == 0 {
		logging.Error.Fatalf("List of remaining missions is empty. Game should have ended!")
	}

	s.missionNameToClientSprite = make(map[string]*pixel.Sprite)
	for _, mission := range remainingMissions {
		s.missionNameToClientSprite[mission.Name] = assets.SpriteRepository.Get(mission.Client)
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
		s.processUserInput()
		s.drawMissionLocations()
		s.drawClientPicture()
		s.drawMissionDescription()
	}
}

func (s *Screen) TearDown() {
	// No tear down action required
}

func (s *Screen) String() string {
	return string(common.Mission)
}

func (s *Screen) drawWorldMap() {
	var factor = s.getWorldMapScalingFactor()
	var mapCenter = s.getWorldMapCenter()
	s.worldMapSprite.Draw(s.window, pixel.IM.Scaled(pixel.Vec{X: 0, Y: 0}, factor).Moved(mapCenter))
}

func (s *Screen) drawMissionLocations() {
	var factor = s.getWorldMapScalingFactor()
	var mapCenter = s.getWorldMapCenter()
	var mapBottomLeft = pixel.Vec{
		X: mapCenter.X - (worldMapWidth*factor)/2,
		Y: mapCenter.Y - (worldMapHeight*factor)/2,
	}

	for _, city := range s.availableMissions {
		var cityLocation = pixel.Vec{
			X: mapBottomLeft.X + (city.Location.Longitude * factor),
			Y: mapBottomLeft.Y + (worldMapHeight * factor) - (city.Location.Latitude * factor),
		}
		s.drawLocationMarker(cityLocation)
		if city.Name == s.selectedMission.Name {
			s.crossHairSprite.Draw(s.window, pixel.IM.Moved(cityLocation))
		}
	}
}

func (s *Screen) drawClientPicture() {
	var clientSprite = s.missionNameToClientSprite[s.selectedMission.Name]
	var scalingFactor = (s.window.Bounds().Max.Y * 3 / 16) / clientSprite.Picture().Bounds().Max.Y
	var margin = s.window.Bounds().Max.Y * 1 / 32
	var positionX = margin + (clientSprite.Picture().Bounds().Max.X * scalingFactor / 2)
	var positionY = s.window.Bounds().Max.Y - margin - (clientSprite.Picture().Bounds().Max.Y * scalingFactor / 2)

	clientSprite.Draw(s.window, pixel.IM.
		Scaled(pixel.Vec{X: 0, Y: 0}, scalingFactor).
		Moved(pixel.V(positionX, positionY)))
}

func (s *Screen) drawMissionDescription() {
	var clientSprite = s.missionNameToClientSprite[s.selectedMission.Name]
	var scalingFactor = (s.window.Bounds().Max.Y * 3 / 16) / clientSprite.Picture().Bounds().Max.Y
	var margin = s.window.Bounds().Max.Y * 1 / 32

	var positionX = margin + (clientSprite.Picture().Bounds().Max.X * scalingFactor) + margin
	var positionY = s.window.Bounds().Max.Y - margin - (clientSprite.Picture().Bounds().Max.Y * scalingFactor)
	var textAreaWidth = s.window.Bounds().Max.X - positionX - margin
	var textAreaHeight = s.window.Bounds().Max.Y/4 - margin - margin
	var text = s.selectedMission.Briefing

	var renderer = fonts.TextRenderer{Window: s.window}
	textLayout, err := renderer.CalculateTextLayout(text, s.briefingFontSize, int(textAreaWidth), int(textAreaHeight))
	if nil == err {
		renderer.RenderTextLayout(textLayout, s.briefingFontSize, common.White, &geometry.Point{
			X: positionX,
			Y: positionY,
		})
	} else {
		logging.Warning.Fatalf("Failed to render text: %v", err)
	}
}

func (s *Screen) getWorldMapScalingFactor() float64 {
	var factorX = (s.window.Bounds().Max.X - 100) / s.worldMapSprite.Picture().Bounds().Max.X
	var factorY = (s.window.Bounds().Max.Y * 3 / 4) / s.worldMapSprite.Picture().Bounds().Max.Y
	return math.Min(factorX, factorY)
}

func (s *Screen) getWorldMapCenter() (result pixel.Vec) {
	result = s.window.Bounds().Center()
	result.Y -= s.window.Bounds().Max.Y / 8
	return
}

func (s *Screen) drawLocationMarker(location pixel.Vec) {
	imd := imdraw.New(nil)
	imd.Color = locationMarkerColor
	imd.Push(location)
	imd.Circle(locationMarkerRadius, 5)
	imd.Draw(s.window)
}

func (s *Screen) processUserInput() {
	var uiEventState = s.inputController.GetUiEventStateCombined()
	if nil == uiEventState {
		return
	}

	if uiEventState.PressedButton {
		engine.MissionController.SelectMission(s.selectedMission)
		s.screenChangeRequired(common.BuyYourWeaponsP1)
		go assets.NewStereo().BufferSong(s.selectedMission.Music)
	} else {
		var nextMission = s.selectedMission
		var err error = nil
		if uiEventState.MovedUp {
			nextMission, err = engine.MissionController.NextMissionNorth(&s.selectedMission.Location)
		} else if uiEventState.MovedDown {
			nextMission, err = engine.MissionController.NextMissionSouth(&s.selectedMission.Location)
		} else if uiEventState.MovedLeft {
			nextMission, err = engine.MissionController.NextMissionWest(&s.selectedMission.Location)
		} else if uiEventState.MovedRight {
			nextMission, err = engine.MissionController.NextMissionEast(&s.selectedMission.Location)
		}
		if nil != err {
			logging.Error.Fatalf("Failed to get next mission: %v", err)
		}
		if nil != nextMission {
			s.selectedMission = nextMission
		}
	}
}
