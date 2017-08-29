package input

// # Package input
import "github.com/DualGo/glfw/v3.2/glfw"

//- ## Const 
//	- ##### KEYDOWN = 1 `int`	*state of key*
//	- ##### KEYUP = 0 `int`	*state of key*
//	- ##### KEYPRESS = 2 `int`	*state of key*
const (
	KEYDOWN  = 1
	KEYUP    = 0
	KEYPRESS = 2
)

//- ## Var 
//	- #### keys type `[]Key` *store keys of the input system* 
//	- #### joysticks type `[]int` 
var keys []Key
var joysticks []Joystick 


//- ## Struct Key 
type Key struct {
	keycode glfw.Key
	action  glfw.Action
}

//	- ### GetKey(keycode `gflw.key`)
//		- > return the state of a key 
// 
//		- > return `int`
// 
func GetKey(keycode glfw.Key) int {
	for _, element := range keys {
		if keycode == element.keycode {
			return int(element.action)
		}
	}
	return 0
}

//	- ### SetKey(keycode `gflw.Key`, action `gflw.Action`)
//		- > addKey to input system
// 
//		- > return `void`
// 
func SetKey(keycode glfw.Key, action glfw.Action) {
	keys = append(keys, Key{keycode, action})
}

//	- ### RemoveKeys()
//		- > clear the input system 
// 
//		- > return `void`
// 
func RemoveKeys() {
	keys = nil
}

//- ## Struct Joystick
type Joystick struct{
	value	int

}

func GetAxe(joy, axe int) float32{
	if isJoyConnected(joy){
		return glfw.GetJoystickAxes(glfw.Joystick(joy))[axe]
	}
	return 0
}

//	- ### ConnectJoystick(joy `int`)
//		- > add joystick to joysticks list 
// 		
//		- > return `void`
//
func ConnectJoystick(joy int){
	find := false
	for _, element := range joysticks{
		if element.value == joy{
			find  = true
			break
		}
	}
	if !find{
		joysticks = append(joysticks, Joystick{value : joy})
	}
}

//	- ### DisconnectJoystick(joy `int`)
//		- > remove joystick from the joysticks list 
// 
//		- > return `void`
// 
func DisconnectJoystick(joy int){
	var tempJoystick []Joystick
	for _, element := range joysticks{
		if element.value != joy{
			tempJoystick = append(tempJoystick, element)
			continue
		}
	}
	joysticks = tempJoystick
}

//	- ### isJoyConnected(joy `int`)
//		- > check if a joystick is connected
// 
//		- > return `bool`
//
func isJoyConnected(joy int) bool{
	for _, element := range joysticks{
		if element.value == joy{
			return true
		}
	}
	return false
}