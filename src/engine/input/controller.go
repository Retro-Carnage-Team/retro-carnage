package input

import "github.com/faiface/pixel/pixelgl"

type Source interface {
	State() *State
	Name() string
}

type Controller struct {
	ControllerPlayerOne Source
	ControllerPlayerTwo Source
	Window              *pixelgl.Window
}

var joysticks = []pixelgl.Joystick{pixelgl.Joystick1, pixelgl.Joystick2, pixelgl.Joystick3, pixelgl.Joystick4,
	pixelgl.Joystick5, pixelgl.Joystick6, pixelgl.Joystick7, pixelgl.Joystick8, pixelgl.Joystick9, pixelgl.Joystick10,
	pixelgl.Joystick11, pixelgl.Joystick12, pixelgl.Joystick13, pixelgl.Joystick14, pixelgl.Joystick15,
	pixelgl.Joystick16}

func (c *Controller) HasTwoOrMoreDevices() bool {
	for _, j := range joysticks {
		if c.Window.JoystickPresent(j) {
			return true
		}
	}
	return false
}

func (c *Controller) AssignControllersToPlayers() {
	for _, j := range joysticks {
		if c.Window.JoystickPresent(j) {
			if nil == c.ControllerPlayerOne {
				c.ControllerPlayerOne = &gamepad{joystick: j, window: c.Window}
			} else if nil == c.ControllerPlayerTwo {
				c.ControllerPlayerTwo = &gamepad{joystick: j, window: c.Window}
				return
			}
		}
	}
	if nil == c.ControllerPlayerOne {
		c.ControllerPlayerOne = &keyboard{Window: c.Window}
	} else if nil == c.ControllerPlayerTwo {
		c.ControllerPlayerTwo = &keyboard{Window: c.Window}
	}
}
