package input

import "time"

type rapidFireState struct {
	pressedSince     *time.Time
	reachedThreshold bool
}

func (rfs *rapidFireState) update(inputState *InputDeviceState) bool {
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
