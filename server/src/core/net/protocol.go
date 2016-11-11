package net

import "io"
import "net"
import "errors"

type PacketProtocol interface {
	HeadSize() int
	ReadHead(net.Conn, []byte) (uint, error)
	WriteHead(net.Conn, []byte, uint) error
	EncodeHead([]byte, uint)
	DecodeHead([]byte) (uint, error)
}

var (
	PacketHeadNotMatchError = errors.New("Packet head not match body length")
)

func PacketN(n int, bo ByteOrder) PacketNProtocol {
	if n != 1 && n != 2 && n != 4 && n != 8 {
		panic("unsupported packet head size")
	}
	return PacketNProtocol{
		n:  n,
		bo: bo,
	}
}

type PacketNProtocol struct {
	n  int
	bo ByteOrder
}

func (p PacketNProtocol) HeadSize() int {
	return p.n
}

func (p PacketNProtocol) ReadHead(conn net.Conn, buff []byte) (uint, error) {
	if _, err := io.ReadFull(conn, buff); err != nil {
		return 0, err
	}

	switch p.n {
	case 1:
		return uint(buff[0]), nil
	case 2:
		return uint(p.bo.Uint16(buff)), nil
	case 4:
		return uint(p.bo.Uint32(buff)), nil
	case 8:
		return uint(p.bo.Uint64(buff)), nil
	}

	panic("unsupported packet head size")
}

func (p PacketNProtocol) WriteHead(conn net.Conn, buff []byte, length uint) error {
	p.EncodeHead(buff, length)

	if _, err := conn.Write(buff); err != nil {
		return err
	}

	return nil
}

func (p PacketNProtocol) EncodeHead(data []byte, length uint) {
	switch p.n {
	case 1:
		data[0] = byte(length)
	case 2:
		p.bo.PutUint16(data, uint16(length))
	case 4:
		p.bo.PutUint32(data, uint32(length))
	case 8:
		p.bo.PutUint64(data, uint64(length))
	default:
		panic("unsupported packet head size")
	}
}

func (p PacketNProtocol) DecodeHead(data []byte) (uint, error) {
	var bodySize uint

	switch p.n {
	case 1:
		bodySize = uint(data[0])
	case 2:
		bodySize = uint(p.bo.Uint16(data))
	case 4:
		bodySize = uint(p.bo.Uint32(data))
	case 8:
		bodySize = uint(p.bo.Uint64(data))
	}

	if bodySize != uint(len(data)-p.n) {
		return 0, PacketHeadNotMatchError
	}

	return bodySize, nil
}
