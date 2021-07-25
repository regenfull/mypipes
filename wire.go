// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/mega8bit/mypipes/domain"
	"github.com/mega8bit/mypipes/entrypoint"
	"github.com/mega8bit/mypipes/models"
	"github.com/mega8bit/mypipes/usecase"
)

func setup() (Application, error) {
	panic(wire.Build(
		NewApplication,
		domain.NewLogger,
		entrypoint.NewCmdUi,
		models.NewExternalEditor,
		models.NewStorageSqlLite,
		usecase.NewCrudUseCase,
	))
}
