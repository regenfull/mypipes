package models

import (
	"github.com/stretchr/testify/mock"
)

type ExternalEditorMock struct {
	mock.Mock
}

func (e *ExternalEditorMock) Launch(text string) (string, error) {
	args := e.Called(text)
	return args.String(0), args.Error(1)
}
