package entrypoint

import (
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/mega8bit/mypipes/domain"
)

type CmdUi struct {
	infoBar   *ComponentInfoBar
	pipesTree *ComponentPipesTree
	pipeInfo  *ComponentPipeInfo

	commands []domain.Command
	crud     domain.ICrudUseCase
	control  domain.IControlUseCase

	l *log.Logger
}

func (c *CmdUi) Update(e ui.Event) {
	switch e.ID {
	case "j", "<Down>":
		c.pipesTree.ScrollDown()
	case "k", "<Up>":
		c.pipesTree.ScrollUp()
	case "n", "N":
		_, err := c.crud.Create()
		if err != nil {
			break
		}

		c.commands, err = c.control.LoadAll()
		if err != nil {
			break
		}

		c.pipesTree.SetCommands(c.commands)

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

func NewCmdUi(crud domain.ICrudUseCase, control domain.IControlUseCase, logger *log.Logger) (*CmdUi, error) {
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

	commands, err := control.LoadAll()
	if err != nil {
		return nil, err
	}

	pipesTree.SetCommands(commands)

	return &CmdUi{
		infoBar,
		pipesTree,
		pipeInfo,
		commands,
		crud,
		control,
		logger,
	}, nil
}
