package assets

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadingValidMissionFileSucceeds(t *testing.T) {
	var mr = MissionRepo{}
	mission, err := mr.loadMissionFile("testdata/Berlin.json")

	assert.Nil(t, err)
	assert.Equal(t, "I am Berlin's hottest influencer.", mission.Briefing)
	assert.Equal(t, "Berlin", mission.Name)
	assert.Equal(t, 2, len(mission.Segments))
}
