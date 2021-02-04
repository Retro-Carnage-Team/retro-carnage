package characters

import (
	"github.com/stretchr/testify/assert"
	"retro-carnage/engine/geometry"
	"testing"
)

func TestLandmineReturnsStaticSprite(t *testing.T) {
	var landmine = Enemy{
		ActivationDistance:      0,
		Active:                  false,
		Dying:                   false,
		DyingAnimationCountDown: 0,
		Movements:               []EnemyMovement{},
		Position: &geometry.Rectangle{
			X:      100,
			Y:      100,
			Width:  50,
			Height: 50,
		},
		Skin:             "",
		SpriteSupplier:   nil,
		Type:             Landmine,
		ViewingDirection: geometry.Down,
	}
	landmine.Activate() // happens when the landmine becomes visible

	var spriteSupplier = landmine.SpriteSupplier
	assert.NotNil(t, spriteSupplier)

	var sprite = spriteSupplier.Sprite(0, landmine)
	assert.NotNil(t, sprite)
	assert.Equal(t, landmineSprite, sprite.Source)

	sprite = spriteSupplier.Sprite(DurationOfMovementFrame*1.4, landmine)
	assert.NotNil(t, sprite)
	assert.Equal(t, landmineSprite, sprite.Source)
}

func TestPersonReturnsSpritesOfAnimation(t *testing.T) {
	InitEnemySkins("testdata/skins")
	var person = buildEnemyPerson()
	person.Activate() // happens when the landmine becomes visible

	var spriteSupplier = person.SpriteSupplier
	assert.NotNil(t, spriteSupplier)

	assert.Equal(t, "images/tiles/enemy-0/down/1.png", spriteSupplier.Sprite(1, person).Source)
	assert.Equal(t, "images/tiles/enemy-0/down/1.png", spriteSupplier.Sprite(2, person).Source)
	assert.Equal(t, "images/tiles/enemy-0/down/2.png", spriteSupplier.Sprite(DurationOfMovementFrame, person).Source)
}

func TestPersonReturnsCorrectSpritesForDeathStateTransition(t *testing.T) {
	InitEnemySkins("testdata/skins")
	var person = buildEnemyPerson()
	person.Activate() // happens when the landmine becomes visible

	var spriteSupplier = person.SpriteSupplier
	assert.NotNil(t, spriteSupplier)
	assert.Equal(t, "images/tiles/enemy-0/down/1.png", spriteSupplier.Sprite(1, person).Source)

	person.Die()

	assert.Equal(t, "images/tiles/enemy-0/death/1.png", spriteSupplier.Sprite(2, person).Source)
	assert.Equal(t, "images/tiles/enemy-0/death/2.png", spriteSupplier.Sprite(DurationOfDeathAnimationFrame, person).Source)
}

func buildEnemyPerson() Enemy {
	return Enemy{
		ActivationDistance:      0,
		Active:                  false,
		Dying:                   false,
		DyingAnimationCountDown: 0,
		Movements:               []EnemyMovement{},
		Position: &geometry.Rectangle{
			X:      100,
			Y:      100,
			Width:  50,
			Height: 50,
		},
		Skin:             WoodlandWithSMG,
		SpriteSupplier:   nil,
		Type:             Person,
		ViewingDirection: geometry.Down,
	}
}
