package main

import (
	"bytes"
	"errors"
	"os"
)

type Storage struct {
	f *os.File
}

func (s *Storage) Init() error {
	var err error

	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	s.f, err = os.OpenFile(
		userHomeDir+"/.mypipes",
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		os.FileMode(0640),
	)

	if err != nil {
		return err
	}

	return nil
}

func (s Storage) Save(b bytes.Buffer) error {
	var err error

	if s.f == nil {
		return errors.New("pointer to the file of storage is not correct")
	}

	err = s.f.Truncate(0)
	if err != nil {
		return err
	}

	_, err = s.f.Write(b.Bytes())
	return err
}

func (s Storage) Read() (*bytes.Buffer, error) {
	var b = new(bytes.Buffer)
	var err error

	if s.f == nil {
		return nil, errors.New("pointer to the file of storage is not correct")
	}

	_, err = b.ReadFrom(s.f)

	return b, err
}
