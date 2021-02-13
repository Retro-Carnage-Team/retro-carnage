package characters

import "fmt"

var (
	playerSkins []*Skin
)

// InitPlayerSkins initializes the player skins. The skins get loaded from the given directory where they are expected
// to be stored as JSON files.
func InitPlayerSkins(skinsDirectory string) {
	if nil == playerSkins {
		playerSkins = make([]*Skin, 0)
		playerSkins = append(playerSkins, loadSkin(fmt.Sprintf("%s/player-0.json", skinsDirectory)))
		playerSkins = append(playerSkins, loadSkin(fmt.Sprintf("%s/player-1.json", skinsDirectory)))
	}
}
