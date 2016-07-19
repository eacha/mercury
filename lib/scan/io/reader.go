package io

import (
	"bufio"
	"os"
)

type Reader struct {
	file      *os.File
	taskQueue chan string
}

func NewReader(fileName string, bufferSize int) (*Reader, error) {
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
	rd.taskQueue = make(chan string, bufferSize)

	return &rd, nil
}

func (rd *Reader) GetQueue() chan string {
	return rd.taskQueue
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
