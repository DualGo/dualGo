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
var keys []Key


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
