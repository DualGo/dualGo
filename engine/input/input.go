package input

// # Package input
import "github.com/DualGo/glfw/v3.2/glfw"

//- ## Struct Key 
type Key struct {
	keycode glfw.Key
	action  glfw.Action
}

const (
	KEYDOWN  = 1
	KEYUP    = 0
	KEYPRESS = 2
)

var keys []Key

//	- ### GetKey(keycode gflw.key)
//		- > return the state of a key 
// 
//		- > return int
// 
func GetKey(keycode glfw.Key) int {
	for _, element := range keys {
		if keycode == element.keycode {
			return int(element.action)
		}
	}
	return 0
}

//	- ### SetKey(keycode gflw.Key, action gflw.Action)
//		- > addKey to input system
// 
//		- > return void
// 
func SetKey(keycode glfw.Key, action glfw.Action) {
	keys = append(keys, Key{keycode, action})
}

//	- ### RemoveKeys()
//		- > clear the input system 
// 
//		- > return void
// 
func RemoveKeys() {
	keys = nil
}
