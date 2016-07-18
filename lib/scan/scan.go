package scan

import "encoding/json"

func Scan(options *Options, statistic *Statistic, fn scannable) {
	defer options.WaitGroup.Done()
	for {
		address, more := <-options.InputChan
		if !more {
			break
		}
		statistic.IncreaseProcessedLines()

		data := fn(options, address)
		j, _ := json.Marshal(*data)

		options.OutputChan <- string(j)
	}
	statistic.SetEndTime()
}
