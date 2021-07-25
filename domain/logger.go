package domain

import (
	"log"
	"os"
)

func NewLogger() (*log.Logger, error) {
	path, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	path += "/.mypipes.log"

	f, err := os.OpenFile(
		path,
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0644,
	)

	if err != nil {
		return nil, err
	}

	return log.New(f, "", log.Lshortfile|log.Lmicroseconds|log.Ldate|log.Ltime), nil
}
