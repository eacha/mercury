package test

import (
	"fmt"
	"log"
	"net"
	"regexp"
	"strconv"
	"time"
)

type BannerServer struct {
	Port      int
	ToWrite   []byte
	WriteWait time.Duration
}

func (bs *BannerServer) RunServer() {
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal(err)
	}
	bs.Port = parsePort(l.Addr().String())
	defer l.Close()

	conn, err := l.Accept()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()

	time.Sleep(bs.WriteWait * time.Second)
	_, err = conn.Write(bs.ToWrite)
	if err != nil {
		fmt.Println(err)
	}

	return
}

func parsePort(address string) int {
	re, _ := regexp.Compile(`.*?(\d+)$`)
	match := re.FindStringSubmatch(address)

	port, _ := strconv.Atoi(match[1])
	return port
}
