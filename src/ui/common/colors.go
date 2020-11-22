package common

import (
	"fmt"
	"image/color"
	"retro-carnage/logging"
)

var (
	Red    = parseHexColor("#ca5512")
	Yellow = parseHexColor("#fbe356")
	Green  = parseHexColor("#12a45d")
)

func parseHexColor(s string) (c color.RGBA) {
	c.A = 0xff
	_, err := fmt.Sscanf(s, "#%02x%02x%02x", &c.R, &c.G, &c.B)
	if nil != err {
		logging.Error.Panicf("Failed to parse Hex formatted color: %v", err)
	}
	return
}
