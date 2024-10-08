package mission

import (
	"retro-carnage/assets"
	"retro-carnage/engine"
	"retro-carnage/input"
	"retro-carnage/logging"
	"retro-carnage/ui/common"
)

type controller struct {
	availableMissions    []*assets.Mission
	inputController      input.InputController
	missionsInitialized  bool
	screenChangeRequired common.ScreenChangeCallback
	selectedMission      *assets.Mission
}

func newController() *controller {
	var result = controller{}
	return &result
}

func (c *controller) setInputController(inputCtrl input.InputController) {
	c.inputController = inputCtrl
}

func (c *controller) setScreenChangeCallback(callback common.ScreenChangeCallback) {
	c.screenChangeRequired = callback
}

func (c *controller) initialize() {
	if !assets.MissionRepository.Initialized() {
		assets.MissionRepository.Initialize()
	}
}

func (c *controller) initializeMissions() {
	remainingMissions, err := engine.MissionController.RemainingMissions()
	if nil != err {
		logging.Error.Fatalf("Failed to retrieve list of remaining missions: %v", err)
	}
	if len(remainingMissions) == 0 {
		logging.Error.Fatalf("List of remaining missions is empty. Game should have ended!")
	}

	c.availableMissions = remainingMissions
	c.selectedMission = remainingMissions[0]
}

func (c *controller) update() {
	if !c.missionsInitialized && assets.MissionRepository.Initialized() {
		c.missionsInitialized = true
		c.initializeMissions()
	}

	if c.missionsInitialized {
		c.processUserInput()
	}
}

func (c *controller) processUserInput() {
	var uiEventState = c.inputController.GetUiEventStateCombined()
	if nil == uiEventState {
		return
	}

	if uiEventState.PressedButton {
		engine.MissionController.SelectMission(c.selectedMission)
		c.screenChangeRequired(common.BuyYourWeaponsP1)
		go assets.NewStereo().BufferSong(c.selectedMission.Music)
	} else {
		var nextMission = c.selectedMission
		var err error = nil
		if uiEventState.MovedUp {
			nextMission, err = engine.MissionController.NextMissionNorth(&c.selectedMission.Location)
		} else if uiEventState.MovedDown {
			nextMission, err = engine.MissionController.NextMissionSouth(&c.selectedMission.Location)
		} else if uiEventState.MovedLeft {
			nextMission, err = engine.MissionController.NextMissionWest(&c.selectedMission.Location)
		} else if uiEventState.MovedRight {
			nextMission, err = engine.MissionController.NextMissionEast(&c.selectedMission.Location)
		}
		if nil != err {
			logging.Error.Fatalf("Failed to get next mission: %v", err)
		}
		if nil != nextMission {
			c.selectedMission = nextMission
		}
	}
}
