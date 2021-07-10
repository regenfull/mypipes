package models

import (
	"database/sql"

	"github.com/mega8bit/mypipes/domain"
)

type StorageSqlLite struct {
	db *sql.DB
}

func (StorageSqlLite) Save(c domain.Command) error {
	return nil
}

func (StorageSqlLite) Get(id uint) (domain.Command, error) {
	return domain.Command{}, nil
}

func (StorageSqlLite) GetAll() ([]domain.Command, error) {
	return nil, nil
}

func (StorageSqlLite) Delete(id uint) error {
	return nil
}

func NewStorageSqlLite() (*StorageSqlLite, error) {
	d, err := sql.Open("sqllite3", "~/.mypipes")
	if err != nil {
		return nil, err
	}

	return &StorageSqlLite{
		d,
	}, nil

}
