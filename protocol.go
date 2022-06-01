package gev

import (
	"github.com/maoge/paas-toolkit-go/ringbuffer"
)

var _ Protocol = &DefaultProtocol{}

// Protocol 自定义协议编解码接口
type Protocol interface {
	UnPacket(c *Connection, buffer *ringbuffer.RingBuffer) (interface{}, []byte)
	Packet(c *Connection, data interface{}) []byte
}

// DefaultProtocol 默认 Protocol
type DefaultProtocol struct{}

// UnPacket 拆包
func (d *DefaultProtocol) UnPacket(c *Connection, buffer *ringbuffer.RingBuffer) (interface{}, []byte) {
	s, e := buffer.PeekAll()
	if len(e) > 0 {
		size := len(s) + len(e)
		userBuffer := *c.UserBuffer()
		if size > cap(userBuffer) {
			userBuffer = make([]byte, size)
			*c.UserBuffer() = userBuffer
		}

		copy(userBuffer, s)
		copy(userBuffer[len(s):], e)
		buffer.RetrieveAll()

		return nil, userBuffer[:size]
	} else {
		buffer.RetrieveAll()

		return nil, s
	}
}

// Packet 封包
func (d *DefaultProtocol) Packet(c *Connection, data interface{}) []byte {
	return data.([]byte)
}
