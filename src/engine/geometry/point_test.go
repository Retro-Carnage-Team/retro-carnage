package geometry

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToString(t *testing.T) {
	var p = &Point{1, 2}
	assert.Equal(t, "Point[X: 1.00000, Y: 2.00000]", p.String())
}
