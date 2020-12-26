package assets

import (
	"encoding/json"
	"io/ioutil"
	"retro-carnage/logging"
	"retro-carnage/util"
	"sync"
)

type GrenadeContainer struct {
	grenades            []*Grenade
	initialized         bool
	initializationMutex sync.Mutex
}

func (gc *GrenadeContainer) GetAll() []*Grenade {
	return gc.grenades
}

func (gc *GrenadeContainer) GetByName(name string) *Grenade {
	for _, grenade := range gc.grenades {
		if grenade.Name == name {
			return grenade
		}
	}
	logging.Error.Fatalf("There is no type of grenade named %s", name)
	return nil
}

var (
	GrenadeCrate = GrenadeContainer{}
)

// Initialize starts the asynchronous initialization. It will load all grenade files from the items/grenades folder.
// You can check the Initialized() method to check when this process has finished.
func (gc *GrenadeContainer) Initialize() {
	gc.initialized = false
	go gc.loadFromDisk("items/grenades/")
}

// InitializeInTest is the test version of the Initialize method. It starts the asynchronous initialization, loading
// grenade files from a given location instead of the default folder. This version doesn't work asynchronously.
// You don't need to check the Initialized() method for when this process has finished.
func (gc *GrenadeContainer) InitializeInTest(folder string) {
	gc.initialized = true
	gc.loadFromDisk(folder)
}

func (gc *GrenadeContainer) Initialized() bool {
	gc.initializationMutex.Lock()
	defer gc.initializationMutex.Unlock()
	return gc.initialized
}

func (gc *GrenadeContainer) loadFromDisk(directory string) {
	var result = make([]*Grenade, 0)

	files, err := util.GetJsonFilesOfDirectory(directory)
	if err != nil {
		logging.Error.Fatalf("failed to access grenade folder: %v", err)
	}

	for _, f := range files {
		grenade, err := gc.loadGrenadeFile(f)
		if err != nil {
			logging.Warning.Printf("failed to load grenade from file %s: %v", f, err)
		}
		result = append(result, grenade)
	}

	gc.grenades = result

	gc.initializationMutex.Lock()
	gc.initialized = true
	gc.initializationMutex.Unlock()
}

func (gc *GrenadeContainer) loadGrenadeFile(filePath string) (*Grenade, error) {
	logging.Trace.Printf("loading grenade: %s", filePath)
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var grenade = &Grenade{}
	err = json.Unmarshal(data, grenade)
	if err != nil {
		return nil, err
	}
	return grenade, nil
}
