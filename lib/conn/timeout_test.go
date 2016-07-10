package conn

import (
	"encoding/base64"
	"testing"

	"time"

	"github.com/eacha/mercury/lib/test"
	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type ConnTimeoutSuite struct{}

var _ = Suite(&ConnTimeoutSuite{})

var BUFFER = "JwAAAAo0L"

func (s *ConnTimeoutSuite) TestNewConnectionRefuse(c *C) {
	conn, err := NewConnTimeout(TCP, "", 1, 10, 10)

	c.Assert(conn, IsNil)
	c.Assert(err, DeepEquals, &ConnError{ConnRefusedMsg, ""})
}

func (s *ConnTimeoutSuite) TestReadError(c *C) {
	sendData, _ := base64.StdEncoding.DecodeString(BUFFER)
	server := test.BannerServer{ToWrite: sendData, WriteWait: 0}

	// Server
	go (&server).RunServer()
	time.Sleep(100 * time.Millisecond)

	// Client
	buffer := make([]byte, 10)
	conn, _ := NewConnTimeout(TCP, "", server.Port, 10, 10)
	conn.Close()

	_, err := conn.Read(buffer)
	c.Assert(err, DeepEquals, &IOError{ReadMsg, ""})
}

func (s *ConnTimeoutSuite) TestWriteError(c *C) {
	sendData, _ := base64.StdEncoding.DecodeString(BUFFER)
	server := test.BannerServer{ToWrite: sendData, WriteWait: 0}

	// Server
	go (&server).RunServer()
	time.Sleep(100 * time.Millisecond)

	// Client
	conn, _ := NewConnTimeout(TCP, "", server.Port, 10, 10)
	conn.Close()

	_, err := conn.Write(sendData)
	c.Assert(err, DeepEquals, &IOError{WriteMsg, ""})

}

func (s *ConnTimeoutSuite) TestReadTimeout(c *C) {
	sendData, _ := base64.StdEncoding.DecodeString(BUFFER)
	server := test.BannerServer{ToWrite: sendData, WriteWait: 2}

	// Server
	go (&server).RunServer()
	time.Sleep(100 * time.Millisecond)

	// Client
	buffer := make([]byte, 10)
	conn, _ := NewConnTimeout(TCP, "", server.Port, 1, 1)
	defer conn.Close()

	_, err := conn.Read(buffer)
	c.Assert(err, DeepEquals, &IOTimeoutError{ReadTimeoutMsg, ""})
}

func (s *ConnTimeoutSuite) TestReadSuccess(c *C) {
	sendData, _ := base64.StdEncoding.DecodeString(BUFFER)
	server := test.BannerServer{ToWrite: sendData, WriteWait: 0}

	// Server
	go (&server).RunServer()
	time.Sleep(100 * time.Millisecond)

	// Client
	buf := make([]byte, 10)
	conn, _ := NewConnTimeout(TCP, "", server.Port, 10, 10)
	defer conn.Close()

	read, _ := conn.Read(buf)
	c.Assert(buf[:read], DeepEquals, sendData)
}

// Errors
func (s *ConnTimeoutSuite) TestConnError(c *C) {
	c.Assert((&ConnError{"test1", "test2"}).Error(), Equals, "test1, Host: test2")
	c.Assert((&IOError{"test1", "test2"}).Error(), Equals, "test1, Host: test2")
	c.Assert((&IOTimeoutError{"test1", "test2"}).Error(), Equals, "test1, Host: test2")
}
