package banner

import (
	"testing"

	"time"

	"github.com/eacha/mercury/lib/conn"
	"github.com/eacha/mercury/lib/scan"
	"github.com/eacha/mercury/lib/test"
	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type bannerSuite struct{}

var _ = Suite(&bannerSuite{})

const banner = "Testing banner"

func (b *bannerSuite) TestReadBannerError(c *C) {
	data := &BannerData{}
	server := test.BannerServer{ToWrite: []byte(banner), WriteWait: 0}

	// Server
	go (&server).RunServer()
	time.Sleep(100 * time.Millisecond)

	// Client
	connection, _ := conn.NewConnTimeout(conn.TCP, "", server.Port, 1, 1)
	connection.Close()

	err := ReadBanner(connection, data)

	c.Assert(err, DeepEquals, &conn.IOError{conn.ReadMsg, ""})
	c.Assert(data.Banner, Equals, "")
}

func (b *bannerSuite) TestReadBanner(c *C) {
	data := &BannerData{}
	server := test.BannerServer{ToWrite: []byte(banner), WriteWait: 0}

	// Server
	go (&server).RunServer()
	time.Sleep(100 * time.Millisecond)

	// Client
	connection, _ := conn.NewConnTimeout(conn.TCP, "", server.Port, 1, 1)
	defer connection.Close()

	err := ReadBanner(connection, data)

	c.Assert(err, DeepEquals, nil)
	c.Assert(data.Banner, Equals, banner)
}

func (b *bannerSuite) TestHostScanConnError(c *C) {
	options := scan.Options{
		Protocol:          conn.TCP,
		Port:              1,
		ConnectionTimeout: 1,
		IOTimeout:         1,
	}

	data := HostScan(&options, "")

	c.Assert(data.(*BannerData).IP, Equals, "")
	c.Assert(data.(*BannerData).Error, Equals, "Connection refued by host, Host: ")
}

func (b *bannerSuite) TestHostScanReadTimeout(c *C) {
	server := test.BannerServer{ToWrite: []byte(banner), WriteWait: 2}
	options := scan.Options{
		Protocol:          conn.TCP,
		ConnectionTimeout: 1,
		IOTimeout:         1,
	}

	// Server
	go (&server).RunServer()
	time.Sleep(100 * time.Millisecond)

	options.Port = server.Port
	data := HostScan(&options, "")

	c.Assert(data.(*BannerData).IP, Equals, "")
	c.Assert(data.(*BannerData).Error, Equals, "Read timeout, Host: ")
}

func (b *bannerSuite) TestHostScan(c *C) {
	server := test.BannerServer{ToWrite: []byte(banner), WriteWait: 0}
	options := scan.Options{
		Protocol:          conn.TCP,
		ConnectionTimeout: 1,
		IOTimeout:         1,
	}

	// Server
	go (&server).RunServer()
	time.Sleep(100 * time.Millisecond)

	options.Port = server.Port
	data := HostScan(&options, "")

	c.Assert(data.(*BannerData).IP, Equals, "")
	c.Assert(data.(*BannerData).Banner, Equals, "Testing banner")
}
