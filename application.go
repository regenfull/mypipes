package main

import (
	ui "github.com/gizak/termui/v3"
	"github.com/mega8bit/mypipes/entrypoint"
)

type Application struct {
	epoint *entrypoint.CmdUi
}

func (a Application) Run() error {
	event := ui.PollEvents()
	for {
		ui.Clear()

		for _, c := range a.epoint.GetComponents() {
			ui.Render(c.GetUiElement())
		}

		e := <-event

		switch e.ID {
		case "<C-c>", "q":
			return nil
		}

		a.epoint.Update(e)

		for _, c := range a.epoint.GetComponents() {
			c.Update()
		}
	}
}

func NewApplication(app *entrypoint.CmdUi) Application {
	return Application{
		app,
	}
}
