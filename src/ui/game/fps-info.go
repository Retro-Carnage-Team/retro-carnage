package game

import (
	"fmt"
	"github.com/faiface/pixel/pixelgl"
	"retro-carnage/engine/geometry"
	"retro-carnage/logging"
	"retro-carnage/ui/common"
	"retro-carnage/ui/common/fonts"
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
	var text = fmt.Sprintf("%d FPS", fi.fps)
	var renderer = fonts.TextRenderer{Window: window}
	textLayout, err := renderer.CalculateTextLayout(text, 20, int(window.Bounds().W()), 30)
	if nil == err {
		var positionX = window.Bounds().W() - textLayout.Lines()[0].Dimension().X - 10
		var positionY = window.Bounds().H() - textLayout.Height() - 15
		renderer.RenderTextLayout(textLayout, 20, common.Black, &geometry.Point{
			X: positionX,
			Y: positionY,
		})
	} else {
		logging.Warning.Fatalf("Failed to render fps: %v", err)
	}
}
