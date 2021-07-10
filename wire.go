// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/mega8bit/mypipes/entrypoint"
)

func setup() (Application, error) {
	panic(wire.Build(
		NewApplication,
		entrypoint.NewCmdUi,
		//models.NewExternalEditor,
		//models.NewStorageSqlLite,
	))
}
