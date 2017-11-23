package input

import (
	"github.com/DualGo/glfw/v3.2/glfw"
	"github.com/DualGo/mathgl/mgl64"
)

const (
	KEYUP    = 0
	KEYDOWN  = 1
	KEYPRESS = 2
)

const (
	BUTTONUP    = 0
	BUTTONDOWN  = 1
	BUTTONPRESS = 2
)

var keys []Key
var mouseButtons []Button
var mousePosition mgl64.Vec2
var joysticks []Joystick

type Key struct {
	keycode glfw.Key
	action  glfw.Action
}

type Button struct {
	buttoncode glfw.MouseButton
	action     glfw.Action
}

func GetKey(keycode glfw.Key) int {
	for _, element := range keys {
		if keycode == element.keycode {
			return int(element.action)
		}
	}
	return 0
}

func SetKey(keycode glfw.Key, action glfw.Action) {
	keys = append(keys, Key{keycode, action})
}

func RemoveKeys() {
	keys = nil
}

func GetButton(buttoncode glfw.MouseButton) int {
	for _, element := range mouseButtons {
		if buttoncode == element.buttoncode {
			return int(element.action)
		}
	}
	return 0
}

func SetButton(buttoncode glfw.MouseButton, action glfw.Action) {
	mouseButtons = append(mouseButtons, Button{buttoncode, action})
}

func RemoveButtons() {
	mouseButtons = nil
}

func SetMousePosition(pos mgl64.Vec2) {
	mousePosition = pos
}

func GetMousePosition() mgl64.Vec2 {
	return mousePosition
}

type Joystick struct {
	value int
}

func GetAxe(joy, axe int) float32 {
	if isJoyConnected(joy) {
		return glfw.GetJoystickAxes(glfw.Joystick(joy))[axe]
	}
	return 0
}

func ConnectJoystick(joy int) {
	find := false
	for _, element := range joysticks {
		if element.value == joy {
			find = true
			break
		}
	}
	if !find {
		joysticks = append(joysticks, Joystick{value: joy})
	}
}

func DisconnectJoystick(joy int) {
	var tempJoystick []Joystick
	for _, element := range joysticks {
		if element.value != joy {
			tempJoystick = append(tempJoystick, element)
			continue
		}
	}
	joysticks = tempJoystick
}

func isJoyConnected(joy int) bool {
	for _, element := range joysticks {
		if element.value == joy {
			return true
		}
	}
	return false
}
