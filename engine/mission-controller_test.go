package engine

import (
	"github.com/stretchr/testify/assert"
	"retro-carnage/assets"
	"retro-carnage/engine/characters"
	"testing"
)

func setUp() (mc *MissionCtrl) {
	assets.MissionRepository.InitializeInTest("./testdata/missions")
	mc = &MissionCtrl{}
	mc.Reset()
	return mc
}

func TestFullSetOfMissionsShouldBeAvailableWhenGameGetsReset(t *testing.T) {
	var mc = setUp()
	var remainingMission, err = mc.RemainingMissions()

	assert.Nil(t, err)
	assert.Equal(t, 12, len(remainingMission))
}

func TestMissionsShouldGetFilteredWhenMarkedAsFinished(t *testing.T) {
	var mc = setUp()
	mc.MarkMissionFinished(assets.MissionRepository.Missions[0])
	var remainingMission, err = mc.RemainingMissions()

	assert.Nil(t, err)
	assert.Equal(t, 11, len(remainingMission))
}

func TestSelectingMissionShouldUpdatePropertyAndCashInInventory(t *testing.T) {
	var mc = setUp()

	var aMission = assets.MissionRepository.Missions[0]
	assert.NotZero(t, aMission.Reward)

	var oldCash = characters.Players[0].Cash()
	mc.SelectMission(aMission)

	assert.Equal(t, aMission.Name, mc.currentMission.Name)
	assert.Greater(t, characters.Players[0].Cash(), oldCash)
}

func TestShouldFindMissionNorthOfGivenMission(t *testing.T) {
	var mc = setUp()

	berlin, err := assets.MissionRepository.ByName("Berlin")
	assert.Nil(t, err)
	assert.NotNil(t, berlin)

	minsk, err := mc.NextMissionNorth(&berlin.Location)
	assert.Nil(t, err)
	assert.Equal(t, "Minsk", minsk.Name)
}

func TestShouldReturnNilIfUnableToFindAnyMissionNorthOfGivenLocation(t *testing.T) {
	var mc = setUp()

	minsk, err := assets.MissionRepository.ByName("Minsk")
	assert.Nil(t, err)
	assert.Equal(t, "Minsk", minsk.Name)

	result, err := mc.NextMissionNorth(&minsk.Location)
	assert.Nil(t, err)
	assert.Nil(t, result)
}

func TestShouldFindMissionSouthOfGivenMission(t *testing.T) {
	var mc = setUp()

	berlin, err := assets.MissionRepository.ByName("Berlin")
	assert.Nil(t, err)
	assert.NotNil(t, berlin)

	bischkek, err := mc.NextMissionSouth(&berlin.Location)
	assert.Nil(t, err)
	assert.Equal(t, "Bischkek", bischkek.Name)
}

func TestShouldReturnNilIfUnableToFindAnyMissionSouthOfGivenLocation(t *testing.T) {
	var mc = setUp()

	amazonas, err := assets.MissionRepository.ByName("Amazonas")
	assert.Nil(t, err)
	assert.Equal(t, "Amazonas", amazonas.Name)

	result, err := mc.NextMissionSouth(&amazonas.Location)
	assert.Nil(t, err)
	assert.Nil(t, result)
}

func TestShouldFindMissionEastToGivenMission(t *testing.T) {
	var mc = setUp()

	berlin, err := assets.MissionRepository.ByName("Berlin")
	assert.Nil(t, err)
	assert.NotNil(t, berlin)

	result, err := mc.NextMissionEast(&berlin.Location)
	assert.Nil(t, err)
	assert.Equal(t, "Tripolis", result.Name)
}

func TestShouldReturnNilIfUnableToFindAnyMissionEastOfGivenLocation(t *testing.T) {
	var mc = setUp()

	vietnam, err := assets.MissionRepository.ByName("Ho-Chi-Minh City")
	assert.Nil(t, err)
	assert.Equal(t, "Ho-Chi-Minh City", vietnam.Name)

	result, err := mc.NextMissionEast(&vietnam.Location)
	assert.Nil(t, err)
	assert.Nil(t, result)
}

func TestShouldFindMissionWestToGivenMission(t *testing.T) {
	var mc = setUp()

	berlin, err := assets.MissionRepository.ByName("Berlin")
	assert.Nil(t, err)
	assert.NotNil(t, berlin)

	result, err := mc.NextMissionWest(&berlin.Location)
	assert.Nil(t, err)
	assert.Equal(t, "Amazonas", result.Name)
}

func TestShouldReturnNilIfUnableToFindAnyMissionWestOfGivenLocation(t *testing.T) {
	var mc = setUp()

	vietnam, err := assets.MissionRepository.ByName("Mexico City")
	assert.Nil(t, err)
	assert.Equal(t, "Mexico City", vietnam.Name)

	result, err := mc.NextMissionWest(&vietnam.Location)
	assert.Nil(t, err)
	assert.Nil(t, result)
}
