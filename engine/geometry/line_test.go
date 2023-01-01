package geometry

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntersectionHorizontalAndVertical(t *testing.T) {
	var lineA = &Line{Start: &Point{4, 0}, End: &Point{4, 239}}
	var lineB = &Line{Start: &Point{8, 20}, End: &Point{2, 20}}
	var intersection = lineA.GetIntersection(lineB)
	var expected = &Point{4, 20}

	assert.Equal(t, expected.String(), intersection.String())
}

func TestIntersectionHorizontalAndDiagonal(t *testing.T) {
	var lineA = &Line{Start: &Point{2, 4}, End: &Point{7, 4}}
	var lineB = &Line{Start: &Point{2, 6}, End: &Point{6, 2}}
	var intersection = lineA.GetIntersection(lineB)
	var expected = &Point{4, 4}

	assert.Equal(t, expected.String(), intersection.String())
}

func TestLineEqualsItself(t *testing.T) {
	var lineA = &Line{Start: &Point{2, 4}, End: &Point{7, 4}}

	assert.True(t, lineA.Equals(lineA))
}

func TestLineEqualsReversedSelf(t *testing.T) {
	var lineA = &Line{Start: &Point{2, 4}, End: &Point{7, 4}}
	var lineB = &Line{Start: &Point{7, 4}, End: &Point{2, 4}}

	assert.True(t, lineA.Equals(lineB))
}

func TestLineNotEqualsOther(t *testing.T) {
	var lineA = &Line{Start: &Point{3, 4}, End: &Point{7, 4}}
	var lineB = &Line{Start: &Point{7, 4}, End: &Point{2, 4}}

	assert.False(t, lineA.Equals(lineB))
}
