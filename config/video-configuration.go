package config

import "github.com/Retro-Carnage-Team/pixel2/backends/opengl"

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

func (vc VideoConfiguration) GetConfiguredMonitor() *opengl.Monitor {
	if !vc.UsePrimaryMonitor {
		for _, m := range opengl.Monitors() {
			if m.Name() == vc.SelectedMonitor {
				return m
			}
		}
	}
	return opengl.PrimaryMonitor()
}
