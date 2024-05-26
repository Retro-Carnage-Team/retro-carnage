package config

import "github.com/Retro-Carnage-Team/pixel/pixelgl"

type VideoConfiguration struct {
	UsePrimaryMonitor bool   `json:"usePrimaryMonitor"`
	SelectedMonitor   string `json:"selectedMonitor"`
	FullScreen        bool   `json:"fullScreen"`
	Width             int    `json:"width"`
	Height            int    `json:"height"`
}

func newDefaultVideoConfiguration() VideoConfiguration {
	return VideoConfiguration{
		UsePrimaryMonitor: true,
		FullScreen:        true,
	}
}

func (vc VideoConfiguration) GetConfiguredMonitor() *pixelgl.Monitor {
	if !vc.UsePrimaryMonitor {
		for _, m := range pixelgl.Monitors() {
			if m.Name() == vc.SelectedMonitor {
				return m
			}
		}
	}
	return pixelgl.PrimaryMonitor()
}
