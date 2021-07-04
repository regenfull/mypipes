package main

import (
	"errors"
	"regexp"
	"strings"
)

type PipeRecord struct {
	Name    string
	Command string
}

func (p *PipeRecord) ParseFileds(rawData string) error {
	var checkRegexp = `^Name=.*\nCommand=.*`

	if !regexp.MustCompile(checkRegexp).MatchString(rawData) {
		return errors.New("wrong raw data to parse")
	}

	var data = strings.Split(rawData, "\n")
	p.Name = strings.ReplaceAll(data[0], "Name=", "")
	p.Command = strings.ReplaceAll(data[1], "Command=", "")
	return nil
}
