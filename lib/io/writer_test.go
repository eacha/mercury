package io

import (
	"log"
	"os"
	"time"

	"bufio"
	. "gopkg.in/check.v1"
)

type WriterSuite struct{}

var _ = Suite(&WriterSuite{})

var outputName = "write.txt"

func (s *WriterSuite) SetUpSuite(c *C) {
	file, err := os.Create(inputName)
	if err != nil {
		log.Fatal(err)
	}

	file.WriteString("1234\n")
	file.WriteString("4567\n")

	file.Close()
}

func (s *WriterSuite) TearDownSuite(c *C) {
	os.Remove(inputName)
	os.Remove(outputName)
}

func (s *WriterSuite) TestNewWriter(c *C) {
	writer, _ := NewWriter(outputName, 1)

	c.Assert(writer.file.Name(), Equals, outputName)

	writer.file.Close()
}

func (s *WriterSuite) TestNewWriterPipe(c *C) {
	writer, _ := NewWriter("", 1)

	c.Assert(writer.file, Equals, os.Stdout)
}

func (s *WriterSuite) TestWriteChannel(c *C) {
	writer, _ := NewWriter(outputName, 1)
	go writer.WriteJson()

	response := writer.GetQueue()

	response <- "1234"
	response <- "4567"
	response <- FINISH_WRITE

	time.Sleep(time.Second * 1)

	file, err := os.Open(outputName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	line, _, err := reader.ReadLine()
	c.Assert(string(line), Equals, "1234")

	line, _, err = reader.ReadLine()
	c.Assert(string(line), Equals, "4567")
}
