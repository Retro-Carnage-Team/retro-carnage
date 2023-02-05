package engine

import (
	"errors"
	"retro-carnage/assets"
	"retro-carnage/engine/characters"
	"sort"
)

type MissionCtrl struct {
	currentMission       *assets.Mission
	finishedMissionNames []string
}

var (
	MissionController = &MissionCtrl{finishedMissionNames: make([]string, 0)}
)

func (mc *MissionCtrl) Reset() {
	mc.currentMission = nil
	mc.finishedMissionNames = make([]string, 0)
}

func (mc *MissionCtrl) RemainingMissions() ([]*assets.Mission, error) {
	if !assets.MissionRepository.Initialized() {
		return nil, errors.New("mission repository has not been initialized, yet")
	}

	result := make([]*assets.Mission, 0)
	for _, mission := range assets.MissionRepository.Missions {
		var finished = false
		for _, name := range mc.finishedMissionNames {
			finished = finished || name == mission.Name
		}
		if !finished {
			result = append(result, mission)
		}
	}
	return result, nil
}

func (mc *MissionCtrl) MarkMissionFinished(mission *assets.Mission) {
	mc.finishedMissionNames = append(mc.finishedMissionNames, mission.Name)
	for _, player := range characters.PlayerController.RemainingPlayers() {
		player.SetCash(player.Cash() + mission.Reward)
	}
}

func (mc *MissionCtrl) CurrentMission() *assets.Mission {
	return mc.currentMission
}

func (mc *MissionCtrl) SelectMission(mission *assets.Mission) {
	if nil != mission {
		mc.currentMission = mission
	} else {
		mc.currentMission = nil
	}
}

func (mc *MissionCtrl) NextMissionNorth(relativeTo *assets.Location) (*assets.Mission, error) {
	var filter = func(mission *assets.Mission) bool {
		return mission.Location.Latitude < relativeTo.Latitude
	}

	var lessBuilder = func(missions []*assets.Mission) func(int, int) bool {
		return func(i, j int) bool {
			return missions[i].Location.Latitude > missions[j].Location.Latitude
		}
	}

	return mc.filterAndSortRemainingMissions(filter, lessBuilder)
}

func (mc *MissionCtrl) NextMissionSouth(relativeTo *assets.Location) (*assets.Mission, error) {
	var filter = func(mission *assets.Mission) bool {
		return mission.Location.Latitude > relativeTo.Latitude
	}

	var lessBuilder = func(missions []*assets.Mission) func(int, int) bool {
		return func(i, j int) bool {
			return missions[i].Location.Latitude < missions[j].Location.Latitude
		}
	}

	return mc.filterAndSortRemainingMissions(filter, lessBuilder)
}

func (mc *MissionCtrl) NextMissionWest(relativeTo *assets.Location) (*assets.Mission, error) {
	var filter = func(mission *assets.Mission) bool {
		return mission.Location.Longitude < relativeTo.Longitude
	}

	var lessBuilder = func(missions []*assets.Mission) func(int, int) bool {
		return func(i, j int) bool {
			return missions[i].Location.Longitude > missions[j].Location.Longitude
		}
	}

	return mc.filterAndSortRemainingMissions(filter, lessBuilder)
}

func (mc *MissionCtrl) NextMissionEast(relativeTo *assets.Location) (*assets.Mission, error) {
	var filter = func(mission *assets.Mission) bool {
		return mission.Location.Longitude > relativeTo.Longitude
	}

	var lessBuilder = func(missions []*assets.Mission) func(int, int) bool {
		return func(i, j int) bool {
			return missions[i].Location.Longitude < missions[j].Location.Longitude
		}
	}

	return mc.filterAndSortRemainingMissions(filter, lessBuilder)
}

func (mc *MissionCtrl) filterAndSortRemainingMissions(test func(*assets.Mission) bool, lessBuilder func([]*assets.Mission) func(int, int) bool) (*assets.Mission, error) {
	var filteredMissions = make([]*assets.Mission, 0)
	remainingMissions, err := mc.RemainingMissions()
	if nil != err {
		return nil, err
	}

	if len(remainingMissions) == 0 {
		return nil, nil
	}

	for _, m := range remainingMissions {
		if test(m) {
			filteredMissions = append(filteredMissions, m)
		}
	}

	if len(filteredMissions) == 0 {
		return nil, nil
	}

	if 1 < len(filteredMissions) {
		sort.SliceStable(filteredMissions, lessBuilder(filteredMissions))
	}

	return filteredMissions[0], nil
}
