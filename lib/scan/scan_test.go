package scan

import (
	"sync"

	"github.com/eacha/mercury/lib/conn"
	. "gopkg.in/check.v1"
)

type scanSuite struct{}

var (
	_ = Suite(&scanSuite{})
)

type dataMock struct {
	Fieled1 int    `json:"field1"`
	Field2  string `json:"field2"`
}

func scannableMock(options *Options, address string) Data {
	return dataMock{Fieled1: 1, Field2: "test"}
}

func (s *scanSuite) TestScan(c *C) {
	var (
		wg      sync.WaitGroup
		options = Options{
			InputChan:         make(chan string, 1),
			OutputChan:        make(chan Data, 1),
			Protocol:          conn.TCP,
			ConnectionTimeout: 1,
			IOTimeout:         1,
		}
		stat = NewStatistic(1)
	)
	wg.Add(1)
	options.WaitGroup = &wg

	options.InputChan <- ""
	close(options.InputChan)

	Scan(&options, stat, scannableMock)
	response := <-options.OutputChan

	c.Assert(response, Equals, "{\"field1\":1,\"field2\":\"test\"}")
}
