package characters

import "fmt"

type EnemySkin string

const (
	WoodlandWithSMG             EnemySkin = "enemy-0"
	GreyJumperWithRifle         EnemySkin = "enemy-1"
	DigitalWithPistols          EnemySkin = "enemy-2"
	WoodlandWithBulletproofVest EnemySkin = "enemy-3"
)

var (
	enemySkins     map[EnemySkin]*Skin
	enemySkinNames = []EnemySkin{WoodlandWithSMG, GreyJumperWithRifle, DigitalWithPistols, WoodlandWithBulletproofVest}
)

// InitEnemySkins initializes the enemy skins. The skins get loaded from the given directory where they are expected to
// be stored as JSON files.
func InitEnemySkins(skinsDirectory string) {
	if nil == enemySkins {
		enemySkins = make(map[EnemySkin]*Skin)
		for _, skin := range enemySkinNames {
			enemySkins[skin] = loadSkin(fmt.Sprintf("%s/%s.json", skinsDirectory, skin))
		}
	}
}

func IsEnemySkin(name string) bool {
	for _, skin := range enemySkinNames {
		if string(skin) == name {
			return true
		}
	}
	return false
}
