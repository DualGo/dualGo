package constants

// # Package constants

//- ## Struct game
//	- #### Width, Height `int`
// 	- #### Title `string`
//  - #### TexturePath `string`
type game struct {
	Width, Height int
	Title         string
	TexturePath   string
}

//- ## Var
//	- #### Param `game` *global game parameters*
var Param game = game{0, 0, "Title", ""}
