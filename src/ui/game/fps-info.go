package game

import (
	"fmt"
	"github.com/faiface/pixel/pixelgl"
	"time"
)

// FpsInfo is a utility to measure the Frames Per Second (FPS) of the graphical output. This info can then be drawn to
// screen. We sync the graphical output to the speed of the monitor. That's why FPS will not be > the frequency of your
// hardware.
type FpsInfo struct {
	fps    int
	frames int
	second <-chan time.Time
}

func (fi *FpsInfo) update() {
	fi.frames++
	select {
	case <-fi.second:
		fi.fps = fi.frames
		fi.frames = 0
	default:
	}
}

func (fi *FpsInfo) drawToScreen(window *pixelgl.Window) {
	window.SetTitle(fmt.Sprintf("RETRO CARNAGE (%d FPS)", fi.fps))
}
