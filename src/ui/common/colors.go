package common

import (
	"fmt"
	"image/color"
	"retro-carnage/logging"
)

var (
	Black  = parseHexColor("#000000")
	Green  = parseHexColor("#12a45d")
	Red    = parseHexColor("#ca5512")
	White  = parseHexColor("#ffffff")
	Yellow = parseHexColor("#fbe356")
)

func parseHexColor(s string) (c color.RGBA) {
	c.A = 0xff
	_, err := fmt.Sscanf(s, "#%02x%02x%02x", &c.R, &c.G, &c.B)
	if nil != err {
		logging.Error.Panicf("Failed to parse Hex formatted color: %v", err)
	}
	return
}
