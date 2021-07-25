package usecase

import (
	"errors"
	"log"
	"regexp"
	"strings"

	"github.com/mega8bit/mypipes/domain"
)

type CrudUseCase struct {
	storage domain.IStorage
	editor  domain.IExternalEditor

	l *log.Logger
}

func (c CrudUseCase) Create() (*domain.Command, error) {
	var err error
	cmdText, err := c.editor.Launch("Name=\nCommand=\n")
	if err != nil {
		c.l.Println("Cannot launch external editor with default text:", err.Error())
		return nil, err
	}

	var checkRegexp = `^Name=.+\nCommand=.+`

	if !regexp.MustCompile(checkRegexp).MatchString(cmdText) {
		return nil, errors.New("wrong raw data to parse")
	}

	var cmd = c.parseCmdFields(cmdText)

	if cmd.Name == "" || cmd.Cmd == "" {
		return nil, errors.New("empty fields in parsed cmd command from the editor")
	}

	id, err := c.storage.Save(cmd)
	if err != nil {
		c.l.Println("Storage has given a error", err.Error())
		return nil, err
	}

	cmd.Id = id
	return cmd, nil
}

func (c CrudUseCase) Update(cmd *domain.Command) error {
	return nil
}

func (c CrudUseCase) Delete(id uint) error {
	return nil
}

func (c CrudUseCase) parseCmdFields(cmdText string) *domain.Command {
	cmdTextParts := strings.Split(cmdText, "\n")
	cmdTextParts[0] = strings.ReplaceAll(cmdTextParts[0], "Name=", "")
	cmdTextParts[1] = strings.ReplaceAll(cmdTextParts[1], "Command=", "")

	cmdTextParts[0] = strings.TrimSpace(cmdTextParts[0])
	cmdTextParts[1] = strings.TrimSpace(cmdTextParts[1])

	return &domain.Command{
		Name: cmdTextParts[0],
		Cmd:  cmdTextParts[1],
	}
}

func NewCrudUseCase(
	storage domain.IStorage,
	editor domain.IExternalEditor,
	logger *log.Logger,
) domain.ICrudUseCase {
	return &CrudUseCase{
		storage,
		editor,
		logger,
	}
}
