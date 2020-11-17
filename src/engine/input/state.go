package input

import "fmt"

type State struct {
	MoveUp     bool
	MoveDown   bool
	MoveLeft   bool
	MoveRight  bool
	Fire       bool
	Grenade    bool
	ToggleUp   bool
	ToggleDown bool
}

func (is *State) String() string {
	return fmt.Sprintf("State[Fire: %t, Grenade: %t, Toggle ↑: %t, Toggle ↓: %t, ↑: %t, →: %t, ↓: %t, ←: %t]",
		is.Fire, is.Grenade, is.ToggleUp, is.ToggleDown, is.MoveUp, is.MoveRight, is.MoveDown, is.MoveLeft)
}
