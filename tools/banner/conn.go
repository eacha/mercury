package banner

import (
	"time"

	"github.com/eacha/mercury/lib/conn"
)

const bufferSize = 1024

type bannerConn struct {
	conn *conn.ConnTimeout
}

func newBanner(protocol, address string, port int, connectionTimeout, ioTimeout time.Duration) (*bannerConn, error) {
	var (
		c   bannerConn
		err error
	)
	c.conn, err = conn.NewConnTimeout(protocol, address, port, connectionTimeout, ioTimeout)

	return &c, err
}

func (bc *bannerConn) ReadBanner() (string, error) {
	buffer := make([]byte, bufferSize)

	_, err := bc.conn.Read(buffer)
	if err != nil {
		return "", err
	}

	return string(buffer), nil
}
