package main

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type ComponentPipeInfo struct {
	element *widgets.Paragraph
}

func (c *ComponentPipeInfo) Init() error {
	c.element = widgets.NewParagraph()
	c.element.TextStyle = ui.NewStyle(ui.ColorYellow)

	width, height := ui.TerminalDimensions()
	c.element.SetRect(width/4, 0, width, height-ComponentInfoBarHeight)
	return nil
}

func (c *ComponentPipeInfo) Update(e ui.Event) {
	width, height := ui.TerminalDimensions()
	c.element.SetRect(width/4, 0, width, height-ComponentInfoBarHeight)
}

func (c ComponentPipeInfo) GetUiElement() ui.Drawable {
	return c.element
}
