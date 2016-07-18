package io

import (
	"os"
)

const (
	FINISH_WRITE = "\\0"
)

type Writer struct {
	file          *os.File
	responseQueue chan string
}

func NewWriter(fileName string, bufferSize int) (*Writer, error) {
	var (
		wr  Writer
		err error
	)

	switch fileName {
	case "":
		wr.file = os.Stdout
	default:
		wr.file, err = os.Create(fileName)
		if err != nil {
			return nil, err
		}
	}
	wr.responseQueue = make(chan string, bufferSize)

	return &wr, nil
}

func (wr *Writer) GetQueue() chan string {
	return wr.responseQueue
}

func (wr *Writer) WriteJson() {
	if wr.file != os.Stdout {

	}

	for {
		if line := <-wr.responseQueue; line != FINISH_WRITE {
			wr.file.WriteString(line + "\n")
		} else {
			return
		}
	}
}
