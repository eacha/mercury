package scan

import (
	"log"
	"os"

	"testing"

	"bufio"
	. "gopkg.in/check.v1"
	"time"
)

func Test(t *testing.T) { TestingT(t) }

type IOSuite struct{}

var (
	_          = Suite(&IOSuite{})
	inputName  = "read.txt"
	outputName = "write.txt"
)

func (s *IOSuite) SetUpSuite(c *C) {
	file, err := os.Create(inputName)
	if err != nil {
		log.Fatal(err)
	}

	file.WriteString("1234\n")
	file.WriteString("4567")

	file.Close()
}

func (s *IOSuite) TearDownSuite(c *C) {
	os.Remove(inputName)
	os.Remove(outputName)
}

func (s *IOSuite) TestReaderNotFoundFile(c *C) {
	_, err := NewReader("error.txt", 1)

	c.Assert(err.Error(), Equals, "open error.txt: no such file or directory")
}

func (s *IOSuite) TestNewReader(c *C) {
	reader, _ := NewReader(inputName, 1)

	c.Assert(reader.file.Name(), Equals, inputName)

	reader.file.Close()
}

func (s *IOSuite) TestNewReaderPipe(c *C) {
	reader, _ := NewReader("", 1)

	c.Assert(reader.file, Equals, os.Stdin)
}

func (s *IOSuite) TestReadChannel(c *C) {
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

func (s *IOSuite) TestNewWriter(c *C) {
	writer, _ := NewWriter(outputName, 1)

	c.Assert(writer.file.Name(), Equals, outputName)

	writer.file.Close()
}

func (s *IOSuite) TestNewWriterPipe(c *C) {
	writer, _ := NewWriter("", 1)

	c.Assert(writer.file, Equals, os.Stdout)
}

func (s *IOSuite) TestWriteJsonChannel(c *C) {
	writer, _ := NewWriter(outputName, 1)
	go writer.WriteJson()

	response := writer.GetQueue()

	response <- "1234"
	response <- "4567"
	response <- FinishWrite

	time.Sleep(time.Second * 1)

	file, err := os.Open(outputName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	line, _, err := reader.ReadLine()
	c.Assert(string(line), Equals, "\"1234\"")

	line, _, err = reader.ReadLine()
	c.Assert(string(line), Equals, "\"4567\"")
}
