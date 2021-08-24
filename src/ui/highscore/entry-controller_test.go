package highscore

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEntryControllerShouldBeInitializedCorrectly(t *testing.T) {
	var entryController = newEntryController()
	assert.Equal(t, 10, len(entryController.entries))

	for i := 1; i < 10; i++ {
		assert.True(t, entryController.entries[i-1].Score > entryController.entries[i].Score)
	}
}

func TestAddEntryToTop10(t *testing.T) {
	var entryController = newEntryController()
	entryController.AddEntry(Entry{
		Name:  "TEST",
		Score: 20_000,
	})
	assert.Equal(t, "TEST", entryController.entries[2].Name)
}

func TestAddEntryBelowTop10(t *testing.T) {
	var entryController = newEntryController()
	entryController.AddEntry(Entry{
		Name:  "TEST",
		Score: 500,
	})
	for i := 0; i < 10; i++ {
		assert.NotEqual(t, "TEST", entryController.entries[i].Name)
	}
}

func TestReachedHighScoreJustBelowScore(t *testing.T) {
	var entryController = newEntryController()
	var p1, p2 = entryController.reachedHighScore(2_500, 2_499)
	assert.Equal(t, false, p1)
	assert.Equal(t, false, p2)
	assert.Equal(t, "Drummer", entryController.entries[len(entryController.entries)-1].Name)
}

func TestReachedHighScoreOnlyP1ReachedIt(t *testing.T) {
	var entryController = newEntryController()
	var p1, p2 = entryController.reachedHighScore(2_750, 2_200)
	assert.Equal(t, true, p1)
	assert.Equal(t, false, p2)
	assert.Equal(t, "Drummer", entryController.entries[len(entryController.entries)-1].Name)
}

func TestReachedHighScoreOnlyP2ReachedIt(t *testing.T) {
	var entryController = newEntryController()
	var p1, p2 = entryController.reachedHighScore(2_000, 2_750)
	assert.Equal(t, false, p1)
	assert.Equal(t, true, p2)
	assert.Equal(t, "Drummer", entryController.entries[len(entryController.entries)-1].Name)
}
