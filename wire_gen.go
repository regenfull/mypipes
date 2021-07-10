// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/mega8bit/mypipes/entrypoint"
)

// Injectors from wire.go:

func setup() (Application, error) {
	cmdUi, err := entrypoint.NewCmdUi()
	if err != nil {
		return Application{}, err
	}
	application := NewApplication(cmdUi)
	return application, nil
}