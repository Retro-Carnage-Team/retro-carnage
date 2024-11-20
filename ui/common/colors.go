package common

import (
	"fmt"
	"image/color"
	"retro-carnage/logging"
)

var (
	Black      = ParseHexColor("#000000")
	Gray       = ParseHexColor("#ababab")
	Green      = ParseHexColor("#12a45d")
	LightGray  = ParseHexColor("#d3d3d3")
	ModalBg    = ParseHexColor("#2f4f4e")
	OliveGreen = ParseHexColor("#897c2a")
	Orange     = ParseHexColor("#ffa055")
	Red        = ParseHexColor("#ca5512")
	White      = ParseHexColor("#ffffff")
	Yellow     = ParseHexColor("#fbe356")
)

func ParseHexColor(s string) (c color.RGBA) {
	c.A = 0xff
	_, err := fmt.Sscanf(s, "#%02x%02x%02x", &c.R, &c.G, &c.B)
	if nil != err {
		logging.Error.Panicf("Failed to parse Hex formatted color: %v", err)
	}
	return
}
