package constants

// # Package constants

import "github.com/DualGo/mathgl/mgl32"

//- ## Struct game
//	- #### Width, Height `int`
// 	- #### Title `string`
//  - #### TexturePath `string`
type game struct {
	Width, Height int
	Title         string
	TexturePath   string
	Debug		  bool
	DefaultColor  mgl32.Vec4
}

//- ## Var
//	- #### Param `game` *global game parameters*
var Param game = game{0, 0, "Title", "",false, mgl32.Vec4{1,1,1,1}}
