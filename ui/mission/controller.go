package mission

import (
	"retro-carnage/assets"
	"retro-carnage/engine"
	"retro-carnage/input"
	"retro-carnage/logging"
	"retro-carnage/ui/common"
)

type controller struct {
	inputController      input.InputController
	model                missionModel
	screenChangeRequired common.ScreenChangeCallback
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

	c.model.availableMissions = remainingMissions
	c.model.selectedMission = remainingMissions[0]
}

func (c *controller) update() {
	if !c.model.initialized && assets.MissionRepository.Initialized() {
		c.initializeMissions()
		c.model.initialized = true
	}

	if c.model.initialized {
		c.processUserInput()
	}
}

func (c *controller) processUserInput() {
	var uiEventState = c.inputController.GetUiEventStateCombined()
	if nil == uiEventState {
		return
	}

	if uiEventState.PressedButton {
		engine.MissionController.SelectMission(c.model.selectedMission)
		c.screenChangeRequired(common.BuyYourWeaponsP1)
		go assets.NewStereo().BufferSong(c.model.selectedMission.Music)
	} else {
		var nextMission = c.model.selectedMission
		var err error = nil
		if uiEventState.MovedUp {
			nextMission, err = engine.MissionController.NextMissionNorth(&c.model.selectedMission.Location)
		} else if uiEventState.MovedDown {
			nextMission, err = engine.MissionController.NextMissionSouth(&c.model.selectedMission.Location)
		} else if uiEventState.MovedLeft {
			nextMission, err = engine.MissionController.NextMissionWest(&c.model.selectedMission.Location)
		} else if uiEventState.MovedRight {
			nextMission, err = engine.MissionController.NextMissionEast(&c.model.selectedMission.Location)
		}
		if nil != err {
			logging.Error.Fatalf("Failed to get next mission: %v", err)
		}
		if nil != nextMission {
			c.model.selectedMission = nextMission
		}
	}
}
