package main

import (
	ui "github.com/gizak/termui/v3"
)

func main() {
	var err error
	if err = ui.Init(); err != nil {
		panic(err)
	}

	defer ui.Close()

	application, err := setup()
	if err != nil {
		panic(err)
	}

	err = application.Run()
	if err != nil {
		panic(err)
	}

}
