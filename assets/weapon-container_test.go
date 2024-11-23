package assets

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	folder := filepath.Join(os.Getenv("RC-ASSETS"), "items/weapons/")
	WeaponCrate.InitializeInTest(folder)
}

func TestWeaponCrateGetAll(t *testing.T) {
	assert.Equal(t, 28, len(WeaponCrate.GetAll()))
}

func TestWeaponCrateGetByNameValid(t *testing.T) {
	weapon := WeaponCrate.GetByName("RPG-7")
	assert.Equal(t, "85 mm", weapon.Ammo)
	assert.Equal(t, 0, weapon.BulletInterval)
	assert.Equal(t, 1200, weapon.BulletRange)
	assert.InDelta(t, 1.0, weapon.BulletSpeed, 0.0001)
	assert.True(t, strings.HasPrefix(weapon.Description, "The RPG-7 is a"))
	assert.Equal(t, "images/weapons/10.png", weapon.Image)
	assert.Equal(t, "images/weapons/10-r.png", weapon.ImageRotated)
	assert.Equal(t, "100.0 cm", weapon.Length)
	assert.Equal(t, 5800, weapon.Price)
	assert.Equal(t, FxRocketLauncher, weapon.Sound)
	assert.Equal(t, RPG, weapon.WeaponType)
	assert.Equal(t, "6.3 kg", weapon.Weight)
}
