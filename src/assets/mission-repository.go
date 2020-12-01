package assets

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"retro-carnage/logging"
	"strings"
	"sync"
)

type MissionRepo struct {
	initialized         bool
	initializationMutex sync.Mutex
	Missions            []*Mission
}

var (
	MissionRepository = &MissionRepo{}
)

// Initialize starts the asynchronous initialization. It will load all mission files from the default mission folder.
// You can check the Initialized() method to check when this process has finished.
func (mr *MissionRepo) Initialize() {
	mr.initialized = false
	go mr.loadFromDisk("./missions")
}

// InitializeInTest is the test version of the Initialize method. It starts the asynchronous initialization, loading
// mission files from a given location instead of the default mission folder. This version doesn't work asynchronously.
// You don't need to check the Initialized() method for when this process has finished.
func (mr *MissionRepo) InitializeInTest(folder string) {
	mr.initialized = true
	mr.loadFromDisk(folder)
}

func (mr *MissionRepo) Initialized() bool {
	mr.initializationMutex.Lock()
	defer mr.initializationMutex.Unlock()
	return mr.initialized
}

func (mr *MissionRepo) ByName(name string) (*Mission, error) {
	if !mr.Initialized() {
		return nil, errors.New("mission repository has not been initialized, yet")
	}

	for _, m := range mr.Missions {
		if m.Name == name {
			return m, nil
		}
	}

	return nil, errors.New("there's no such mission")
}

func (mr *MissionRepo) loadFromDisk(directory string) {
	mr.Missions = make([]*Mission, 0)

	files, err := mr.getMissionFiles(directory)
	if err != nil {
		logging.Error.Fatalf("failed to access mission folder: %v", err)
	}

	for _, f := range files {
		mission, err := mr.loadMissionFile(f)
		if err != nil {
			logging.Warning.Printf("failed to load mission from file %s: %v", f, err)
		}
		mr.Missions = append(mr.Missions, mission)
	}

	mr.initializationMutex.Lock()
	mr.initialized = true
	mr.initializationMutex.Unlock()
}

func (mr *MissionRepo) getMissionFiles(directory string) ([]string, error) {
	if !strings.HasSuffix(directory, "/") {
		directory += "/"
	}

	files, err := ioutil.ReadDir(directory)
	if err != nil {
		return nil, err
	}

	var result = make([]string, 0)
	for _, f := range files {
		if !f.IsDir() && strings.HasSuffix(f.Name(), ".json") {
			result = append(result, directory+f.Name())
		}
	}
	return result, nil
}

func (mr *MissionRepo) loadMissionFile(filePath string) (*Mission, error) {
	logging.Trace.Printf("loading mission: %s", filePath)
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var mission = &Mission{}
	err = json.Unmarshal(data, mission)
	if err != nil {
		return nil, err
	}
	return mission, nil
}
