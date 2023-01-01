package main

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"os"
	"retro-carnage/assets"
	"retro-carnage/ui"
	"retro-carnage/ui/common/fonts"
)

func run() {
	var monitor = pixelgl.PrimaryMonitor()
	var pixelX, pixelY = monitor.Size()
	// Set variables to different values to adjust window size
	pixelX = 1280.0
	pixelY = 800.0
	cfg := pixelgl.WindowConfig{
		Bounds: pixel.R(0, 0, pixelX, pixelY),
		// commenting out this line will make the app run in a window instead of full screen
		//Monitor: monitor,
		Title: "RETRO CARNAGE",
		VSync: true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	win.SetCursorVisible(false)
	win.SetSmooth(true)

	fonts.Initialize(pixelX)
	assets.NewStereo()

	var mainScreen = ui.MainScreen{Monitor: monitor, Window: win}
	mainScreen.Initialize()
	mainScreen.RunMainLoop()
}

func main() {
	if len(os.Args) == 1 {

	} else {
		fmt.Println(os.Args[1])
		err := os.Chdir(os.Args[1])
		if err != nil {
			panic(err)
		}
	}

	pixelgl.Run(run)
}
