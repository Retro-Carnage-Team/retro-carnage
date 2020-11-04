package engine

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAccessAmmunition(t *testing.T) {
	ammoList := Ammunitions

	assert.Equal(t, 10, len(ammoList))
	assert.Equal(t, "9 x 19 mm", ammoList[0].Name)
}
