package usecase

import (
	"log"

	"github.com/mega8bit/mypipes/domain"
)

type ControlUseCase struct {
	storage domain.IStorage

	l *log.Logger
}

func (c ControlUseCase) LoadAll() ([]domain.Command, error) {
	result, err := c.storage.GetAll()
	if err != nil {
		c.l.Println("Cannot select all saved commands: ", err.Error())
	}

	return result, err
}

func NewControlUseCase(storage domain.IStorage, logger *log.Logger) domain.IControlUseCase {
	return &ControlUseCase{
		storage: storage,
		l:       logger,
	}
}
