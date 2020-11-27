package assets

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMissionFilesFindsTestFile(t *testing.T) {
	var mr = MissionRepo{}
	files, err := mr.getMissionFiles("./testdata")

	assert.Nil(t, err)
	var found = false
	for _, f := range files {
		found = found || f == "./testdata/Berlin.json"
	}
	assert.Equal(t, true, found)
}

func TestLoadingValidMissionFileSucceeds(t *testing.T) {
	var mr = MissionRepo{}
	mission, err := mr.loadMissionFile("./testdata/Berlin.json")

	assert.Nil(t, err)
	assert.Equal(t, "I am Berlin's hottest influencer.", mission.Briefing)
	assert.Equal(t, "Berlin", mission.Name)
	assert.Equal(t, 2, len(mission.Segments))
}
