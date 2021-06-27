package main

import (
	"fmt"

	ui "github.com/gizak/termui/v3"
)

func init() {
	if err := ui.Init(); err != nil {
		panic(err)
	}

	components = append(components, new(ComponentInfoBar))
	components = append(components, new(ComponentPipesTree))
	components = append(components, new(ComponentPipeInfo))

	for _, c := range components {
		err := c.Init()
		if err == nil {
			continue
		}

		fmt.Println("Initialization failed: ", err.Error())
	}
}
