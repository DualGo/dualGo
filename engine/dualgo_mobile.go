//+build android

package engine

import (
	"golang.org/x/mobile/app"
	"golang.org/x/mobile/gl"
	gl2 "github.com/DualGo/gl/all-core/gl"
)
s
var(
	Gl *gl.Context
)

func CreateWindow(title string, width, height int, fullscreen bool, msaa int){
	Gl = gl2.NewContext()


	Update(0)

}

func Update(delta float32){
	app.Main(func(a app.App) {

	})
}

func DeleteWindow(){
}
