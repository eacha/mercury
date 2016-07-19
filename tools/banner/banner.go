package banner

import (
	"github.com/eacha/mercury/lib/conn"
)

const bufferSize = 1024

func ReadBanner(conn *conn.ConnTimeout, data *BannerData) error {
	buffer := make([]byte, bufferSize)
	len, err := conn.Read(buffer)
	data.Banner = string(buffer[0:len])

	if err != nil {
		return err
	}

	return nil
}
