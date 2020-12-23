package fonts

type TextLayout struct {
	fontSize int
	lines    []TextLine
}

func (tl *TextLayout) FontSize() int {
	return tl.fontSize
}

func (tl *TextLayout) Lines() []TextLine {
	return tl.lines
}

func (tl *TextLayout) Height() float64 {
	var result = 0.0
	if 1 <= len(tl.lines) {
		result += tl.lines[0].Dimension().Y
	}
	if 2 <= len(tl.lines) {
		result += (float64(len(tl.lines)) - 1) * tl.lines[0].Dimension().Y * 1.2
	}
	return result
}
