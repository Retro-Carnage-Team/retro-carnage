package geometry

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDirectionForCardinals(t *testing.T) {
	var result = GetDirectionForCardinals(false, true, true, false)
	assert.Equal(t, DownLeft, *result)
}

func TestIsDiagonal(t *testing.T) {
	assert.Equal(t, false, Up.IsDiagonal())
	assert.Equal(t, true, UpRight.IsDiagonal())
	assert.Equal(t, false, Right.IsDiagonal())
	assert.Equal(t, true, DownRight.IsDiagonal())
	assert.Equal(t, false, Down.IsDiagonal())
	assert.Equal(t, true, DownLeft.IsDiagonal())
	assert.Equal(t, false, Left.IsDiagonal())
	assert.Equal(t, true, UpLeft.IsDiagonal())
}
