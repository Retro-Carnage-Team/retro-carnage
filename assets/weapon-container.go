package assets

import (
	"encoding/json"
	"os"
	"retro-carnage/logging"
	"retro-carnage/util"
	"sync"
)

type WeaponContainer struct {
	initialized         bool
	initializationMutex sync.Mutex
	weapons             []*Weapon
}

func (wc *WeaponContainer) GetAll() []*Weapon {
	return wc.weapons
}

func (wc *WeaponContainer) GetByName(name string) *Weapon {
	for _, weapon := range wc.weapons {
		if weapon.Name == name {
			return weapon
		}
	}
	logging.Error.Fatalf("failed to find weapon named %s", name)
	return nil
}

var (
	WeaponCrate = WeaponContainer{}
)

// Initialize starts the asynchronous initialization. It will load all ammunition files from the items/ammunition
// folder. You can check the Initialized() method to check when this process has finished.
func (wc *WeaponContainer) Initialize() {
	wc.initialized = false
	go wc.loadFromDisk("items/weapons/")
}

// InitializeInTest is the test version of the Initialize method. It starts the asynchronous initialization, loading
// ammunition files from a given location instead of the default folder. This version doesn't work asynchronously.
// You don't need to check the Initialized() method for when this process has finished.
func (wc *WeaponContainer) InitializeInTest(folder string) {
	wc.initialized = true
	wc.loadFromDisk(folder)
}

func (wc *WeaponContainer) Initialized() bool {
	wc.initializationMutex.Lock()
	defer wc.initializationMutex.Unlock()
	return wc.initialized
}

func (wc *WeaponContainer) loadFromDisk(directory string) {
	var result = make([]*Weapon, 0)

	files, err := util.GetJsonFilesOfDirectory(directory)
	if err != nil {
		logging.Error.Fatalf("failed to access weapons folder: %v", err)
	}

	for _, f := range files {
		ammo, err := wc.loadWeaponFile(f)
		if err != nil {
			logging.Warning.Printf("failed to load weapon from file %s: %v", f, err)
		}
		result = append(result, ammo)
	}
	wc.weapons = result

	wc.initializationMutex.Lock()
	wc.initialized = true
	wc.initializationMutex.Unlock()
}

func (wc *WeaponContainer) loadWeaponFile(filePath string) (*Weapon, error) {
	logging.Trace.Printf("loading weapon: %s", filePath)
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var weapon = &Weapon{}
	err = json.Unmarshal(data, weapon)
	if err != nil {
		return nil, err
	}
	return weapon, nil
}
