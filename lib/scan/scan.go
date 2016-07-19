package scan

import (
	"encoding/json"
	"sync"
	"time"
)

type Data interface{}

type Options struct {
	// Basic Scan Setup
	WaitGroup *sync.WaitGroup

	InputFileName  string
	OutputFileName string

	InputChan  chan string
	OutputChan chan string

	Port              int
	Module            string
	Protocol          string
	Threads           uint
	ConnectionTimeout time.Duration
	IOTimeout         time.Duration

	// More options in the future
}

type scannable func(*Options, string) Data

func Scan(options *Options, statistic *Statistic, fn scannable) {
	defer options.WaitGroup.Done()
	for {
		address, more := <-options.InputChan
		if !more {
			break
		}
		statistic.IncreaseProcessedLines()

		data := fn(options, address)
		j, _ := json.Marshal(data)

		options.OutputChan <- string(j)
	}
	statistic.SetEndTime()
}
