package characters

import "retro-carnage/engine/geometry"

type SkinFrame struct {
	SpritePath string         `json:"sprite"`
	Offset     geometry.Point `json:"offset"`
}
