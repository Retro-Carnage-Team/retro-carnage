package characters

import (
	"retro-carnage/assets"
	"retro-carnage/engine/geometry"
	"retro-carnage/engine/graphics"
)

type SkinFrame struct {
	SpritePath string         `json:"sprite"`
	Offset     geometry.Point `json:"offset"`
}

func (sf *SkinFrame) ToSpriteWithOffset() *graphics.SpriteWithOffset {
	return &graphics.SpriteWithOffset{
		Offset: sf.Offset,
		Sprite: assets.SpriteRepository.Get(sf.SpritePath),
	}
}
