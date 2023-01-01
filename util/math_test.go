package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxInt(t *testing.T) {
	assert.Equal(t, 3, MaxInt(2, 3))
	assert.Equal(t, 3, MaxInt(3, 2))
	assert.Equal(t, 3, MaxInt(3, 3))
}

func TestMinInt(t *testing.T) {
	assert.Equal(t, 2, MinInt(2, 3))
	assert.Equal(t, 2, MinInt(3, 2))
	assert.Equal(t, 2, MinInt(2, 2))
}

func TestMax(t *testing.T) {
	assert.InDelta(t, 2.0, Max([]float64{0.1, 0.7, 1.6, 2.0, 0.456}), 0.0001)
}
