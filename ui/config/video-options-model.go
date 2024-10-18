package config

import "retro-carnage/config"

type videoOptionsModel struct {
	selectedMonitorIndex int
	selectedOption       int
	videoConfig          config.VideoConfiguration
}
