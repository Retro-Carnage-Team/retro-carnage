package util

import (
	"os"
	"strings"
)

func GetJsonFilesOfDirectory(directory string) ([]string, error) {
	if !strings.HasSuffix(directory, "/") {
		directory += "/"
	}

	files, err := os.ReadDir(directory)
	if err != nil {
		return nil, err
	}

	var result = make([]string, 0)
	for _, f := range files {
		if !f.IsDir() && strings.HasSuffix(strings.ToLower(f.Name()), ".json") {
			result = append(result, directory+f.Name())
		}
	}
	return result, nil
}
