package input

import "fmt"

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
