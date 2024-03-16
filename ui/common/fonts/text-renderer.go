package fonts

import (
	"errors"
	"fmt"
	"image/color"
	"retro-carnage/engine/geometry"
	"strings"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
)

const err_msg_text_doesnt_fit = "text doesn't fit in output region"

// TextRenderer can be used to render text to the screen
type TextRenderer struct {
	Window *pixelgl.Window
}

// DrawLineToScreenCenter renders a given line of text to the center of the screen. The horizontal position is centered.
// The vertical position of the line can be modified with the offsetMultiplier parameter. Positive offsetMultiplier
// values move the line up, negative values down.
func (tr *TextRenderer) DrawLineToScreenCenter(line string, offsetMultiplier float64, color color.Color) {
	var defaultFontSize = DefaultFontSize()
	var lineDimensions = GetTextDimension(defaultFontSize, line)

	var vertCenter = tr.Window.Bounds().Max.Y / 2
	var lineX = (tr.Window.Bounds().Max.X - lineDimensions.X) / 2
	var lineY = vertCenter + offsetMultiplier*lineDimensions.Y

	var txt = text.New(pixel.V(lineX, lineY), SizeToFontAtlas[defaultFontSize])
	txt.Color = color
	_, _ = fmt.Fprint(txt, line)
	txt.Draw(tr.Window, pixel.IM)
}

// CalculateTextLayout calculates the layout of a given text when rendered into an area of a specific size. It handles
// line breaks, removes additional whitespace and so on. An error is returned if the given text doesn't fit in the
// specified area when using the specified font size.
func (tr *TextRenderer) CalculateTextLayout(text string, fontSize int, width int, height int) (*TextLayout, error) {
	var words = strings.Fields(text)
	var result = &TextLayout{
		fontSize: fontSize,
		lines:    make([]TextLine, 0),
	}

	var lineText = ""
	var lineNumber = 0
	for _, word := range words {
		var newLineText = lineText
		if newLineText != "" {
			newLineText += " "
		}
		newLineText += word

		var lineDimensions = GetTextDimension(fontSize, newLineText)
		if lineDimensions.X > float64(width) {
			if lineText == "" {
				return nil, errors.New(err_msg_text_doesnt_fit)
			}
			var positionY = float64(height) - lineDimensions.Y - (float64(lineNumber) * lineDimensions.Y * 1.2)
			if float64(height) < positionY {
				return nil, errors.New(err_msg_text_doesnt_fit)
			}
			result.lines = append(result.lines, TextLine{
				dimension: lineDimensions,
				position:  &geometry.Point{X: 0, Y: positionY},
				text:      lineText,
			})
			lineText = word
			lineNumber++
		} else {
			lineText = newLineText
		}
	}

	if lineText != "" {
		var lineDimensions = GetTextDimension(fontSize, lineText)
		var positionY = float64(height) - lineDimensions.Y - (float64(lineNumber) * lineDimensions.Y * 1.2)
		if float64(height) < positionY {
			return nil, errors.New(err_msg_text_doesnt_fit)
		}
		result.lines = append(result.lines, TextLine{
			dimension: lineDimensions,
			position:  &geometry.Point{X: 0, Y: positionY},
			text:      lineText,
		})
	}

	return result, nil
}

func (tr *TextRenderer) RenderTextLayout(textLayout *TextLayout, fontSize int, color color.Color, position *geometry.Point) {
	var atlas = SizeToFontAtlas[fontSize]
	for _, line := range textLayout.Lines() {
		var txt = text.New(position.Add(line.position).ToVec(), atlas)
		txt.Color = color
		_, _ = fmt.Fprint(txt, line.text)
		txt.Draw(tr.Window, pixel.IM)
	}
}
