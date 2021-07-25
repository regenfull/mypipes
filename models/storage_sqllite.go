package models

import (
	"database/sql"
	"os"

	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mega8bit/mypipes/domain"
)

type StorageSqlLite struct {
	db *sql.DB
}

func (s *StorageSqlLite) Save(c *domain.Command) (uint, error) {
	var id uint
	err := s.db.QueryRow(`
		INSERT INTO commands(
			name,
			command
		) VALUES(?,?)
		RETURNING rowid
	`, c.Name, c.Cmd).Scan(&id)

	return id, err
}

func (StorageSqlLite) Get(id uint) (*domain.Command, error) {
	return &domain.Command{}, nil
}

func (StorageSqlLite) GetAll() ([]domain.Command, error) {
	return nil, nil
}

func (StorageSqlLite) Delete(id uint) error {
	return nil
}

func NewStorageSqlLite() (domain.IStorage, error) {
	path, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	path = path + "/.mypipes.db"

	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0640)
	if err != nil {
		return nil, err
	}
	f.Close()

	d, err := sql.Open("sqlite3", fmt.Sprintf("file:%s?mode=rw", path))
	if err != nil {
		return nil, err
	}

	_, err = d.Exec(`
		CREATE TABLE IF NOT EXISTS
			commands (
				name TEXT,
				command TEXT,
				create_datetime DATETIME DEFAULT CURRENT_TIMESTAMP
			)
		
	`)

	if err != nil {
		return nil, err
	}

	return &StorageSqlLite{
		d,
	}, nil

}
