package game

import "retro-carnage/assets"

type model struct {
	inProgress bool
	lost       bool
	mission    *assets.Mission
	paused     bool
	won        bool
}
