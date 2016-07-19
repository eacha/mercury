package banner

import (
	"testing"

	"time"

	"github.com/eacha/mercury/lib/conn"
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
