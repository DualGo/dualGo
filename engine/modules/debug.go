package modules

//create by thewrath
//official debug module
//version 0.1

import (
	"math/rand"

	"github.com/DualGo/dualGo/engine/extends"
	"github.com/DualGo/dualGo/engine/graphics/d2d"
	"github.com/DualGo/dualGo/engine/renderer"
	"github.com/DualGo/dualGo/engine/utils"
	"github.com/DualGo/glfw/v3.2/glfw"
	"github.com/DualGo/mathgl/mgl32"
)

type Debug struct {
	extends.Module
	fps      int32
	lastFps  int32
	lastTime float64
	color    mgl32.Vec4
	shape    d2d.Rectangle
}

func (debug *Debug) Init(objects []d2d.Drawable2D) {
	debug.fps = 16
	debug.lastFps = 16
	debug.color = mgl32.Vec4{0.4, 1, 0.4, 1}
	debug.lastTime = 0
	debug.shape.Init(mgl32.Vec2{0, 0}, mgl32.Vec2{50, 50})
	debug.shape.SetStrokeColor(debug.color)
	debug.shape.SetColor(mgl32.Vec4{0, 0, 0, 0})
	debug.shape.SetStroke(0.005)

}

func (debug *Debug) Update(renderer *renderer.Renderer2D, objects []d2d.Drawable2D) {
	if constants.Param.Debug {
		for _, element := range objects {
			debug.shape.SetPosition(element.GetPosition())
			debug.shape.SetSize(element.GetSize())
			renderer.Draw(debug.shape)
		}
		renderer.DrawText(10, 50, 0.5, "DEBUG MODE ACTIVATED", debug.color)
		//joystick debuging
		renderer.DrawText(10, 75, 0.5, "MS : ", debug.color)
		currentTime := glfw.GetTime()
		debug.fps++
		if currentTime-debug.lastTime >= 1.0 {
			renderer.DrawText(70, 75, 0.5, String(1000/debug.fps), debug.color)
			debug.lastFps = debug.fps
			debug.fps = 0
			debug.lastTime = glfw.GetTime()
		} else {
			renderer.DrawText(70, 75, 0.5, String(1000/debug.lastFps), debug.color)
		}
	}

}

func (debuf *Debug) GetUpdatePosition() string {
	return "last"
}

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

func RandInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
