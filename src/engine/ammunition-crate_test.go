package engine

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAll(t *testing.T) {
	assert.Equal(t, 10, len(AmmunitionCrate.GetAll()))
}

func TestGetByNameValid(t *testing.T) {
	ammo, err := AmmunitionCrate.GetByName("9 x 19 mm")
	assert.Nil(t, err)
	assert.Equal(t, "9 x 19 mm", ammo.name)
}
