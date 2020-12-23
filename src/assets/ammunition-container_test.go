package assets

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAmmunitionCrateGetAll(t *testing.T) {
	assert.Equal(t, 10, len(AmmunitionCrate.GetAll()))
}

func TestAmmunitionCrateGetByNameValid(t *testing.T) {
	ammo := AmmunitionCrate.GetByName("9 x 19 mm")
	assert.Equal(t, "9 x 19 mm", ammo.name)
}
