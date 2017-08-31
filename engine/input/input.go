package input

import "github.com/DualGo/glfw/v3.2/glfw"

const (
	KEYDOWN  = 1
	KEYUP    = 0
	KEYPRESS = 2
)

var keys []Key
var joysticks []Joystick

type Key struct {
	keycode glfw.Key
	action  glfw.Action
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
