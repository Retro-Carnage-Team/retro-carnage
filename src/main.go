package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"retro-carnage.net/engine/input"
	"retro-carnage.net/util"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	win.Clear(colornames.Black)

	inputCtrl := input.Controller{Window: win}
	inputCtrl.HasTwoOrMoreDevices()
	inputCtrl.AssignControllersToPlayers()

	for !win.Closed() {
		win.Update()
		if nil != inputCtrl.ControllerPlayerOne {
			util.Trace.Print(inputCtrl.ControllerPlayerOne.State().String())
		}
	}
}

func main() {
	pixelgl.Run(run)
}
