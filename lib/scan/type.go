package scan

import (
	"sync"
	"time"
)

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

type Scannable interface {
	Scan(options *Options, statistic *Statistic)
}
