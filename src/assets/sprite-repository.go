package assets

import (
	"fmt"
	"github.com/faiface/pixel"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"os"
	"retro-carnage/logging"
	"retro-carnage/util"
	"strings"
	"sync"
)

type SpriteRepo struct {
	initializationMutex sync.Mutex
	initialized         bool
	tiles               map[string]*pixel.Sprite
}

var (
	SpriteRepository = &SpriteRepo{
		initialized: false,
		tiles:       make(map[string]*pixel.Sprite),
	}
)

func (sr *SpriteRepo) Initialized() bool {
	sr.initializationMutex.Lock()
	defer sr.initializationMutex.Unlock()
	return sr.initialized
}

func (sr *SpriteRepo) Initialize() {
	if !sr.Initialized() {
		go sr.loadFromDirectory("images/")

		sr.initializationMutex.Lock()
		sr.initialized = true
		sr.initializationMutex.Unlock()
	}
}

func (sr *SpriteRepo) Get(path string) *pixel.Sprite {
	return sr.tiles[path]
}

func (sr *SpriteRepo) loadFromDirectory(directory string) {
	if !strings.HasSuffix(directory, "/") {
		directory += "/"
	}

	files, err := ioutil.ReadDir(directory)
	if err != nil {
		logging.Warning.Fatalf("failed to read directory %s: %v", directory, err)
	}

	for _, f := range files {
		if f.IsDir() {
			sr.loadFromDirectory(directory + f.Name() + "/")
		} else {
			var filePath = directory + f.Name()
			var picture = sr.loadPicture(filePath)
			sr.tiles[filePath] = pixel.NewSprite(picture, picture.Bounds())
		}
	}
}

func (sr *SpriteRepo) loadPicture(filePath string) pixel.Picture {
	stopWatch := util.StopWatch{Name: fmt.Sprintf("Loading sprite: %s", filePath)}
	stopWatch.Start()

	file, err := os.Open(filePath)
	if err != nil {
		logging.Error.Fatalf("Failed to load image file %s: %v", filePath, err)
		return nil
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		logging.Error.Fatalf("Failed to decode image file %s: %v", filePath, err)
		return nil
	}
	var result = pixel.PictureDataFromImage(img)

	_ = stopWatch.Stop()
	logging.Trace.Println(stopWatch.PrintDebugMessage())
	return result
}
