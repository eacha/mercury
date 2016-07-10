package io

import (
	"os"
	"time"
)

const WAITING_TIMEOUT = 1

type Writer struct {
	file          *os.File
	responseQueue chan string
	finish         chan bool
}

func NewWriter(fileName string, responseQueue chan string, finish chan bool) (*Writer, error) {
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
	wr.responseQueue = responseQueue
	wr.finish = finish

	return &wr, nil
}

func (wr *Writer) WriteJson() {
	if wr.file != os.Stdout {

	}

	for {
		select {
		case line := <- wr.responseQueue:
			wr.file.WriteString(line + "\n")
		case <- wr.finish:
			for {
				select {
				case line := <- wr.responseQueue:
					wr.file.WriteString(line + "\n")
				case <- time.After(time.Second * WAITING_TIMEOUT):
					return
				}
			}
		}
	}
}
