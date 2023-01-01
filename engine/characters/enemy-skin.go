package characters

import (
	"fmt"
	"retro-carnage/logging"
)

// EnemySkin typed string representing an enemy skin.
type EnemySkin string

const (
	// WoodlandWithSMG skin of an enemy wearing woodland camouflage and a sub machine gun
	WoodlandWithSMG EnemySkin = "enemy-0"

	// GreyJumperWithRifle skin if an enemy wearing a grey jumpsuit and a rifle
	GreyJumperWithRifle EnemySkin = "enemy-1"

	// DigitalWithPistols skin of an enemy wearing digital camouflage and two pistols
	DigitalWithPistols EnemySkin = "enemy-2"

	// WoodlandWithBulletproofVest skin of an enemy wearing woodland camouflage and a bullet proof vest and a pistol
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

func GetEnemySkin(name EnemySkin) *Skin {
	if !IsEnemySkin(string(name)) {
		logging.Error.Fatalf("No such enemy skin found: %s", name)
	}
	return enemySkins[name]
}

// IsEnemySkin returns whether or no the given name represents an installed enemy skin.
func IsEnemySkin(name string) bool {
	for _, skin := range enemySkinNames {
		if string(skin) == name {
			return true
		}
	}
	return false
}
