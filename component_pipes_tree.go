package main

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type ComponentPipesTree struct {
	element      *widgets.Tree
	nodes        []*widgets.TreeNode
	lastPipesCnt int
}

func (c *ComponentPipesTree) Init() error {
	for _, p := range pipes {
		c.nodes = append(c.nodes, &widgets.TreeNode{
			Value: WidgetTreeNodeValue(p.Name),
		})
	}
	c.lastPipesCnt = len(pipes)

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

	if c.lastPipesCnt != len(pipes) {
		c.nodes = c.nodes[:0]
		for _, p := range pipes {
			c.nodes = append(c.nodes, &widgets.TreeNode{
				Value: WidgetTreeNodeValue(p.Name),
			})
		}
		c.lastPipesCnt = len(pipes)
		c.element.SetNodes(c.nodes)
	}

	switch e.ID {
	case "j", "<Down>":
		c.element.ScrollDown()
	case "k", "<Up>":
		c.element.ScrollUp()
	}
}

func (c ComponentPipesTree) GetSelectedTree() int {
	return c.element.SelectedRow
}

func (c ComponentPipesTree) GetUiElement() ui.Drawable {
	return c.element
}
