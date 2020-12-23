package fonts

import (
	"fmt"
	"github.com/faiface/pixel"
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
var textDimensionCache map[string]*geometry.Point

func Initialize() {
	SizeToFontAtlas = make(map[int]*text.Atlas)
	textDimensionCache = make(map[string]*geometry.Point)

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
	for _, line := range input {
		result[line] = GetTextDimension(fontSize, line)
	}
	return result
}

func GetTextDimension(fontSize int, input string) *geometry.Point {
	var key = fmt.Sprintf("%d___%s", fontSize, input)
	var value = textDimensionCache[key]
	if nil == value {
		var txt = text.New(pixel.V(0, 0), SizeToFontAtlas[fontSize])
		_, _ = fmt.Fprint(txt, input)
		value = &geometry.Point{X: txt.Dot.X, Y: txt.LineHeight}
		textDimensionCache[key] = value
	}
	return value
}

func GetMaxTextWidth(fontSize int, input []string) float64 {
	var txt = text.New(pixel.V(0, 0), SizeToFontAtlas[fontSize])
	for _, line := range input {
		_, _ = fmt.Fprintln(txt, line)
	}
	return txt.Bounds().W()
}

func BuildText(position pixel.Vec, fontSize int, color color.Color, content string) *text.Text {
	var txt = text.New(position, SizeToFontAtlas[fontSize])
	txt.Color = color
	_, _ = fmt.Fprint(txt, content)
	return txt
}

func BuildMultiLineText(position pixel.Vec, fontSize int, color color.Color, content []string) *text.Text {
	var txt = text.New(position, SizeToFontAtlas[fontSize])
	txt.Color = color
	for _, line := range content {
		_, _ = fmt.Fprintln(txt, line)
	}
	return txt
}
