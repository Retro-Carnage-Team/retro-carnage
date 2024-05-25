package game

import (
	"fmt"
	"time"

	"github.com/Retro-Carnage-Team/pixel/pixelgl"
)

// fpsInfo is a utility to measure the Frames Per Second (FPS) of the graphical output. This info can then be drawn to
// screen. We sync the graphical output to the speed of the monitor. That's why FPS will not be > the frequency of your
// hardware.
type fpsInfo struct {
	fps    int
	frames int
	second <-chan time.Time
}

func (fi *fpsInfo) update() {
	fi.frames++
	select {
	case <-fi.second:
		fi.fps = fi.frames
		fi.frames = 0
	default:
	}
}

func (fi *fpsInfo) drawToScreen(window *pixelgl.Window) {
	window.SetTitle(fmt.Sprintf("RETRO CARNAGE (%d FPS)", fi.fps))
}
