package characters

import (
	"github.com/stretchr/testify/assert"
	"retro-carnage/engine/geometry"
	"testing"
)

func TestInitEnemySkins(t *testing.T) {
	InitEnemySkins("testdata/skins")

	assert.Equal(t, 4, len(enemySkins))
	assert.Equal(t, string(WoodlandWithSMG), enemySkins[WoodlandWithSMG].Name)
	assert.Equal(t, 6, len(enemySkins[WoodlandWithSMG].MovementByDirection[geometry.Right.Name]))
	assert.Equal(t, "images/tiles/enemy-0/right/1.png", enemySkins[WoodlandWithSMG].MovementByDirection[geometry.Right.Name][0].SpritePath)
	assert.Equal(t, -63.0, enemySkins[WoodlandWithSMG].MovementByDirection[geometry.Right.Name][0].Offset.X)
	assert.Equal(t, -30.0, enemySkins[WoodlandWithSMG].MovementByDirection[geometry.Right.Name][0].Offset.Y)
}
