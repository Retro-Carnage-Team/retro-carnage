package characters

import (
	"retro-carnage/engine/geometry"
	"testing"

	"github.com/stretchr/testify/assert"
)

const ENEMY0_DOWN1 = "images/enemy-0/down/1.png"
const ENEMY0_DOWN2 = "images/enemy-0/down/2.png"

func TestPersonReturnsSpritesOfAnimation(t *testing.T) {
	InitEnemySkins("testdata/skins")
	var person = buildEnemyPerson()
	var spriteSupplier = NewEnemyPersonSpriteSupplier(*person.ViewingDirection)
	assert.NotNil(t, spriteSupplier)

	assert.Equal(t, ENEMY0_DOWN1, spriteSupplier.Sprite(1, person).Source)
	assert.Equal(t, ENEMY0_DOWN1, spriteSupplier.Sprite(2, person).Source)
	assert.Equal(t, ENEMY0_DOWN2, spriteSupplier.Sprite(DurationOfEnemyMovementFrame, person).Source)
}

func TestPersonReturnsCorrectSpritesForDeathStateTransition(t *testing.T) {
	InitEnemySkins("testdata/skins")
	var person = buildEnemyPerson()
	var spriteSupplier = NewEnemyPersonSpriteSupplier(*person.ViewingDirection)

	assert.NotNil(t, spriteSupplier)
	assert.Equal(t, ENEMY0_DOWN1, spriteSupplier.Sprite(1, person).Source)

	person.Die()

	assert.Equal(t, ENEMY0_DOWN1, spriteSupplier.Sprite(2, person).Source)
	assert.Equal(t, ENEMY0_DOWN1, spriteSupplier.Sprite(DurationOfEnemyDeathAnimationFrame, person).Source)
}

func buildEnemyPerson() ActiveEnemy {
	return ActiveEnemy{
		Dying:                   false,
		DyingAnimationCountDown: 0,
		Movements:               []EnemyMovement{},
		position: geometry.Rectangle{
			X:      100,
			Y:      100,
			Width:  50,
			Height: 50,
		},
		Skin:             WoodlandWithSMG,
		SpriteSupplier:   nil,
		Type:             Person,
		ViewingDirection: &geometry.Down,
	}
}
