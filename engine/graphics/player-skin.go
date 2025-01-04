package graphics

import "fmt"

var (
	PlayerSkins []*Skin
)

// InitPlayerSkins initializes the player skins. The skins get loaded from the given directory where they are expected
// to be stored as JSON files.
func InitPlayerSkins(skinsDirectory string) {
	if nil == PlayerSkins {
		PlayerSkins = make([]*Skin, 0)
		PlayerSkins = append(PlayerSkins, loadSkin(fmt.Sprintf("%s/player-0.json", skinsDirectory)))
		PlayerSkins = append(PlayerSkins, loadSkin(fmt.Sprintf("%s/player-1.json", skinsDirectory)))
	}
}

// SkinForPlayer returns the skin configuration used for the specified Player.
func SkinForPlayer(playerIdx int) *Skin {
	return PlayerSkins[playerIdx]
}
