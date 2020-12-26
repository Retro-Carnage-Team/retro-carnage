package assets

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func init() {
	WeaponCrate.InitializeInTest("../../items/weapons/")
}

func TestWeaponCrateGetAll(t *testing.T) {
	assert.Equal(t, 18, len(WeaponCrate.GetAll()))
}

func TestWeaponCrateGetByNameValid(t *testing.T) {
	weapon := WeaponCrate.GetByName("Panzerfaust 44")
	assert.Equal(t, "60 mm", weapon.Ammo)
	assert.Equal(t, 0, weapon.BulletInterval)
	assert.Equal(t, 900, weapon.BulletRange)
	assert.InDelta(t, 1.2, weapon.BulletSpeed, 0.0001)
	assert.True(t, strings.HasPrefix(weapon.Description, "The Panzerfaust 44 2A1"))
	assert.Equal(t, "images/weapons/Panzerfaust-44.png", weapon.Image)
	assert.Equal(t, "images/weapons/Panzerfaust-44-r.png", weapon.ImageRotated)
	assert.Equal(t, "88 cm", weapon.Length)
	assert.Equal(t, 5800, weapon.Price)
	assert.Equal(t, FxRocketLauncher, weapon.Sound)
	assert.Equal(t, RPG, weapon.WeaponType)
	assert.Equal(t, "7.8 kg", weapon.Weight)
}
