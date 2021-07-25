package models

import (
	"github.com/mega8bit/mypipes/domain"
	"github.com/stretchr/testify/mock"
)

type StorageSqlLiteMock struct {
	mock.Mock
}

func (s *StorageSqlLiteMock) Save(c *domain.Command) (uint, error) {
	args := s.Called(c)
	return uint(args.Int(0)), args.Error(1)
}

func (s *StorageSqlLiteMock) Get(id uint) (*domain.Command, error) {
	args := s.Called(id)
	return args.Get(0).(*domain.Command), args.Error(1)
}

func (s *StorageSqlLiteMock) GetAll() ([]domain.Command, error) {
	args := s.Called()
	return args.Get(0).([]domain.Command), args.Error(1)
}

func (s *StorageSqlLiteMock) Delete(id uint) error {
	args := s.Called(id)
	return args.Error(0)
}
