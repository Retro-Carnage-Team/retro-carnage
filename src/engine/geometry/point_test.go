package geometry

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToString(t *testing.T) {
	var p = &Point{1, 2}
	assert.Equal(t, "Point[X: 1.00000, Y: 2.00000]", p.String())
}

func TestAdd(t *testing.T) {
	var p1 = &Point{1, 2}
	var p2 = &Point{2, 3}
	var result = p1.Add(p2)
	assert.Equal(t, "Point[X: 3.00000, Y: 5.00000]", result.String())
}

func TestToVec(t *testing.T) {
	var p1 = &Point{1, 2}
	var result = p1.ToVec()
	assert.InDelta(t, 1.0, result.X, 0.00001)
	assert.InDelta(t, 2.0, result.Y, 0.00001)
}
