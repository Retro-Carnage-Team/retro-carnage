package input

import (
	"fmt"
	"time"
)

type DeviceState struct {
	MoveUp     bool
	MoveDown   bool
	MoveLeft   bool
	MoveRight  bool
	Fire       bool
	Grenade    bool
	ToggleUp   bool
	ToggleDown bool
}

func (ds *DeviceState) String() string {
	return fmt.Sprintf("DeviceState[Fire: %t, Grenade: %t, Toggle ↑: %t, Toggle ↓: %t, ↑: %t, →: %t, ↓: %t, ←: %t]",
		ds.Fire, ds.Grenade, ds.ToggleUp, ds.ToggleDown, ds.MoveUp, ds.MoveRight, ds.MoveDown, ds.MoveLeft)
}

func (ds *DeviceState) IsButtonPressed() bool {
	return ds.Fire || ds.Grenade || ds.ToggleUp || ds.ToggleDown
}

func (ds *DeviceState) Combine(other *DeviceState) *DeviceState {
	return &DeviceState{
		MoveUp:     ds.MoveUp || other.MoveUp,
		MoveDown:   ds.MoveDown || other.MoveDown,
		MoveLeft:   ds.MoveLeft || other.MoveLeft,
		MoveRight:  ds.MoveRight || other.MoveRight,
		Fire:       ds.Fire || other.Fire,
		Grenade:    ds.Grenade || other.Grenade,
		ToggleUp:   ds.ToggleUp || other.ToggleUp,
		ToggleDown: ds.ToggleDown || other.ToggleDown,
	}
}

type UiEventState struct {
	MovedUp       bool
	MovedDown     bool
	MovedLeft     bool
	MovedRight    bool
	PressedButton bool
}

func (es *UiEventState) String() string {
	return fmt.Sprintf("UiEventState[PressedButton: %t, ↑: %t, →: %t, ↓: %t, ←: %t]",
		es.PressedButton, es.MovedUp, es.MovedRight, es.MovedDown, es.MovedLeft)
}

type rapidFireState struct {
	pressedSince     *time.Time
	reachedThreshold bool
}

func (rfs *rapidFireState) update(inputState *DeviceState) bool {
	if inputState.IsButtonPressed() {
		if nil == rfs.pressedSince {
			var now = time.Now()
			rfs.pressedSince = &now
			return true
		} else {
			if !rfs.reachedThreshold && time.Since(*rfs.pressedSince).Milliseconds() > rapidFireThreshold {
				rfs.reachedThreshold = true
				var now = time.Now()
				rfs.pressedSince = &now
				return true
			} else if rfs.reachedThreshold && time.Since(*rfs.pressedSince).Milliseconds() > rapidFireOffset {
				var now = time.Now()
				rfs.pressedSince = &now
				return true
			}
			return false
		}
	} else {
		rfs.pressedSince = nil
		return false
	}
}
