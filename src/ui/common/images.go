package common

import (
	"github.com/faiface/pixel"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"retro-carnage/logging"
)

func LoadSprite(filePath string) *pixel.Sprite {
	var backgroundImage = loadPicture(filePath)
	return pixel.NewSprite(backgroundImage, backgroundImage.Bounds())
}

func loadPicture(filePath string) pixel.Picture {
	file, err := os.Open(filePath)
	if err != nil {
		logging.Error.Fatalf("Failed to load background image for title screen: %v", err)
		return nil
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		logging.Error.Fatalf("Failed to decode background image for title screen: %v", err)
		return nil
	}
	return pixel.PictureDataFromImage(img)
}
