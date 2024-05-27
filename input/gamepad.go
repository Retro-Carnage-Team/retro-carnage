package input

import (
	"math"
	"retro-carnage/config"

	pixel "github.com/Retro-Carnage-Team/pixel2"
	"github.com/Retro-Carnage-Team/pixel2/backends/opengl"
)

// gamepad can be used to access the device state of a gamepad or joystick.
// Both digital and analog sticks are supported. The class uses the pixel framework / OpenGL to get required hardware
// information.
type gamepad struct {
	configuration config.InputDeviceConfiguration
	window        *opengl.Window
}

const (
	// Old analog controllers might be a bit wobbly and need a higher value.
	inputThreshold = 0.05
	PiOver8        = math.Pi / 8
	PiTimes3Over8  = (3 * math.Pi) / 8
	PiTimes5Over8  = (5 * math.Pi) / 8
	PiTimes7Over8  = (7 * math.Pi) / 8
	PiTimes9Over8  = (9 * math.Pi) / 8
	PiTimes11Over8 = (11 * math.Pi) / 8
	PiTimes13Over8 = (13 * math.Pi) / 8
	PiTimes15Over8 = (15 * math.Pi) / 8
)

// isStickMovedFully checks the values of the X & Y axis of an analog stick on whether the stick has been moved fully
// to (any given) direction. This check has to be performed because the math used to determine the angle of the stick
// works only when it it moved (more or less) fully.
func (g *gamepad) isStickMovedFully(x float64, y float64) bool {
	var radius = math.Sqrt(x*x + y*y) // Use Pythagorean theorem
	return 1-inputThreshold < radius
}

// Computes the angle (given in radians) for any point of the unit circle.
func (g *gamepad) computeStickAngle(x float64, y float64) float64 {
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
func (g *gamepad) convertStickAngleToCardinalDirections(angle float64) (up, down, left, right bool) {
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

// State returns the DeviceState of the gamepad.
func (g *gamepad) State() *InputDeviceState {
	var joystick = pixel.Joystick(g.configuration.GamepadConfiguration.JoystickIndex)
	var horizontal = g.window.JoystickAxis(joystick, pixel.AxisLeftX)
	var vertical = g.window.JoystickAxis(joystick, pixel.AxisLeftY)

	var state = InputDeviceState{
		PrimaryAction: g.window.JoystickPressed(joystick, pixel.GamepadButton(g.configuration.InputFire)),
		ToggleDown:    g.window.JoystickPressed(joystick, pixel.GamepadButton(g.configuration.InputPreviousWeapon)),
		ToggleUp:      g.window.JoystickPressed(joystick, pixel.GamepadButton(g.configuration.InputNextWeapon)),
	}

	if g.configuration.GamepadConfiguration.HasDigitalAxis {
		state.MoveUp = vertical == -1
		state.MoveDown = vertical == 1
		state.MoveLeft = horizontal == -1
		state.MoveRight = horizontal == 1
	} else if g.isStickMovedFully(horizontal, vertical) {
		var angle = g.computeStickAngle(horizontal, vertical*-1)
		state.MoveUp, state.MoveDown, state.MoveLeft, state.MoveRight = g.convertStickAngleToCardinalDirections(angle)
	}

	return &state
}

// Name returns the human readable name of the gamepad.
func (g *gamepad) Name() string {
	return g.configuration.DeviceName
}
