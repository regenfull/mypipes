package domain

type ICrudUseCase interface {
	Create() (*Command, error)
	Update(c *Command) error
	Delete(id uint) error
}
