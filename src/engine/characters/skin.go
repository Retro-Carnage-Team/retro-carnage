package characters

import (
	"encoding/json"
	"io/ioutil"
	"retro-carnage/logging"
)

type Skin struct {
	DeathAnimation      []SkinFrame            `json:"deathAnimation"`
	Idle                map[string]SkinFrame   `json:"idle"`
	MovementByDirection map[string][]SkinFrame `json:"movement"`
	Name                string                 `json:"name"`
}

// DurationOfDeathAnimation returns the duration of a death animation in milliseconds
func (s *Skin) DurationOfDeathAnimation() int64 {
	return int64(len(s.DeathAnimation) * DurationOfPlayerDeathAnimationFrame)
}

func loadSkin(filePath string) *Skin {
	logging.Trace.Printf("loading skin: %s", filePath)
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		logging.Error.Fatalf("failed to read skin file %s: %v", filePath, err)
	}

	var skin = Skin{}
	err = json.Unmarshal(data, &skin)
	if err != nil {
		logging.Error.Fatalf("failed to parse skin file %s: %v", filePath, err)
	}
	return &skin
}
