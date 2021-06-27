package main

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type ComponentPipesTree struct {
	element *widgets.Tree
	nodes   []*widgets.TreeNode
}

func (c *ComponentPipesTree) Init() error {
	c.nodes = []*widgets.TreeNode{
		{
			Value: WidgetTreeNodeValue("test1"),
		},
		{
			Value: WidgetTreeNodeValue("test2"),
		},
		{
			Value: WidgetTreeNodeValue("test3"),
		},
	}

	c.element = widgets.NewTree()
	c.element.SetNodes(c.nodes)
	c.element.TextStyle = ui.NewStyle(ui.ColorYellow)

	width, height := ui.TerminalDimensions()
	c.element.SetRect(0, 0, width/4, height-ComponentInfoBarHeight)
	return nil
}

func (c *ComponentPipesTree) Update(e ui.Event) {
	width, height := ui.TerminalDimensions()
	c.element.SetRect(0, 0, width/4, height-ComponentInfoBarHeight)

	switch e.ID {
	case "j", "<Down>":
		c.element.ScrollDown()
	case "k", "<Up>":
		c.element.ScrollUp()
	}
}

func (c ComponentPipesTree) GetUiElement() ui.Drawable {
	return c.element
}
