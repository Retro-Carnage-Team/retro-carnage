package engine

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToString(t *testing.T) {
	var p = &Point{1, 2}

	assert.Equal(t, "1.00000/2.00000", p.String())
}
