package graphics

import (
	"encoding/json"
	"os"
	"retro-carnage/engine/geometry"
	"retro-carnage/logging"
)

// Skin is the configuration of the visual appearance of a Player or an Enemy.
type Skin struct {
	BulletOffsets       map[string]geometry.Point `json:"bulletOffsets"`
	DeathAnimation      map[string][]SkinFrame    `json:"deathAnimation"`
	Idle                map[string]SkinFrame      `json:"idle"`
	MovementByDirection map[string][]SkinFrame    `json:"movement"`
	Name                string                    `json:"name"`
	RpgOffsets          map[string]geometry.Point `json:"rpgOffsets"`
}

// DurationOfDeathAnimation returns the duration of a death animation in milliseconds
func (s *Skin) DurationOfDeathAnimation(direction geometry.Direction) int64 {
	return int64(len(s.DeathAnimation[direction.Name]) * DurationOfPlayerDeathAnimationFrame)
}

func loadSkin(filePath string) *Skin {
	logging.Trace.Printf("loading skin: %s", filePath)
	data, err := os.ReadFile(filePath)
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
