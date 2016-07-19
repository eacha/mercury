package io

import (
	"log"
	"os"

	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type ReaderSuite struct{}

var (
	_         = Suite(&ReaderSuite{})
	inputName = "read.txt"
)

func (s *ReaderSuite) SetUpSuite(c *C) {
	file, err := os.Create(inputName)
	if err != nil {
		log.Fatal(err)
	}

	file.WriteString("1234\n")
	file.WriteString("4567")

	file.Close()
}

func (s *ReaderSuite) TearDownSuite(c *C) {
	os.Remove(inputName)
}

func (s *ReaderSuite) TestNotFoundFile(c *C) {
	_, err := NewReader("error.txt", 1)

	c.Assert(err.Error(), Equals, "open error.txt: no such file or directory")
}

func (s *ReaderSuite) TestNewReader(c *C) {
	reader, _ := NewReader(inputName, 1)

	c.Assert(reader.file.Name(), Equals, inputName)

	reader.file.Close()
}

func (s *ReaderSuite) TestNewReaderPipe(c *C) {
	reader, _ := NewReader("", 1)

	c.Assert(reader.file, Equals, os.Stdin)
}

func (s *ReaderSuite) TestReadChannel(c *C) {
	reader, _ := NewReader(inputName, 1)
	go reader.ReadIP()

	queue := reader.GetQueue()

	r, more := <-queue
	c.Assert(r, Equals, "1234")
	c.Assert(more, Equals, true)

	r, more = <-queue
	c.Assert(r, Equals, "4567")
	c.Assert(more, Equals, true)

	r, more = <-queue
	c.Assert(more, Equals, false)
}
