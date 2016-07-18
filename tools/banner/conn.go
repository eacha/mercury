package banner

import (
	"github.com/eacha/aps/tools/conn"
	"time"
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
		return nil, err
	}

	return string(buffer), nil
}
