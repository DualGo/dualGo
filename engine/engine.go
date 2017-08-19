package engine

import (
	"fmt"
	"log"

	"github.com/DualGo/dualGo/engine/graphics/d2d"
	"github.com/DualGo/dualGo/engine/input"
	"github.com/DualGo/dualGo/engine/renderer"
	"github.com/DualGo/dualGo/engine/utils"

	"github.com/DualGo/dualGo/engine/extends"

	"github.com/DualGo/gl/v4.1-core/gl"
	"github.com/DualGo/glfw/v3.2/glfw"
)

type Engine struct {
	width    int
	height   int
	title    string
	window   *glfw.Window
	objects  []d2d.Drawable2D
	modules  []extends.Module
	renderer *renderer.Renderer2D
}
type InitFunc func()
type UpdateFunc func()

/*Init all components require by opengl and the engine*/
func (engine *Engine) Init(width, height int, renderer *renderer.Renderer2D, title string, callbackInit InitFunc, callbackUpdate UpdateFunc) {

	engine.renderer = renderer
	engine.width, engine.height = width, height

	engine.title = title

	//setup constants
	constants.Param.Width = engine.width
	constants.Param.Height = engine.height
	constants.Param.Title = engine.title

	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	window, err := glfw.CreateWindow(engine.width, engine.height, title, nil, nil)
	if err != nil {
		panic(err)
	}
	engine.window = window
	engine.window.MakeContextCurrent()
	glfw.SwapInterval(1)
	onResize := func(w *glfw.Window, width int, height int) {
		engine.renderer.SetWidth(width)
		engine.renderer.SetHeight(height)
	}
	onKeyPressed := func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		input.SetKey(key, action)
	}
	engine.window.SetSizeCallback(onResize)
	engine.window.SetKeyCallback(onKeyPressed)

	// Initialize Glow
	if err := gl.Init(); err != nil {
		panic(err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version", version)
	gl.ClearColor(0, 0, 0, 0.0)

	callbackInit()
	renderer.SetWidth(width)
	renderer.SetHeight(height)
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
		input.RemoveKeys()
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
