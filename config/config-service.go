package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"retro-carnage/logging"
)

type ConfigService struct{}

// loadControllerConfigurations reads the controller configurations that is stored on disk.
// Returns empty array of not configurations can be found.
func (cs *ConfigService) LoadInputDeviceConfigurations() []InputDeviceConfiguration {
	var result = make([]InputDeviceConfiguration, 0)
	for i := 0; i < 2; i++ {
		filePath, err := cs.buildConfigurationFilePath(i)
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
	folderPath, err := cs.buildConfigurationFolderPath()
	if nil != err {
		return err
	}

	if _, err = os.Stat(folderPath); os.IsNotExist(err) {
		err = os.MkdirAll(folderPath, 0700)
		if nil != err {
			logging.Warning.Printf("failed to create folder for configurations: %s", folderPath)
			return err
		}
	}

	filePath, err := cs.buildConfigurationFilePath(playerIdx)
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

func (cs *ConfigService) buildConfigurationFilePath(playerIdx int) (string, error) {
	var folder, err = cs.buildConfigurationFolderPath()
	if err != nil {
		return "", err
	}

	return path.Join(folder, fmt.Sprintf("controller-%d.json", playerIdx)), nil
}

func (cs *ConfigService) buildConfigurationFolderPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if nil != err {
		logging.Warning.Printf("failed to calculate config folder path")
		return "", err
	}

	return path.Join(homeDir, ".retro-carnage", "settings"), nil
}
