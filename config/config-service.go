package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"retro-carnage/logging"
)

var (
	configService *ConfigService
)

type ConfigService struct{}

func GetConfigService() *ConfigService {
	if configService == nil {
		configService = &ConfigService{}
	}
	return configService
}

// LoadAudioConfigurations reads the audio configuration that is stored on disk.
// Returns configuration with default values if no data can be found.
func (cs *ConfigService) LoadAudioConfiguration() AudioConfiguration {
	filePath, err := cs.buildAudioConfigurationFilePath()
	if nil != err {
		logging.Warning.Printf("failed to compute audio configuration file path")
		return newDefaultAudioConfiguration()
	}

	logging.Trace.Printf("loading audio configuration from %s", filePath)
	data, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			logging.Info.Printf("audio config file not present %s", filePath)
		} else {
			logging.Warning.Printf("failed to read audio config %s", filePath)
		}
		return newDefaultAudioConfiguration()
	}

	var config = &AudioConfiguration{}
	err = json.Unmarshal(data, config)
	if err != nil {
		logging.Warning.Printf("failed to deserialize audio config %s", filePath)
	}
	return *config
}

// SaveAudioConfiguration stores the given audio configuration
func (cs *ConfigService) SaveAudioConfiguration(ac AudioConfiguration) error {
	var err = cs.initializeConfigurationFolder()
	if nil != err {
		return err
	}

	filePath, err := cs.buildAudioConfigurationFilePath()
	if nil != err {
		logging.Warning.Println("failed to calculate audio config path")
		return err
	}

	logging.Trace.Printf("saving audio configuration to %s", filePath)
	jsonData, _ := json.Marshal(ac)
	err = os.WriteFile(filePath, jsonData, 0600)
	if err != nil {
		logging.Warning.Printf("failed to write audio config %s", filePath)
		return err
	}

	return nil
}

// LoadInputDeviceConfigurations reads the controller configurations that are stored on disk.
// Returns empty array if no configurations can be found.
func (cs *ConfigService) LoadInputDeviceConfigurations() []InputDeviceConfiguration {
	var result = make([]InputDeviceConfiguration, 0)
	for i := 0; i < 2; i++ {
		filePath, err := cs.buildInputDeviceConfigurationFilePath(i)
		if nil != err {
			continue
		}

		logging.Trace.Printf("loading controller configuration for player %d from %s", i, filePath)
		data, err := os.ReadFile(filePath)
		if err != nil {
			if os.IsNotExist(err) {
				logging.Info.Printf("controller config file not present %s", filePath)
			} else {
				logging.Warning.Printf("failed to read controller config %s", filePath)
			}
			continue
		}

		var config = &InputDeviceConfiguration{}
		err = json.Unmarshal(data, config)
		if err != nil {
			logging.Warning.Printf("failed to deserialize controller config %s", filePath)
		}
		result = append(result, *config)
	}

	return result
}

// SaveInputDeviceConfiguration stores the given controller configuration for the specified player
func (cs *ConfigService) SaveInputDeviceConfiguration(cc InputDeviceConfiguration, playerIdx int) error {
	var err = cs.initializeConfigurationFolder()
	if nil != err {
		return err
	}

	filePath, err := cs.buildInputDeviceConfigurationFilePath(playerIdx)
	if nil != err {
		logging.Warning.Printf("failed to calculate config path for input device %d", playerIdx)
		return err
	}

	logging.Trace.Printf("saving input device configuration for player %d to %s", playerIdx, filePath)
	jsonData, _ := json.Marshal(cc)
	err = os.WriteFile(filePath, jsonData, 0600)
	if err != nil {
		logging.Warning.Printf("failed to write input device config %s", filePath)
		return err
	}

	return nil
}

// LoadVideoConfigurations reads the video configuration that is stored on disk.
// Returns configuration with default values if no data can be found.
func (cs *ConfigService) LoadVideoConfiguration() VideoConfiguration {
	filePath, err := cs.buildVideoConfigurationFilePath()
	if nil != err {
		logging.Warning.Printf("failed to compute video configuration file path")
		return newDefaultVideoConfiguration()
	}

	logging.Trace.Printf("loading video configuration from %s", filePath)
	data, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			logging.Info.Printf("video config file not present %s", filePath)
		} else {
			logging.Warning.Printf("failed to read video config %s", filePath)
		}
		return newDefaultVideoConfiguration()
	}

	var config = &VideoConfiguration{}
	err = json.Unmarshal(data, config)
	if err != nil {
		logging.Warning.Printf("failed to deserialize video config %s", filePath)
	}
	return *config
}

// SaveVideoConfiguration stores the given video configuration
func (cs *ConfigService) SaveVideoConfiguration(vc VideoConfiguration) error {
	var err = cs.initializeConfigurationFolder()
	if nil != err {
		return err
	}

	filePath, err := cs.buildVideoConfigurationFilePath()
	if nil != err {
		logging.Warning.Println("failed to calculate video config path")
		return err
	}

	logging.Trace.Printf("saving video configuration to %s", filePath)
	jsonData, _ := json.Marshal(vc)
	err = os.WriteFile(filePath, jsonData, 0600)
	if err != nil {
		logging.Warning.Printf("failed to write video config %s", filePath)
		return err
	}

	return nil
}

func (cs *ConfigService) buildAudioConfigurationFilePath() (string, error) {
	var folder, err = cs.buildConfigurationFolderPath()
	if err != nil {
		return "", err
	}

	return path.Join(folder, "audio.json"), nil
}

func (cs *ConfigService) buildInputDeviceConfigurationFilePath(playerIdx int) (string, error) {
	var folder, err = cs.buildConfigurationFolderPath()
	if err != nil {
		return "", err
	}

	return path.Join(folder, fmt.Sprintf("controller-%d.json", playerIdx)), nil
}

func (cs *ConfigService) buildVideoConfigurationFilePath() (string, error) {
	var folder, err = cs.buildConfigurationFolderPath()
	if err != nil {
		return "", err
	}

	return path.Join(folder, "video.json"), nil
}

func (cs *ConfigService) initializeConfigurationFolder() error {
	folderPath, err := cs.buildConfigurationFolderPath()
	if nil != err {
		return err
	}

	if _, err = os.Stat(folderPath); os.IsNotExist(err) {
		err = os.MkdirAll(folderPath, 0700)
		if nil != err {
			logging.Warning.Printf("failed to create config folder: %s", folderPath)
			return err
		}
	}

	return nil
}

func (cs *ConfigService) buildConfigurationFolderPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if nil != err {
		logging.Warning.Printf("failed to calculate config folder path")
		return "", err
	}

	return path.Join(homeDir, ".retro-carnage", "settings"), nil
}
