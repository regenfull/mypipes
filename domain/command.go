package domain

type Command struct {
	Id   uint
	Name string
	Cmd  string
}

type IControlUseCase interface {
	LoadAll() ([]Command, error)
}
