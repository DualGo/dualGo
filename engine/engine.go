package engine

import (
	"fmt"
	"log"
	"opengl/hello/engine/utils"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

type Engine struct {
	width  int
	height int
	title  string
	window *glfw.Window
}
type InitFunc func()
type UpdateFunc func()

/*Init all components require by opengl and the engine*/
func (engine *Engine) Init(width, height int, title string, callbackInit InitFunc, callbackUpdate UpdateFunc) {
	engine.width = width
	engine.height = height
	engine.title = title

	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	window, err := glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		panic(err)
	}
	engine.window = window
	engine.window.MakeContextCurrent()

	// Initialize Glow
	if err := gl.Init(); err != nil {
		panic(err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version", version)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	gl.ClearColor(0.0, 0.0, 0.0, 1.0)

	//setup constants
	constants.Param.Width = engine.width
	constants.Param.Height = engine.height
	callbackInit()
	engine.loop(callbackUpdate)
}

/* render loop */
func (engine *Engine) loop(callback UpdateFunc) {
	for !engine.window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		callback()
		engine.window.SwapBuffers()
		glfw.PollEvents()
	}
}
