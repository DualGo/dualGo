//+build !netgo, !android

package engine

import(
	"github.com/DualGo/glfw/v3.2/glfw"

	"github.com/DualGo/glHelper"
)

var(
	window *glfw.Window
	Gl *glHelper.Context
)

func CreateWindow(title string, width, height int, fullscreen bool, msaa int){
	err := glfw.Init()

	RaisedError(err)

	window, err = glfw.CreateWindow(width, height, title, nil, nil)

	RaisedError(err)

	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window.MakeContextCurrent()

	Gl = glHelper.NewContext()
	Gl.ClearColor(1,0,0,1)
	loop(0)

}

func loop(delta float32){
	for !window.ShouldClose() {
		glfw.PollEvents()
		Gl.Clear(Gl.COLOR_BUFFER_BIT)
		loopEngine(delta)
		window.SwapBuffers()
	}
}

func DeleteWindow(){
	glfw.Terminate()
}