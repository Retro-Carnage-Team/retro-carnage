package assets

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGrenadeCrateGetAll(t *testing.T) {
	assert.Equal(t, 2, len(GrenadeCrate.GetAll()))
}

func TestGrenadeCrateGetByNameValid(t *testing.T) {
	grenade, err := GrenadeCrate.GetByName("DM41")
	assert.Nil(t, err)
	assert.Equal(t, 500, int(grenade.Price()))
}
