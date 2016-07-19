package conn

import (
	"fmt"
)

const (
	ConnRefusedMsg  = "Connection refued by host"
	ConnTimeoutMsg  = "Connection timeout"
	SetTimeoutMsg   = "Can't set IO timeout"
	ReadTimeoutMsg  = "Read timeout"
	WriteTimeoutMsg = "Write timeout"
	ReadMsg         = "Can't do the read operation"
	WriteMsg        = "Can't do the write operation"
)

type ConnError struct {
	Msg     string
	Address string
}

func (e *ConnError) Error() string {
	return fmt.Sprintf("%s, Host: %s", e.Msg, e.Address)
}

type IOError struct {
	Msg     string
	Address string
}

func (e *IOError) Error() string {
	return fmt.Sprintf("%s, Host: %s", e.Msg, e.Address)
}

type IOTimeoutError struct {
	Msg     string
	Address string
}

func (e *IOTimeoutError) Error() string {
	return fmt.Sprintf("%s, Host: %s", e.Msg, e.Address)
}
