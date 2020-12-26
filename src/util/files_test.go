package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMissionFilesFindsTestFile(t *testing.T) {
	files, err := GetJsonFilesOfDirectory("testdata/")

	assert.Nil(t, err)
	var found = false
	for _, f := range files {
		found = found || f == "testdata/Berlin.json"
	}
	assert.Equal(t, true, found)
}
