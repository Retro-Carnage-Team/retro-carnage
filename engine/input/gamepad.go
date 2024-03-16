package input

import (
	"math"
	"strings"

	"github.com/faiface/pixel/pixelgl"
)

// gamepad can be used to access the device state of a gamepad or joystick.
// Both digital and analog sticks are supported. The class uses the pixel framework / OpenGL to get required hardware
// information.
type gamepad struct {
	analog   *bool
	joystick pixelgl.Joystick
	name     string
	window   *pixelgl.Window
}

const (
	// Old analog controllers might be a bit wobbly and need a higher value.
	inputThreshold = 0.05
)

var (
	digitalControllers = []string{"Competition Pro"}

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
func (g *gamepad) State() *DeviceState {
	var state DeviceState
	var horizontal = g.window.JoystickAxis(g.joystick, pixelgl.AxisLeftX)
	var vertical = g.window.JoystickAxis(g.joystick, pixelgl.AxisLeftY)
	if g.isAnalog() {
		// Checked this with XBox360 and PlayStation controllers
		state.Fire = g.window.JoystickPressed(g.joystick, pixelgl.ButtonA)
		state.Grenade = g.window.JoystickPressed(g.joystick, pixelgl.ButtonB)
		state.ToggleUp = g.window.JoystickPressed(g.joystick, pixelgl.ButtonX)
		state.ToggleDown = g.window.JoystickPressed(g.joystick, pixelgl.ButtonY)
		if g.isStickMovedFully(horizontal, vertical) {
			var angle = g.computeStickAngle(horizontal, vertical*-1)
			state.MoveUp, state.MoveDown, state.MoveLeft, state.MoveRight = g.convertStickAngleToCardinalDirections(angle)
		}
	} else {
		// Checked this with a SpeedLink Competition Pro USB
		state.Fire = g.window.JoystickPressed(g.joystick, pixelgl.ButtonTriangle)
		state.Grenade = g.window.JoystickPressed(g.joystick, pixelgl.ButtonCross)
		state.ToggleDown = g.window.JoystickPressed(g.joystick, pixelgl.ButtonLeftBumper)
		state.ToggleUp = g.window.JoystickPressed(g.joystick, pixelgl.ButtonCircle)
		state.MoveUp = vertical == -1
		state.MoveDown = vertical == 1
		state.MoveLeft = horizontal == -1
		state.MoveRight = horizontal == 1
	}
	return &state
}

// Name returns the human readable name of the gamepad.
func (g *gamepad) Name() string {
	if g.name != "" {
		return g.name
	}

	var name = g.window.JoystickName(g.joystick)
	g.name = name
	return name
}

// isAnalog returns true when the controller axis allows values other than [-1, 0, 1].
func (g *gamepad) isAnalog() bool {
	if nil != g.analog {
		return *g.analog
	}

	var analogController = true
	for _, controllerName := range digitalControllers {
		if strings.Contains(strings.ToLower(g.Name()), strings.ToLower(controllerName)) {
			analogController = false
			break
		}
	}
	g.analog = &analogController
	return analogController
}
