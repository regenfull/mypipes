package usecase

import (
	"bytes"
	"log"
	"testing"

	"github.com/mega8bit/mypipes/domain"
	"github.com/mega8bit/mypipes/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate_Ok(t *testing.T) {
	var fBuf = new(bytes.Buffer)

	mockStorage := new(models.StorageSqlLiteMock)
	mockStorage.
		On("Save", &domain.Command{
			Name: "firstcommand",
			Cmd:  "grep 1",
		}).
		Return(1, nil)

	mockEditor := new(models.ExternalEditorMock)
	mockEditor.
		On("Launch", "Name=\nCommand=\n").
		Return("Name=firstcommand\nCommand=grep 1", nil)

	usecase := CrudUseCase{
		storage: mockStorage,
		editor:  mockEditor,
		l:       log.New(fBuf, "", 0),
	}

	result, _ := usecase.Create()

	assert.Equal(t, result.Cmd, "grep 1")
	assert.Equal(t, result.Name, "firstcommand")
	assert.Equal(t, result.Id, uint(1))

	mockStorage.AssertCalled(t, "Save", mock.Anything)
	mockEditor.AssertCalled(t, "Launch", mock.Anything)
}

func TestCreate_WrongEditorData(t *testing.T) {
	var fBuf = new(bytes.Buffer)

	mockStorage := new(models.StorageSqlLiteMock)

	mockEditor := new(models.ExternalEditorMock)
	mockEditor.
		On("Launch", "Name=\nCommand=\n").
		Return("", nil)

	usecase := CrudUseCase{
		storage: mockStorage,
		editor:  mockEditor,
		l:       log.New(fBuf, "", 0),
	}

	_, e := usecase.Create()

	assert.NotEqual(t, e, nil)
	mockStorage.AssertNotCalled(t, "Save", mock.Anything)
	mockEditor.AssertCalled(t, "Launch", mock.Anything)
}
