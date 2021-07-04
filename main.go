package main

import (
	ui "github.com/gizak/termui/v3"
)

func main() {
	defer ui.Close()

	event := ui.PollEvents()
	for {
		ui.Clear()

		for _, c := range components {
			ui.Render(c.GetUiElement())
		}

		e := <-event

		switch e.ID {
		case "<C-c>", "q":
			return
		case "n":
			createPipe()
		}

		for _, c := range components {
			c.Update(e)
		}
	}

}
