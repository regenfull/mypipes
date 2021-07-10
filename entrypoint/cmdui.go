package entrypoint

import (
	ui "github.com/gizak/termui/v3"
)

type CmdUi struct {
	infoBar   *ComponentInfoBar
	pipesTree *ComponentPipesTree
	pipeInfo  *ComponentPipeInfo
}

func (c CmdUi) Update(e ui.Event) {
	switch e.ID {
	case "j", "<Down>":
		c.pipesTree.ScrollDown()
	case "k", "<Up>":
		c.pipesTree.ScrollUp()
	case "n", "N":
	case "d", "D":
	}
}

func (c CmdUi) GetComponents() []IComponent {
	return []IComponent{
		c.infoBar,
		c.pipesTree,
		c.pipeInfo,
	}
}

func NewCmdUi() (*CmdUi, error) {
	var err error
	var infoBar = new(ComponentInfoBar)
	var pipesTree = new(ComponentPipesTree)
	var pipeInfo = new(ComponentPipeInfo)

	if err = infoBar.Init(); err != nil {
		return nil, err
	}

	if err = pipesTree.Init(); err != nil {
		return nil, err
	}

	if err = pipeInfo.Init(); err != nil {
		return nil, err
	}

	return &CmdUi{
		infoBar,
		pipesTree,
		pipeInfo,
	}, nil
}
