package extends

// # Package extends
import (
	"github.com/DualGo/dualGo/engine/graphics/d2d"
	"github.com/DualGo/dualGo/engine/renderer"
)

//- ## Interface Module
//- > Init(objects []d2d.Drawable2D)
//-
//- > Update(renderer *renderer.Renderer2D)
//-
//- > GetUpdatePosition() string
//-
type Module interface {
	Init(objects []d2d.Drawable2D)
	Update(renderer *renderer.Renderer2D)
	GetUpdatePosition() string
}
