package characters

import (
	"retro-carnage/assets"
	"retro-carnage/engine/geometry"
	"retro-carnage/engine/graphics"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	spawn_delay = 50
)

func TestSpawnWithLimit(t *testing.T) {
	// Capacity of spawn area is 3
	spawnArea := createSpawnArea(3)

	// The first two calls should return an enemy and indicate non-depletion.
	for i := 0; i < 2; i++ {
		enemy, depleted := spawnArea.Spawn(spawn_delay)
		assert.NotNil(t, enemy)
		assert.Equal(t, false, depleted)
	}

	// Third call should return an enemy and indicate depletion.
	enemy, depleted := spawnArea.Spawn(spawn_delay)
	assert.NotNil(t, enemy)
	assert.Equal(t, true, depleted)

	// Further calls should return nil and indicate depletion.
	for i := 0; i < 2; i++ {
		enemy, depleted = spawnArea.Spawn(spawn_delay)
		assert.Nil(t, enemy)
		assert.Equal(t, true, depleted)
	}
}

func TestSpawnWithoutLimit(t *testing.T) {
	spawnArea := createSpawnArea(0)
	for i := 0; i < 10; i++ {
		enemy, depleted := spawnArea.Spawn(spawn_delay)
		assert.NotNil(t, enemy)
		assert.Equal(t, false, depleted)
	}
}

func createSpawnArea(capacity int) *ActiveEnemy {
	var enemyType = enemyTypeSpawnArea
	var result = ActiveEnemy{
		Actions:                 make([]assets.EnemyAction, 0),
		Dying:                   false,
		DyingAnimationCountDown: 0,
		Movements:               make([]EnemyMovement, 0),
		Skin:                    graphics.WoodlandWithSMG,
		SpawnCapacity:           capacity,
		SpawnDelays:             append(make([]int64, 0), spawn_delay),
		SpriteSupplier:          nil,
		Type:                    enemyTypeSpawnArea,
		ViewingDirection:        &geometry.Down,
	}
	result.SetPosition(&geometry.Rectangle{})
	enemyType.OnActivation(&result)
	return &result
}
