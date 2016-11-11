package net

import "io"
import "net"
import "time"
import "sync"
import "errors"

// Dial connects to the address on the named network.
func Dial(network, address string, protocol PacketProtocol) (*Conn, error) {
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return NewConn(conn, protocol), nil
}

// Dial connects to the address on the named network.
func DialTimeout(network, address string, protocol PacketProtocol, timeout time.Duration) (*Conn, error) {
	conn, err := net.DialTimeout(network, address, timeout)
	if err != nil {
		return nil, err
	}
	return NewConn(conn, protocol), nil
}

// The read packet connection
type Conn struct {
	conn      net.Conn       // The connection
	protocol  PacketProtocol // Packet protocol
	rHeadBuff []byte         // head buffer for read
	wHeadBuff []byte         // head buffer for write
	rMutext   sync.Mutex     // Read lock
	wMutext   sync.Mutex     // Write lock
	settings                 // Settings

	recvBytes uint32 // decode
	sendBytes uint32 // encode
}

// Create a real packet connection.
func NewConn(conn net.Conn, protocol PacketProtocol) *Conn {
	c := &Conn{
		conn:      conn,
		protocol:  protocol,
		rHeadBuff: make([]byte, protocol.HeadSize()),
		wHeadBuff: make([]byte, protocol.HeadSize()),
		recvBytes: 0,
		sendBytes: 20141021, // 第一次通讯使用魔数20141021作为期望值
	}
	return c
}

// Get the connection
func (c *Conn) Conn() net.Conn {
	return c.conn
}

// Get the packet protocol.
func (c *Conn) GetProtocol() PacketProtocol {
	return c.protocol
}

// Set the packet protocol.
func (c *Conn) SetProtocol(protocol PacketProtocol) {
	c.protocol = protocol
}

// Close the connection.
func (c *Conn) Close() error {
	return c.conn.Close()
}

// LocalAddr returns the local network address.
func (c *Conn) LocalAddr() net.Addr {
	return c.conn.LocalAddr()
}

// RemoteAddr returns the remote network address.
func (c *Conn) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

func (c *Conn) alloc(size uint) []byte {
	return make([]byte, size)
}

// Read a packet.
func (c *Conn) Read() ([]byte, error) {
	c.rMutext.Lock()
	defer c.rMutext.Unlock()

	if c.rTimeout > 0 {
		c.conn.SetReadDeadline(time.Now().Add(c.rTimeout))
	} else {
		c.conn.SetReadDeadline(time.Time{})
	}

	if size, err := c.protocol.ReadHead(c.conn, c.rHeadBuff); err == nil {
		var recvBytes uint32
		size, recvBytes = c.decodePacketHead(c.rHeadBuff)

		if c.rMaxSize > 0 && size > c.rMaxSize {
			return nil, errors.New("size big then max read size")
		}
		if recvBytes <= c.recvBytes {
			return nil, errors.New("[Conn::Read] Replay request")
		}
		c.recvBytes = recvBytes
		data := c.alloc(size)
		_, err := io.ReadFull(c.conn, data)
		if err == nil {
			c.decodePacket(data)
		}
		return data, err
	} else {
		return nil, err
	}
}

// Write a packet.
func (c *Conn) Write(data []byte) error {
	c.wMutext.Lock()
	defer c.wMutext.Unlock()

	size := uint(len(data))
	if c.wMaxSize > 0 && size > c.wMaxSize {
		return errors.New("size big then max write size")
	}

	if c.wTimeout > 0 {
		c.conn.SetWriteDeadline(time.Now().Add(c.wTimeout))
	} else {
		c.conn.SetWriteDeadline(time.Time{})
	}

	c.encodePacketHead(c.wHeadBuff, size)
	_, err := c.conn.Write(c.wHeadBuff)
	if err != nil {
		return err
	}

	c.encodePacket(data, size, 0)
	_, err = c.conn.Write(data)
	return err
}

// Write a raw data. The raw data is a fully packet include head and body.
func (c *Conn) WriteRaw(data []byte) error {
	c.wMutext.Lock()
	defer c.wMutext.Unlock()

	bodySize, err := c.protocol.DecodeHead(data)
	if err != nil {
		return err
	}

	if c.wMaxSize > 0 && bodySize > c.wMaxSize {
		return errors.New("size big then max write size")
	}

	if c.wTimeout > 0 {
		c.conn.SetWriteDeadline(time.Now().Add(c.wTimeout))
	} else {
		c.conn.SetWriteDeadline(time.Time{})
	}

	c.encodePacketHead(data, bodySize)
	c.encodePacket(data, uint(len(data)), c.protocol.HeadSize())
	_, err = c.conn.Write(data)

	return err
}

// encode/decode统一小端处理
// [size(4), moduleID(1), actionID(1), data(n)] encode to [{sendBytes^size(高4)|sendBytes(低4)}(8), encode....]
func (c *Conn) encodePacketHead(data []byte, size uint) {
	return
	// HEAD(8byte) = sendBytes^size(4byte)+sendBytes(4byte)
	data[0] = byte(c.sendBytes)
	data[1] = byte(c.sendBytes >> 8)
	data[2] = byte(c.sendBytes >> 16)
	data[3] = byte(c.sendBytes >> 24)

	// 保护包长度
	v := c.sendBytes ^ uint32(size)
	data[4] = byte(v)
	data[5] = byte(v >> 8)
	data[6] = byte(v >> 16)
	data[7] = byte(v >> 24)
}

func (c *Conn) encodePacket(data []byte, size uint, pos int) {
	return
	code := byte(c.sendBytes)
	// 保护包内容
	for ; pos <= int(size)-1; pos++ {
		data[pos] ^= code
	}
	c.sendBytes += uint32(size)
}

func (c *Conn) decodePacketHead(data []byte) (size uint, recvBytes uint32) {
	recvBytes = uint32(data[0]) | uint32(data[1])<<8 | uint32(data[2])<<16 | uint32(data[3])<<24
	size = uint(recvBytes ^ uint32((uint(data[4]) | uint(data[5])<<8 | uint(data[6])<<16 | uint(data[7])<<24)))
	return
}

//[{size(高4)|recvBytes(低4)}(8), encode....]
func (c *Conn) decodePacket(data []byte) {
	code := byte(c.recvBytes)
	for i := 0; i < len(data); i++ {
		data[i] ^= code
	}
}
