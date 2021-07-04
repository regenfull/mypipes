package main

import (
	"os"
	"os/exec"
)

func launchExternalEditor(filename string) error {
	editor, params := makeExternalEditorParams(filename)
	cmd := exec.Command(editor, params...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func makeExternalEditorParams(filename string) (string, []string) {
	return "vim", []string{filename}
}
