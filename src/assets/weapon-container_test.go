package assets

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWeaponCrateGetAll(t *testing.T) {
	assert.Equal(t, 18, len(WeaponCrate.GetAll()))
}

func TestWeaponCrateGetByNameValid(t *testing.T) {
	weapon, err := WeaponCrate.GetByName("Panzerfaust 44")
	assert.Nil(t, err)
	assert.Equal(t, RPG, weapon.WeaponType())
}