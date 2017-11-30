package main

import (
	"github.com/DualGo/dualgo/engine"
)

var(
	renderer engine.Renderer
)

func main(){
	renderer.Init()
	engine.AddSystem(renderer.GetSystem())
	println("I'm init by glfw")
	engine.CreateWindow("test", 800, 600, false, 0)
}

