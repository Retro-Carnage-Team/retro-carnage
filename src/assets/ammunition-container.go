package assets

import (
	"encoding/json"
	"io/ioutil"
	"retro-carnage/logging"
	"retro-carnage/util"
	"sync"
)

type AmmunitionContainer struct {
	ammunition          []*Ammunition
	initialized         bool
	initializationMutex sync.Mutex
}

func (ac *AmmunitionContainer) GetAll() []*Ammunition {
	return ac.ammunition
}

func (ac *AmmunitionContainer) GetByName(name string) *Ammunition {
	for _, ammo := range ac.ammunition {
		if ammo.Name == name {
			return ammo
		}
	}
	logging.Error.Fatalf("no such ammunition: %s", name)
	return nil
}

var (
	AmmunitionCrate = AmmunitionContainer{}
)

// Initialize starts the asynchronous initialization. It will load all ammunition files from the items/ammunition
// folder. You can check the Initialized() method to check when this process has finished.
func (ac *AmmunitionContainer) Initialize() {
	ac.initialized = false
	go ac.loadFromDisk("items/ammunition/")
}

// InitializeInTest is the test version of the Initialize method. It starts the asynchronous initialization, loading
// ammunition files from a given location instead of the default folder. This version doesn't work asynchronously.
// You don't need to check the Initialized() method for when this process has finished.
func (ac *AmmunitionContainer) InitializeInTest(folder string) {
	ac.initialized = true
	ac.loadFromDisk(folder)
}

func (ac *AmmunitionContainer) Initialized() bool {
	ac.initializationMutex.Lock()
	defer ac.initializationMutex.Unlock()
	return ac.initialized
}

func (ac *AmmunitionContainer) loadFromDisk(directory string) {
	var result = make([]*Ammunition, 0)

	files, err := util.GetJsonFilesOfDirectory(directory)
	if err != nil {
		logging.Error.Fatalf("failed to access ammunition folder: %v", err)
	}

	for _, f := range files {
		ammo, err := ac.loadAmmunitionFile(f)
		if err != nil {
			logging.Warning.Printf("failed to load ammunition from file %s: %v", f, err)
		}
		result = append(result, ammo)
	}
	ac.ammunition = result

	ac.initializationMutex.Lock()
	ac.initialized = true
	ac.initializationMutex.Unlock()
}

func (ac *AmmunitionContainer) loadAmmunitionFile(filePath string) (*Ammunition, error) {
	logging.Trace.Printf("loading ammunition: %s", filePath)
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var ammo = &Ammunition{}
	err = json.Unmarshal(data, ammo)
	if err != nil {
		return nil, err
	}
	return ammo, nil
}
