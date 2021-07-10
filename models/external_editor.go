package models

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"
	"unsafe"
)

type ExternalEditor struct {
	editor string
	args   []string
}

func (e ExternalEditor) Launch(text string) (string, error) {
	var b = make([]byte, 0)
	var str = *(*string)(unsafe.Pointer(&b))

	var filename = fmt.Sprintf("/tmp/%d%s", time.Now().UnixNano(), str)

	file, err := os.OpenFile(
		filename,
		os.O_RDWR|os.O_CREATE|os.O_EXCL,
		os.ModePerm,
	)

	if err != nil {
		return "", err
	}

	_, err = file.WriteString(text)
	if err != nil {
		return "", err
	}

	file.Close()

	e.args = append(e.args, "filename")
	cmd := exec.Command(e.editor, e.args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
		return "", err
	}

	file, err = os.OpenFile(
		filename,
		os.O_RDWR,
		os.ModePerm,
	)

	if err != nil {
		return "", err
	}

	var buffer = make([]byte, 128)
	var data []byte

	for {
		n, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			file.Close()
			return "", err
		}

		if err == io.EOF {
			break
		}

		data = append(data, buffer[:n]...)
	}

	file.Close()
	os.Remove(filename)

	return string(data), nil
}

func NewExternalEditor() *ExternalEditor {
	return &ExternalEditor{"vim", []string{}}
}
