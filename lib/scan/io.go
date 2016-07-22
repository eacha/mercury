package scan

import (
	"bufio"
	"encoding/json"
	"os"
)

var FinishWrite Data = nil

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

type Writer struct {
	file          *os.File
	responseQueue chan Data
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
	wr.responseQueue = make(chan Data, bufferSize)

	return &wr, nil
}

func (wr *Writer) GetQueue() chan Data {
	return wr.responseQueue
}

func (wr *Writer) WriteJson() {
	if wr.file != os.Stdout {
		defer wr.file.Close()
	}

	for {
		if data := <-wr.responseQueue; data != FinishWrite {
			if j, err := json.Marshal(data); err == nil {
				wr.file.WriteString(string(j) + "\n")
			}
		} else {
			return
		}
	}
}
