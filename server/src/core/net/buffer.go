package net

import "encoding/binary"

type ByteOrder binary.ByteOrder

var (
	BigEndian    = binary.BigEndian
	LittleEndian = binary.LittleEndian
)

// Byte writer.
type Buffer struct {
	raw  []byte
	wpos int
	rpos int
}

func NewBuffer(raw []byte) *Buffer {
	return &Buffer{
		raw: raw,
	}
}

// Reset buffer content and read/writer position.
func (b *Buffer) Reset() {
	b.wpos = 0
	b.rpos = 0
}

// Set buffer content.
func (b *Buffer) Set(raw []byte) {
	b.raw = raw
}

// Get buffer content.
func (b *Buffer) Get() []byte {
	return b.raw[:b.wpos]
}

func (b *Buffer) Cap() int {
	return len(b.raw)
}

func (b *Buffer) Len() int {
	return int(b.wpos)
}

func (b *Buffer) SetWritePosition(pos int) {
	b.wpos = pos
}

func (b *Buffer) GetWritePosition() int {
	return b.wpos
}

func (b *Buffer) GetReadPosition() int {
	return b.rpos
}

func (b *Buffer) SetReadPosition(pos int) {
	b.rpos = pos
}

func (b *Buffer) grows(n int) int {
	begin := b.wpos
	b.wpos += n
	if b.wpos >= len(b.raw) {
		newRaw := make([]byte, len(b.raw)+n+512)
		copy(newRaw, b.raw)
		b.raw = newRaw
	}
	return begin
}

func (b *Buffer) WriteBytes(v []byte) {
	begin := b.grows(len(v))
	copy(b.raw[begin:b.wpos], v)
}

func (b *Buffer) WriteUint8(v uint8) {
	begin := b.grows(1)
	b.raw[begin] = v
}

func (b *Buffer) WriteUint16(v uint16, byteOrder ByteOrder) {
	begin := b.grows(2)
	byteOrder.PutUint16(b.raw[begin:b.wpos], v)
}

func (b *Buffer) WriteUint32(v uint32, byteOrder ByteOrder) {
	begin := b.grows(4)
	byteOrder.PutUint32(b.raw[begin:b.wpos], v)
}

func (b *Buffer) WriteUint64(v uint64, byteOrder ByteOrder) {
	begin := b.grows(8)
	byteOrder.PutUint64(b.raw[begin:b.wpos], v)
}

func (b *Buffer) WriteUint16LE(v uint16) {
	begin := b.grows(2)
	LittleEndian.PutUint16(b.raw[begin:b.wpos], v)
}

func (b *Buffer) WriteUint32LE(v uint32) {
	begin := b.grows(4)
	LittleEndian.PutUint32(b.raw[begin:b.wpos], v)
}

func (b *Buffer) WriteUint64LE(v uint64) {
	begin := b.grows(8)
	LittleEndian.PutUint64(b.raw[begin:b.wpos], v)
}

func (b *Buffer) WriteUint16BE(v uint16) {
	begin := b.grows(2)
	BigEndian.PutUint16(b.raw[begin:b.wpos], v)
}

func (b *Buffer) WriteUint32BE(v uint32) {
	begin := b.grows(4)
	BigEndian.PutUint32(b.raw[begin:b.wpos], v)
}

func (b *Buffer) WriteUint64BE(v uint64) {
	begin := b.grows(8)
	BigEndian.PutUint64(b.raw[begin:b.wpos], v)
}

func (b *Buffer) ReadBytes(n int) (x []byte) {
	begin := b.rpos
	b.rpos += n
	x = b.raw[begin:b.rpos]
	return
}

func (b *Buffer) ReadUint8() (x uint8) {
	x = b.raw[b.rpos]
	b.rpos += 1
	return
}

func (b *Buffer) ReadUint16(byteOrder ByteOrder) (x uint16) {
	begin := b.rpos
	b.rpos += 2
	x = byteOrder.Uint16(b.raw[begin:b.rpos])
	return
}

func (b *Buffer) ReadUint32(byteOrder ByteOrder) (x uint32) {
	begin := b.rpos
	b.rpos += 4
	x = byteOrder.Uint32(b.raw[begin:b.rpos])
	return
}

func (b *Buffer) ReadUint64(byteOrder ByteOrder) (x uint64) {
	begin := b.rpos
	b.rpos += 8
	x = byteOrder.Uint64(b.raw[begin:b.rpos])
	return
}

func (b *Buffer) ReadUint16LE() (x uint16) {
	begin := b.rpos
	b.rpos += 2
	x = LittleEndian.Uint16(b.raw[begin:b.rpos])
	return
}

func (b *Buffer) ReadUint32LE() (x uint32) {
	begin := b.rpos
	b.rpos += 4
	x = LittleEndian.Uint32(b.raw[begin:b.rpos])
	return
}

func (b *Buffer) ReadUint64LE() (x uint64) {
	begin := b.rpos
	b.rpos += 8
	x = LittleEndian.Uint64(b.raw[begin:b.rpos])
	return
}

func (b *Buffer) ReadUint16BE() (x uint16) {
	begin := b.rpos
	b.rpos += 2
	x = BigEndian.Uint16(b.raw[begin:b.rpos])
	return
}

func (b *Buffer) ReadUint32BE() (x uint32) {
	begin := b.rpos
	b.rpos += 4
	x = BigEndian.Uint32(b.raw[begin:b.rpos])
	return
}

func (b *Buffer) ReadUint64BE() (x uint64) {
	begin := b.rpos
	b.rpos += 8
	x = BigEndian.Uint64(b.raw[begin:b.rpos])
	return
}
