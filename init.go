package main

import (
	"fmt"

	ui "github.com/gizak/termui/v3"
)

func init() {
	var err error
	err = storage.Init()
	if err != nil {
		panic(err)
	}

	if err = ui.Init(); err != nil {
		panic(err)
	}

	pipesBuffer, err := storage.Read()
	if err != nil {
		panic(err)
	}

	pipes, err = parsePipes(pipesBuffer)
	if err != nil {
		panic(err)
	}

	var infoBar = new(ComponentInfoBar)
	var pipeTree = new(ComponentPipesTree)
	var pipeInfo = new(ComponentPipeInfo)
	pipeInfo.SetPipesTree(pipeTree)

	components = append(components, infoBar)
	components = append(components, pipeTree)
	components = append(components, pipeInfo)

	for _, c := range components {
		err := c.Init()
		if err == nil {
			continue
		}

		fmt.Println("Initialization failed: ", err.Error())
	}
}
