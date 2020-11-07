package engine

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetDirectionForCardinals(t *testing.T) {
	var result = GetDirectionForCardinals(false, true, true, false)
	assert.Equal(t, DownLeft, *result)
}
