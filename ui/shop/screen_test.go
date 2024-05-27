package shop

import (
	"testing"

	pixel "github.com/Retro-Carnage-Team/pixel2"
	"github.com/stretchr/testify/assert"
)

func TestGetItemRectOfFirstItem(t *testing.T) {
	var screenSize = pixel.V(3840, 2160)
	var result = getItemRect(screenSize, 0)

	assert.InDelta(t, 10, result.X, 0.0001)
	assert.InDelta(t, 1813.3333, result.Y, 0.0001)     // (2160 - 10 - 336.6666)
	assert.InDelta(t, 756.0, result.Width, 0.0001)     // (3840 - (6 * 10)) / 5
	assert.InDelta(t, 336.6666, result.Height, 0.0001) // (2160 - 70 - (7 * 10)) / 6
}

func TestGetItemRectOfSeventhItem(t *testing.T) {
	var screenSize = pixel.V(3840, 2160)
	var result = getItemRect(screenSize, 6)

	assert.InDelta(t, 776.0, result.X, 0.0001)         // 10 + 756 + 10
	assert.InDelta(t, 1466.6666, result.Y, 0.0001)     // (2160 - 10 - 336.6666 - 10 - 336.6666)
	assert.InDelta(t, 756.0, result.Width, 0.0001)     // (3840 - (6 * 10)) / 5
	assert.InDelta(t, 336.6666, result.Height, 0.0001) // (2160 - 70 - (7 * 10)) / 6
}
