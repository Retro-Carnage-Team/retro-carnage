package highscore

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEntryControllerShouldBeInitializedCorrectly(t *testing.T) {
	var entryController = newEntryController()
	assert.Equal(t, 10, len(entryController.entries))

	for i := 1; i < 10; i++ {
		assert.True(t, entryController.entries[i-1].score > entryController.entries[i].score)
	}
}

func TestAddEntryToTop10(t *testing.T) {
	var entryController = newEntryController()
	entryController.addEntry(entry{
		name:  "TEST",
		score: 20_000,
	})
	assert.Equal(t, "TEST", entryController.entries[2].name)
}

func TestAddEntryBelowTop10(t *testing.T) {
	var entryController = newEntryController()
	entryController.addEntry(entry{
		name:  "TEST",
		score: 500,
	})
	for i := 0; i < 10; i++ {
		assert.NotEqual(t, "TEST", entryController.entries[i].name)
	}
}
