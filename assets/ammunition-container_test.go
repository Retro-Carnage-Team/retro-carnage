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
	assert.Equal(t, 10, len(AmmunitionCrate.GetAll()))
}

func TestAmmunitionCrateGetByNameValid(t *testing.T) {
	ammo := AmmunitionCrate.GetByName("9 x 19 mm")
	assert.Equal(t, "9 x 19 mm", ammo.Name)
}
