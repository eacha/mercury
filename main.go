package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"encoding/json"

	"github.com/eacha/mercury/lib"
	"github.com/eacha/mercury/lib/conn"
	"github.com/eacha/mercury/lib/scan"
	"github.com/eacha/mercury/lib/scan/io"
	"github.com/eacha/mercury/tools/banner"
)

const (
	inputChannelBuffer  = 1
	outputChannelBuffer = 1
)

var (
	modulesList       = []string{"Banner"}
	protocolList      = []string{conn.UDP, conn.TCP}
	showModules       bool
	showProtocols     bool
	options           scan.Options
	connectionTimeout uint
	ioTimeout         uint
)

func init() {
	flag.BoolVar(&showModules, "module-list", false, "Print module list and exit")
	flag.BoolVar(&showProtocols, "protocol-list", false, "Print protocol list and exit")

	flag.StringVar(&options.InputFileName, "input-file", "", "Input file name, empty for stdin")
	flag.StringVar(&options.OutputFileName, "output-file", "", "Output file name, empty for stdout")
	flag.IntVar(&options.Port, "port", 0, "Port number to scan")
	flag.StringVar(&options.Module, "module", "", "Set module to scan")
	flag.StringVar(&options.Protocol, "protocol", conn.TCP, "Set protocol to scan")
	flag.UintVar(&options.Threads, "threads", 1, "Set the number of corutines")
	flag.UintVar(&connectionTimeout, "connection-timeout", 10, "Set connection timeout in seconds")
	flag.UintVar(&ioTimeout, "io-timeout", 10, "Set input output timeout in seconds")

	flag.Parse()

	// Help arguments
	if showModules {
		printModules()
	}

	if showProtocols {
		printProtocols()
	}

	// Check the arguments
	if options.Port < 0 || options.Port > 65535 {
		log.Fatal("--port must be in the range [0, 65535]")
	}

	if options.Module == "" || !lib.StringInSlice(options.Module, modulesList) {
		log.Fatal("--module must be in the --module-list")
	}

	if !lib.StringInSlice(options.Protocol, protocolList) {
		log.Fatal("--protocol must be in the --protocol-list")
	}

	if connectionTimeout <= 0 && ioTimeout <= 0 {
		log.Fatal("--connection-timeout and  --io-timeout must be positive")
	}

	options.ConnectionTimeout = time.Duration(connectionTimeout)
	options.IOTimeout = time.Duration(ioTimeout)
}

func printModules() {
	fmt.Println("Modules:")
	for _, mod := range modulesList {
		fmt.Printf("\t- %s\n", mod)
	}
	os.Exit(0)
}

func printProtocols() {
	fmt.Println("protocols:")
	for _, mod := range protocolList {
		fmt.Printf("\t- %s\n", mod)
	}
	os.Exit(0)
}

func main() {
	var (
		wg sync.WaitGroup
		ts = make([]*scan.Statistic, int(options.Threads))
	)

	reader, err := io.NewReader(options.InputFileName, inputChannelBuffer)
	if err != nil {
		log.Fatal("Can't open the input file")
	}

	writer, err := io.NewWriter(options.OutputFileName, outputChannelBuffer)
	if err != nil {
		log.Fatal("Can't open the output file")
	}

	options.WaitGroup = &wg
	options.InputChan = reader.GetQueue()
	options.OutputChan = writer.GetQueue()

	go reader.ReadIP()
	go writer.WriteJson()

	switch options.Module {
	case "Banner":
		for i := 0; i < int(options.Threads); i++ {
			wg.Add(1)
			ts[i] = scan.NewStatistic(i)
			go scan.Scan(&options, ts[i], banner.HostScan)
		}
	}

	wg.Wait()
	options.OutputChan <- io.FINISH_WRITE
	time.Sleep(10 * time.Second)

	for _, value := range ts {
		j, _ := json.Marshal(*value)
		fmt.Println(string(j))
	}

}
