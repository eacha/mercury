package io

import (
	"log"
	"os"
	"time"

	. "gopkg.in/check.v1"
	"bufio"
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
	response := make(chan string, 1)
	finish := make(chan bool, 1)
	writer, _ := NewWriter(outputName, response, finish)

	c.Assert(writer.file.Name(), Equals, outputName)

	writer.file.Close()
}

func (s *WriterSuite) TestNewWriterPipe(c *C) {
	response := make(chan string, 1)
	finish := make(chan bool, 1)
	writer, _ := NewWriter("", response, finish)

	c.Assert(writer.file, Equals, os.Stdout)
}

func (s *WriterSuite) TestWriteChannel(c *C) {
	response := make(chan string, 1)
	finish := make(chan bool, 1)
	writer, _ := NewWriter(outputName, response, finish)
	go writer.WriteJson()

	response <- "1234"
	finish <- true
	response <- "4567"

	time.Sleep(time.Second * 2)

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
