package main

import (
	"fmt"
	"os"
	"retro-carnage/assets"
	"retro-carnage/logging"
	"retro-carnage/ui"
	"retro-carnage/ui/common/fonts"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
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
	if len(os.Args) != 1 {
		fmt.Println(os.Args[1])
		err := os.Chdir(os.Args[1])
		if err != nil {
			panic(err)
		}

		if len(os.Args) > 2 {
			for i := 2; i < len(os.Args); i++ {
				if os.Args[i] == "debug" {
					logging.ActivateTraceLogger()
					logging.Trace.Println("Activated trace logger")
				}
			}
		}
	}

	pixelgl.Run(run)
}
