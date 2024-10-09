package mission

import "retro-carnage/assets"

type model struct {
	availableMissions []*assets.Mission
	initialized       bool
	selectedMission   *assets.Mission
}
