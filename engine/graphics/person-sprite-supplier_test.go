package graphics

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
	var spriteSupplier = NewPersonSpriteSupplier(&person)
	assert.NotNil(t, spriteSupplier)

	assert.Equal(t, ENEMY0_DOWN1, spriteSupplier.Sprite(1).Source)
	assert.Equal(t, ENEMY0_DOWN1, spriteSupplier.Sprite(2).Source)
	assert.Equal(t, ENEMY0_DOWN2, spriteSupplier.Sprite(durationOfEnemyMovementFrame).Source)
}

func TestPersonReturnsCorrectSpritesForDeathStateTransition(t *testing.T) {
	InitEnemySkins("testdata/skins")
	var person = buildEnemyPerson()
	var spriteSupplier = NewPersonSpriteSupplier(&person)

	assert.NotNil(t, spriteSupplier)
	assert.Equal(t, ENEMY0_DOWN1, spriteSupplier.Sprite(1).Source)

	person.dying = true

	assert.Equal(t, ENEMY0_DOWN1, spriteSupplier.Sprite(2).Source)
	assert.Equal(t, ENEMY0_DOWN1, spriteSupplier.Sprite(durationOfEnemyDeathAnimationFrame).Source)
}

func buildEnemyPerson() mockPerson {
	return mockPerson{
		dying:            false,
		skin:             WoodlandWithSMG,
		viewingDirection: geometry.Down,
	}
}

type mockPerson struct {
	dying            bool
	skin             EnemySkin
	viewingDirection geometry.Direction
}

func (mp *mockPerson) Dying() bool {
	return mp.dying
}

func (mp *mockPerson) Skin() EnemySkin {
	return mp.skin
}

func (mp *mockPerson) ViewingDirection() *geometry.Direction {
	return &mp.viewingDirection
}
