package characters

import (
	"encoding/json"
	"io/ioutil"
	"retro-carnage/logging"
)

type Skin struct {
	FramesByDirection map[string][]SkinFrame `json:"frames"`
	Name              string                 `json:"name"`
}

func loadSkin(filePath string) Skin {
	logging.Trace.Printf("loading skin: %s", filePath)
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		logging.Error.Fatalf("failed to read skin file %s: %v", filePath, err)
	}

	var skin = Skin{}
	err = json.Unmarshal(data, &skin)
	if err != nil {
		logging.Error.Fatalf("failed to parse skin file %s: %v", filePath, err)
	}
	return skin
}
