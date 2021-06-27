package main

import (
	ui "github.com/gizak/termui/v3"
)

type IComponent interface {
	Init() error
	GetUiElement() ui.Drawable
	Update(ui.Event)
}

type WidgetTreeNodeValue string
func (w WidgetTreeNodeValue) String() string {
	return string(w)
}