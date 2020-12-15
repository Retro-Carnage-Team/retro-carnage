package common

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/golang/freetype/truetype"
	"image/color"
	"io/ioutil"
	"os"
	"retro-carnage/engine/geometry"
	"retro-carnage/logging"
	"unicode"
)

const defaultFontPath = "./fonts/XXII-DIRTY-ARMY.ttf"
const DefaultFontSize = 52

var SizeToFontAtlas map[int]*text.Atlas

func InitializeFonts() {
	SizeToFontAtlas = make(map[int]*text.Atlas)

	defaultFont, err := loadTTF(defaultFontPath)
	if nil != err {
		logging.Error.Panicf("Failed to load font %s: %v", defaultFontPath, err)
	}

	for i := 16; i <= DefaultFontSize; i += 2 {
		var fontFace = truetype.NewFace(defaultFont, &truetype.Options{
			Size:              float64(i),
			GlyphCacheEntries: 1,
		})
		SizeToFontAtlas[i] = text.NewAtlas(fontFace, text.ASCII, text.RangeTable(unicode.Latin))
	}
}

func loadTTF(path string) (*truetype.Font, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	font, err := truetype.Parse(bytes)
	if err != nil {
		return nil, err
	}

	return font, nil
}

func GetTextDimensions(fontSize int, input ...string) map[string]*geometry.Point {
	var result = make(map[string]*geometry.Point)
	var txt = text.New(pixel.V(0, 0), SizeToFontAtlas[fontSize])
	for _, line := range input {
		_, _ = fmt.Fprint(txt, line)
		result[line] = &geometry.Point{X: txt.Dot.X, Y: txt.LineHeight}
		txt.Clear()
	}
	return result
}

func DrawLineToScreenCenter(window *pixelgl.Window, line string, offsetMultiplier float64, color color.Color,
	lineDimensions *geometry.Point) {

	var vertCenter = window.Bounds().Max.Y / 2
	var lineX = (window.Bounds().Max.X - lineDimensions.X) / 2
	var lineY = vertCenter + offsetMultiplier*lineDimensions.Y

	var txt = text.New(pixel.V(lineX, lineY), SizeToFontAtlas[DefaultFontSize])
	txt.Color = color
	_, _ = fmt.Fprint(txt, line)
	txt.Draw(window, pixel.IM)
}
