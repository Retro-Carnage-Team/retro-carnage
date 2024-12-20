package assets

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	folder := filepath.Join(os.Getenv("RC-ASSETS"), "items/grenades/")
	GrenadeCrate.InitializeInTest(folder)
}

func TestGrenadeCrateGetAll(t *testing.T) {
	assert.Equal(t, 2, len(GrenadeCrate.GetAll()))
}

func TestGrenadeCrateGetByNameValid(t *testing.T) {
	grenade := GrenadeCrate.GetByName("DM51")
	assert.Equal(t, 600, int(grenade.Price))
}
