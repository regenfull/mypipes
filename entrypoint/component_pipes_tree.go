package entrypoint

import (
	"fmt"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/mega8bit/mypipes/domain"
)

type ComponentPipesTree struct {
	element *widgets.Tree
	nodes   []*widgets.TreeNode
}

func (c *ComponentPipesTree) Init() error {
	c.element = widgets.NewTree()
	c.element.TextStyle = ui.NewStyle(ui.ColorYellow)

	width, height := ui.TerminalDimensions()
	c.element.SetRect(0, 0, width/3, height-ComponentInfoBarHeight)
	return nil
}

func (c *ComponentPipesTree) Update() {
	width, height := ui.TerminalDimensions()
	c.element.SetRect(0, 0, width/3, height-ComponentInfoBarHeight)

}

func (c ComponentPipesTree) GetSelectedTree() int {
	return c.element.SelectedRow
}

func (c ComponentPipesTree) GetUiElement() ui.Drawable {
	return c.element
}

func (c ComponentPipesTree) ScrollUp() {
	c.element.ScrollUp()
}

func (c ComponentPipesTree) ScrollDown() {
	c.element.ScrollDown()
}

func (c *ComponentPipesTree) SetCommands(commands []domain.Command) {
	c.nodes = make([]*widgets.TreeNode, 0)
	for i, p := range commands {
		name := fmt.Sprintf("%d: %s", i+1, p.Name)
		c.nodes = append(c.nodes, &widgets.TreeNode{
			Value: WidgetTreeNodeValue(name),
		})
	}
	c.element.SetNodes(c.nodes)
}
