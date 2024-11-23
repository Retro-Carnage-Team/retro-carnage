package assets

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	folder := filepath.Join(os.Getenv("RC-ASSETS"), "items/ammunition/")
	AmmunitionCrate.InitializeInTest(folder)
}

func TestAmmunitionCrateGetAll(t *testing.T) {
	// Attention: This test depends on elements specified in the retro-carnage-assets repository.
	assert.Equal(t, 13, len(AmmunitionCrate.GetAll()))
}

func TestAmmunitionCrateGetByNameValid(t *testing.T) {
	ammo := AmmunitionCrate.GetByName("9 x 19 mm")
	assert.Equal(t, "9 x 19 mm", ammo.Name)
}
