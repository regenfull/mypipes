package main

import (
	"bytes"
	"encoding/gob"
	"io"
)

func parsePipes(b *bytes.Buffer) ([]PipeRecord, error) {
	var pipes []PipeRecord
	var err error

	var decoder = gob.NewDecoder(b)
	err = decoder.Decode(&pipes)
	if err == io.EOF {
		err = nil
	}
	return pipes, err
}
