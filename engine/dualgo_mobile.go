//+build android

package engine

import (
	"golang.org/x/mobile/app"
	"golang.org/x/mobile/gl"
)

var(
	Gl *gl.Context
)

func CreateWindow(title string, width, height int, fullscreen bool, msaa int){
	Gl = gl.NewContext()

	initEngine()
	Update(0)

}

func Update(delta float32){
	app.Main(func(a app.App) {
		loopEngine()
	})
}

func DeleteWindow(){
}
