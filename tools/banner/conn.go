package banner

import (
	"github.com/eacha/mercury/lib/conn"
)

const bufferSize = 1024

//type bannerConn struct {
//	conn *conn.ConnTimeout
//}
//
//func newBannerConn(protocol, address string, port int, connectionTimeout, ioTimeout time.Duration) (*bannerConn, error) {
//	var (
//		c   bannerConn
//		err error
//	)
//	c.conn, err = conn.NewConnTimeout(protocol, address, port, connectionTimeout, ioTimeout)
//
//	return &c, err
//}

//func (bc *bannerConn) ReadBanner() (string, error) {
//	buffer := make([]byte, bufferSize)
//
//	_, err := bc.conn.Read(buffer)
//	if err != nil {
//		return "", err
//	}
//
//	return string(buffer), nil
//}

func ReadBanner(conn *conn.ConnTimeout, data *BannerData) error {
	buffer := make([]byte, bufferSize)
	len, err := conn.Read(buffer)
	data.Banner = string(buffer[0:len])

	if err != nil {
		return err
	}

	return nil
}
