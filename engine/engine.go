package engine

//# Package Engine
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

// - ## func InitFunc
// - ## func UpdateFunc
type InitFunc func()
type UpdateFunc func()

// - ## Struct Engine
type Engine struct {
	width    int
	height   int
	title    string
	window   *glfw.Window
	objects  []d2d.Drawable2D
	modules  []extends.Module
	renderer *renderer.Renderer2D
}

//	- ### Init(width, height int, renderer `*renderer.Renderer2D`)
//		- > init glfw and opengl
// 
//		- > return `void`
// 
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
	//glfw.SwapInterval(1)
	onResize := func(w *glfw.Window, width int, height int) {
		engine.renderer.SetWidth(width)
		engine.renderer.SetHeight(height)
	}
	onKeyPressed := func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		input.SetKey(key, action)
	}
	joystickCallback := func(joy, event int){
		if  event == 0x00040001 {
			input.ConnectJoystick(joy)
		}else if event == 0x00040002 {
			input.DisconnectJoystick(joy)
		} 
	}
	engine.window.SetSizeCallback(onResize)
	engine.window.SetKeyCallback(onKeyPressed)
	glfw.SetJoystickCallback(joystickCallback)

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

//	- ### AddObect(object `d2d.Drawable2D`)
//		- > add object to object list  
// 
//		- > return `void`
// 
func (engine *Engine) AddObject(object d2d.Drawable2D) {
	engine.objects = append(engine.objects, object)
}

//	- ### AddModule(module `extends.Module`)
//		- > add module to module list  
// 
//		- > return `void`
// 
func (engine *Engine) AddModule(module extends.Module) {
	engine.modules = append(engine.modules, module)
}

//	- ### loop(callback `UpdateFunc`)
//		- > main loop, call update function 
// 
//		- > return `void`
// 
func (engine *Engine) loop(callback UpdateFunc) {
	for !engine.window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		for _, element := range engine.modules {
			if element.GetUpdatePosition() == "first" {
				element.Update(engine.renderer)
			}

		}
		callback()
		for _, element := range engine.modules {
			if element.GetUpdatePosition() == "last" {
				element.Update(engine.renderer)
			}
		}
		input.RemoveKeys()
		engine.window.SwapBuffers()
		glfw.PollEvents()
	}
}