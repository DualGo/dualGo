package modules

// # Package modules
import (
	"github.com/DualGo/dualGo/engine/extends"
	"github.com/DualGo/dualGo/engine/graphics/d2d"
	"github.com/DualGo/dualGo/engine/renderer"
	"github.com/DualGo/glfw/v3.2/glfw"
	"github.com/DualGo/mathgl/mgl32"
)
//- ## Struct Debug
type Debug struct {
	extends.Module
	fps      int32
	lastFps  int32
	lastTime float64
	color    mgl32.Vec4
}

//	- ### Init(objects `[]d2d.Drawable2D`)
//		- > init the debug module 
// 
//		- > return `void` 
// 
func (debug *Debug) Init(objects []d2d.Drawable2D) {
	debug.fps = 16
	debug.lastFps = 16
	debug.color = mgl32.Vec4{0.4, 1, 0.4, 1}
	debug.lastTime = 0
}

//	- ### Update(renderer `*renderer.Renderer2D`)
//		- > update the  debug module 
// 
//		- > return `void`
// 
func (debug *Debug) Update(renderer *renderer.Renderer2D) {
	renderer.DrawText(10, 50, 0.5, "DEBUG MODE ACTIVATED", debug.color)
	currentTime := glfw.GetTime()
	debug.fps++
	if currentTime-debug.lastTime >= 1.0 {
		renderer.DrawText(10, 75, 0.5, String(1000/debug.fps), debug.color)
		debug.lastFps = debug.fps
		debug.fps = 0
		debug.lastTime = glfw.GetTime()
	} else {
		renderer.DrawText(10, 75, 0.5, String(1000/debug.lastFps), debug.color)
	}

}

//	- ### GetUpdatePosition()
//		- > return the udpate poisition ( first, midle, last ) 
// 
//		- > return `string`
// 
func (debuf *Debug) GetUpdatePosition() string {
	return "last"
}

//- ### String(n `int32`)
//	- > convert int32 into string 
// 
//	- > return `string`
// 
func String(n int32) string {
	buf := [11]byte{}
	pos := len(buf)
	i := int64(n)
	signed := i < 0
	if signed {
		i = -i
	}
	for {
		pos--
		buf[pos], i = '0'+byte(i%10), i/10
		if i == 0 {
			if signed {
				pos--
				buf[pos] = '-'
			}
			return string(buf[pos:])
		}
	}
}
