package main

import (
	"fmt"
	"os"
	"retro-carnage/assets"
	"retro-carnage/config"
	"retro-carnage/logging"
	"retro-carnage/ui"
	"retro-carnage/ui/common/fonts"

	"github.com/Retro-Carnage-Team/pixel"
	"github.com/Retro-Carnage-Team/pixel/pixelgl"
)

func run() {
	var configuration = config.GetConfigService().LoadVideoConfiguration()
	var monitor = configuration.GetConfiguredMonitor()
	var monitorW, monitorH = monitor.Size()

	var cfg pixelgl.WindowConfig
	if configuration.FullScreen {
		cfg = pixelgl.WindowConfig{
			Bounds:  pixel.R(0, 0, monitorW, monitorH),
			Monitor: monitor,
			Title:   "RETRO CARNAGE",
			VSync:   true,
		}
	} else {
		var confW = float64(configuration.Width)
		var confH = float64(configuration.Height)
		var x = (monitorW - confW) / 2
		var y = (monitorH - confH) / 2
		cfg = pixelgl.WindowConfig{
			Bounds:   pixel.R(0, 0, confW, confH),
			Position: pixel.Vec{X: x, Y: y},
			Title:    "RETRO CARNAGE",
			VSync:    true,
		}
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		logging.Error.Fatalf("failed to create window: %s", err)
	}

	win.SetCursorVisible(false)
	win.SetSmooth(true)

	fonts.Initialize(monitorW)
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
