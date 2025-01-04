package graphics

import (
	"retro-carnage/assets"
	"retro-carnage/engine/geometry"
)

type SkinFrame struct {
	SpritePath string         `json:"sprite"`
	Offset     geometry.Point `json:"offset"`
}

func (sf *SkinFrame) ToSpriteWithOffset() *SpriteWithOffset {
	return &SpriteWithOffset{
		Offset: sf.Offset,
		Source: sf.SpritePath,
		Sprite: assets.SpriteRepository.Get(sf.SpritePath),
	}
}
