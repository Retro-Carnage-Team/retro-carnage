package input

import (
	"fmt"
)

type InputDeviceState struct {
	MoveUp        bool
	MoveDown      bool
	MoveLeft      bool
	MoveRight     bool
	PrimaryAction bool
	ToggleUp      bool
	ToggleDown    bool
}

func (ds *InputDeviceState) String() string {
	return fmt.Sprintf("DeviceState[Primary: %t, Toggle ↑: %t, Toggle ↓: %t, ↑: %t, →: %t, ↓: %t, ←: %t]",
		ds.PrimaryAction, ds.ToggleUp, ds.ToggleDown, ds.MoveUp, ds.MoveRight, ds.MoveDown, ds.MoveLeft)
}

func (ds *InputDeviceState) IsButtonPressed() bool {
	return ds.PrimaryAction || ds.ToggleUp || ds.ToggleDown
}

func (ds *InputDeviceState) Combine(other *InputDeviceState) *InputDeviceState {
	return &InputDeviceState{
		MoveUp:        ds.MoveUp || other.MoveUp,
		MoveDown:      ds.MoveDown || other.MoveDown,
		MoveLeft:      ds.MoveLeft || other.MoveLeft,
		MoveRight:     ds.MoveRight || other.MoveRight,
		PrimaryAction: ds.PrimaryAction || other.PrimaryAction,
		ToggleUp:      ds.ToggleUp || other.ToggleUp,
		ToggleDown:    ds.ToggleDown || other.ToggleDown,
	}
}
