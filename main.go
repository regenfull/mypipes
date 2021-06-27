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

		if e.ID == "<C-c>" {
			break
		}

		if e.ID == "q" {
			break
		}

		for _, c := range components {
			c.Update(e)
		}
	}

}
