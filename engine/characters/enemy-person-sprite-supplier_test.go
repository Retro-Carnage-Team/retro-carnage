package characters

import (
	"github.com/stretchr/testify/assert"
	"retro-carnage/engine/geometry"
	"testing"
)

func TestPersonReturnsSpritesOfAnimation(t *testing.T) {
	InitEnemySkins("testdata/skins")
	var person = buildEnemyPerson()
	var spriteSupplier = NewEnemyPersonSpriteSupplier(*person.ViewingDirection)
	assert.NotNil(t, spriteSupplier)

	assert.Equal(t, "images/enemy-0/down/1.png", spriteSupplier.Sprite(1, person).Source)
	assert.Equal(t, "images/enemy-0/down/1.png", spriteSupplier.Sprite(2, person).Source)
	assert.Equal(t, "images/enemy-0/down/2.png", spriteSupplier.Sprite(DurationOfEnemyMovementFrame, person).Source)
}

func TestPersonReturnsCorrectSpritesForDeathStateTransition(t *testing.T) {
	InitEnemySkins("testdata/skins")
	var person = buildEnemyPerson()
	var spriteSupplier = NewEnemyPersonSpriteSupplier(*person.ViewingDirection)

	assert.NotNil(t, spriteSupplier)
	assert.Equal(t, "images/enemy-0/down/1.png", spriteSupplier.Sprite(1, person).Source)

	person.Die()

	assert.Equal(t, "images/enemy-0/down/1.png", spriteSupplier.Sprite(2, person).Source)
	assert.Equal(t, "images/enemy-0/down/1.png", spriteSupplier.Sprite(DurationOfEnemyDeathAnimationFrame, person).Source)
}

func buildEnemyPerson() ActiveEnemy {
	return ActiveEnemy{
		Dying:                   false,
		DyingAnimationCountDown: 0,
		Movements:               []*EnemyMovement{},
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
