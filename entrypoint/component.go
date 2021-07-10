package entrypoint

import (
	ui "github.com/gizak/termui/v3"
)

type IComponent interface {
	Init() error
	GetUiElement() ui.Drawable
	Update()
}

type WidgetTreeNodeValue string

func (w WidgetTreeNodeValue) String() string {
	return string(w)
}
