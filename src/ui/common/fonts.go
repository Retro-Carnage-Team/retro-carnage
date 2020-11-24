package common

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"image/color"
	"io/ioutil"
	"os"
	"retro-carnage/engine/geometry"
	"retro-carnage/logging"
	"unicode"
)

const defaultFontPath = "./fonts/XXII-DIRTY-ARMY.ttf"

var DefaultAtlas *text.Atlas

func InitializeFonts() {
	defaultFont, err := loadTTF(defaultFontPath, 52)
	if nil != err {
		logging.Error.Panicf("Failed to load font %s: %v", defaultFontPath, err)
	}
	DefaultAtlas = text.NewAtlas(defaultFont, text.ASCII, text.RangeTable(unicode.Latin))
}

func loadTTF(path string, size float64) (font.Face, error) {
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

	return truetype.NewFace(font, &truetype.Options{
		Size:              size,
		GlyphCacheEntries: 1,
	}), nil
}

func GetTextDimensions(txt *text.Text, input ...string) map[string]*geometry.Point {
	var result = make(map[string]*geometry.Point)
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

	var txt = text.New(pixel.V(lineX, lineY), DefaultAtlas)
	txt.Color = color
	_, _ = fmt.Fprint(txt, line)
	txt.Draw(window, pixel.IM)
}