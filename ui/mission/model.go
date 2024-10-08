package mission

import "retro-carnage/assets"

type missionModel struct {
	availableMissions []*assets.Mission
	initialized       bool
	selectedMission   *assets.Mission
}
