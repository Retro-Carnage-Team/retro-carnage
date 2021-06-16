package input

import (
	"github.com/faiface/pixel/pixelgl"
	"math"
)

type gamepad struct {
	joystick pixelgl.Joystick
	name     string
	window   *pixelgl.Window
}

// Old controllers might be a bit wobbly and need a higher value.
const inputThreshold = 0.15

var (
	PiOver8        = math.Pi / 8
	PiTimes3Over8  = (3 * math.Pi) / 8
	PiTimes5Over8  = (5 * math.Pi) / 8
	PiTimes7Over8  = (7 * math.Pi) / 8
	PiTimes9Over8  = (9 * math.Pi) / 8
	PiTimes11Over8 = (11 * math.Pi) / 8
	PiTimes13Over8 = (13 * math.Pi) / 8
	PiTimes15Over8 = (15 * math.Pi) / 8
)

func isStickMovedFully(x float64, y float64) bool {
	var radius = math.Sqrt(x*x + y*y) // Use Pythagorean theorem
	return 1-inputThreshold < radius
}

// Computes the angle (given in radians) for any point of the unit circle.
func computeStickAngle(x float64, y float64) float64 {
	if 0 <= x && 0 <= y {
		return math.Asin(y)
	}
	if 0 > x && 0 <= y {
		return math.Pi - math.Asin(y)
	}
	if 0 > x && 0 > y {
		return math.Pi + math.Asin(-1*y)
	}
	return 2*math.Pi - math.Asin(-1*y)
}

// Converts the given angle (in radians) into a combination of 4 cardinal directions
func convertStickAngleToCardinalDirections(angle float64) (up, down, left, right bool) {
	if PiOver8 <= angle && PiTimes3Over8 > angle {
		return true, false, false, true
	}
	if PiTimes3Over8 <= angle && PiTimes5Over8 > angle {
		return true, false, false, false
	}
	if PiTimes5Over8 <= angle && PiTimes7Over8 > angle {
		return true, false, true, false
	}
	if PiTimes7Over8 <= angle && PiTimes9Over8 > angle {
		return false, false, true, false
	}
	if PiTimes9Over8 <= angle && PiTimes11Over8 > angle {
		return false, true, true, false
	}
	if PiTimes11Over8 <= angle && PiTimes13Over8 > angle {
		return false, true, false, false
	}
	if PiTimes13Over8 <= angle && PiTimes15Over8 > angle {
		return false, true, false, true
	}
	return false, false, false, true
}

func (g *gamepad) State() *DeviceState {
	var result DeviceState
	result.Fire = g.window.JoystickPressed(g.joystick, pixelgl.ButtonA)
	result.Grenade = g.window.JoystickPressed(g.joystick, pixelgl.ButtonB)
	var horizontal = g.window.JoystickAxis(g.joystick, pixelgl.AxisLeftX)
	var vertical = g.window.JoystickAxis(g.joystick, pixelgl.AxisLeftY)
	if isStickMovedFully(horizontal, vertical) {
		var angle = computeStickAngle(horizontal, vertical*-1)
		var up, down, left, right = convertStickAngleToCardinalDirections(angle)
		result.MoveUp = up
		result.MoveDown = down
		result.MoveLeft = left
		result.MoveRight = right
	}
	result.ToggleUp = g.window.JoystickPressed(g.joystick, pixelgl.ButtonX)
	result.ToggleDown = g.window.JoystickPressed(g.joystick, pixelgl.ButtonY)
	return &result
}

func (g *gamepad) Name() string {
	if "" != g.name {
		return g.name
	}

	var name = g.window.JoystickName(g.joystick)
	g.name = name
	return name
}
