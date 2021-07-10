package entrypoint

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type ComponentPipeInfo struct {
	element  *widgets.Paragraph
	pipeTree *ComponentPipesTree
}

func (c *ComponentPipeInfo) Init() error {
	c.element = widgets.NewParagraph()
	c.element.TextStyle = ui.NewStyle(ui.ColorYellow)

	width, height := ui.TerminalDimensions()
	c.element.SetRect(width/4, 0, width, height-ComponentInfoBarHeight)
	return nil
}

func (c *ComponentPipeInfo) SetPipesTree(p *ComponentPipesTree) {
	c.pipeTree = p
}

func (c *ComponentPipeInfo) Update() {
	width, height := ui.TerminalDimensions()
	c.element.SetRect(width/4, 0, width, height-ComponentInfoBarHeight)
}

func (c ComponentPipeInfo) GetUiElement() ui.Drawable {
	return c.element
}
