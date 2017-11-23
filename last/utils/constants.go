package constants

import "github.com/DualGo/mathgl/mgl32"

type game struct {
	Width, Height int
	Title         string
	TexturePath   string
	Debug         bool
	DefaultColor  mgl32.Vec4
}

var Param game = game{0, 0, "Title", "", false, mgl32.Vec4{1, 1, 1, 1}}
