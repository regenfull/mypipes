package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"os"
	"time"
	"unsafe"
)

func createPipe() error {
	var b = make([]byte, 0)
	var str = *(*string)(unsafe.Pointer(&b))

	var filename = fmt.Sprintf("/tmp/%d%s", time.Now().UnixNano(), str)

	file, err := os.OpenFile(
		filename,
		os.O_RDWR|os.O_CREATE|os.O_EXCL,
		os.ModePerm,
	)

	if err != nil {
		return err
	}

	_, err = file.WriteString("Name=\nCommand=")
	if err != nil {
		return err
	}

	file.Close()

	err = launchExternalEditor(filename)
	if err != nil {
		return err
	}

	file, err = os.OpenFile(
		filename,
		os.O_RDWR,
		os.ModePerm,
	)

	if err != nil {
		return err
	}

	var buffer = make([]byte, 128)
	var data []byte

	for {
		n, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			file.Close()
			return err
		}

		if err == io.EOF {
			break
		}

		data = append(data, buffer[:n]...)
	}

	file.Close()

	os.Remove(filename)

	var dataRecord PipeRecord
	dataRecord.ParseFileds(string(data))

	pipes = append(pipes, dataRecord)

	var encodeBuffer bytes.Buffer
	var encoder = gob.NewEncoder(&encodeBuffer)
	encoder.Encode(pipes)

	return storage.Save(encodeBuffer)
}
