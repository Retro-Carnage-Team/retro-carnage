package engine

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntersectionHorizontalAndVertical(t *testing.T) {
	var lineA = NewLine(NewPoint(4, 0), NewPoint(4, 239))
	var lineB = NewLine(NewPoint(8, 20), NewPoint(2, 20))
	var intersection = lineA.GetIntersection(lineB)
	var expected = NewPoint(4, 20)

	assert.Equal(t, expected.String(), intersection.String())
}

func TestIntersectionHorizontalAndDiagonal(t *testing.T) {
	var lineA = NewLine(NewPoint(2, 4), NewPoint(7, 4))
	var lineB = NewLine(NewPoint(2, 6), NewPoint(6, 2))
	var intersection = lineA.GetIntersection(lineB)
	var expected = NewPoint(4, 4)

	assert.Equal(t, expected.String(), intersection.String())
}

func TestLineEqualsItself(t *testing.T) {
	var lineA = NewLine(NewPoint(2, 4), NewPoint(7, 4))

	assert.True(t, lineA.Equals(lineA))
}

func TestLineEqualsReversedSelf(t *testing.T) {
	var lineA = NewLine(NewPoint(2, 4), NewPoint(7, 4))
	var lineB = NewLine(NewPoint(7, 4), NewPoint(2, 4))

	assert.True(t, lineA.Equals(lineB))
}

func TestLineNotEqualsOther(t *testing.T) {
	var lineA = NewLine(NewPoint(3, 4), NewPoint(7, 4))
	var lineB = NewLine(NewPoint(7, 4), NewPoint(2, 4))

	assert.False(t, lineA.Equals(lineB))
}
