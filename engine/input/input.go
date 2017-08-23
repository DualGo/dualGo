package input

// # Package input
import "github.com/DualGo/glfw/v3.2/glfw"

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
