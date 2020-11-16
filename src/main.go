package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"retro-carnage.net/ui"
	uiUtil "retro-carnage.net/ui/util"
)

func run() {
	var monitor = pixelgl.PrimaryMonitor()
	var pixelX, pixelY = monitor.Size()
	cfg := pixelgl.WindowConfig{
		Bounds:  pixel.R(0, 0, pixelX, pixelY),
		Monitor: monitor,
		Title:   "RETRO CARNAGE",
		VSync:   true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	win.SetSmooth(true)

	uiUtil.InitializeFonts()
	uiUtil.NewStereo()

	var mainScreen = ui.MainScreen{Monitor: monitor, Window: win}
	mainScreen.Initialize()
	mainScreen.RunMainLoop()
}

func main() {
	pixelgl.Run(run)
}
