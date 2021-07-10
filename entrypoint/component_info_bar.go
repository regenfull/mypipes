package entrypoint

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

const (
	ComponentInfoBarHeight = 3
)

type ComponentInfoBar struct {
	element *widgets.Paragraph
}

func (c *ComponentInfoBar) Init() error {
	c.element = widgets.NewParagraph()
	c.element.Text = "N - create | E - edit | D - create directory | R - remove"
	c.element.TextStyle = ui.NewStyle(ui.ColorYellow)
	width, height := ui.TerminalDimensions()
	c.element.SetRect(0, height-ComponentInfoBarHeight, width, height)
	return nil
}

func (c *ComponentInfoBar) Update() {
	width, height := ui.TerminalDimensions()
	c.element.SetRect(0, height-ComponentInfoBarHeight, width, height)
}

func (c ComponentInfoBar) GetUiElement() ui.Drawable {
	return c.element
}
