package io

import (
	"bufio"
	"os"
)

type Reader struct {
	file      *os.File
	taskQueue chan string
}

func NewReader(fileName string, taskQueue chan string) (*Reader, error) {
	var (
		rd  Reader
		err error
	)

	switch fileName {
	case "":
		rd.file = os.Stdin
	default:
		rd.file, err = os.Open(fileName)
		if err != nil {
			return nil, err
		}
	}
	rd.taskQueue = taskQueue

	return &rd, nil
}

func (rd *Reader) ReadIP() {
	var reader *bufio.Reader

	if rd.file != os.Stdin {
		defer rd.file.Close()
	}

	reader = bufio.NewReader(rd.file)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			close(rd.taskQueue)
			return
		}
		rd.taskQueue <- string(line)
	}
}
