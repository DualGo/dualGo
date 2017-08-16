package modules

import (
	"github.com/DualGo/dualGo/engine/extends"
	"github.com/DualGo/dualGo/engine/graphics/d2d"
)

type Debug struct {
	extends.Module
}

func (debug Debug) Init(objects []d2d.Drawable2D) {

}

func (debug Debug) Update() {

}

func GetUpdatePosition() string {
	return "last"
}
