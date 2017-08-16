package extends

import "github.com/DualGo/dualGo/engine/graphics/d2d"

type Module interface {
	Init(objects []d2d.Drawable2D)
	Update()
	GetUpdatePosition() string
}
