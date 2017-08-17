package engine

import (
	"fmt"
	"log"

	"github.com/DualGo/dualGo/engine/graphics/d2d"

	"github.com/DualGo/dualGo/engine/extends"
	"github.com/DualGo/dualGo/engine/utils"

	"github.com/DualGo/gl/v4.1-core/gl"
	"github.com/DualGo/glfw/v3.2/glfw"
)

type Engine struct {
	width   int
	height  int
	title   string
	window  *glfw.Window
	objects []d2d.Drawable2D
	modules []extends.Module
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
	for _, element := range engine.modules {
		element.Init(engine.objects)
	}
	engine.loop(callbackUpdate)
}

/* render loop */
func (engine *Engine) loop(callback UpdateFunc) {
	for !engine.window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		for _, element := range engine.modules {
			if element.GetUpdatePosition() == "first" {
				element.Update()
			}

		}
		callback()
		for _, element := range engine.modules {
			if element.GetUpdatePosition() == "last" {
				element.Update()
			}
		}
		engine.window.SwapBuffers()
		glfw.PollEvents()
	}
}

func (engine *Engine) AddObject(object d2d.Drawable2D) {
	engine.objects = append(engine.objects, object)
}

func (engine *Engine) AddModule(module extends.Module) {
	engine.modules = append(engine.modules, module)
}
