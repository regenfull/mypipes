package domain

type IStorage interface {
	Save(c *Command) (uint, error)
	Get(id uint) (*Command, error)
	GetAll() ([]Command, error)
	Delete(id uint) error
}
