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
