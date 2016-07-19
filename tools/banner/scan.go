package banner

import (
	"github.com/eacha/mercury/lib/conn"
	"github.com/eacha/mercury/lib/scan"
)

func HostScan(options *scan.Options, address string) scan.Data {
	data := &BannerData{IP: address}

	conn, err := conn.NewConnTimeout(options.Protocol, address, options.Port, options.ConnectionTimeout, options.IOTimeout)
	if err != nil {
		data.Error = err.Error()
		return data
	}
	defer conn.Close()

	if err = ReadBanner(conn, data); err != nil {
		data.Error = err.Error()
		return data
	}

	return data
}
