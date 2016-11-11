package net

import "bytes"
import "testing"

func Test_Buffer(t *testing.T) {
	var test = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	buffer := NewBuffer(nil)

	buffer.WriteUint8(12)
	buffer.WriteUint16(123, BigEndian)
	buffer.WriteUint32(123456789, LittleEndian)
	buffer.WriteUint64(1234567890, LittleEndian)
	buffer.WriteBytes(test)

	if buffer.ReadUint8() != 12 {
		t.Fatal("buffer.ReadByte()")
	}

	if buffer.ReadUint16(BigEndian) != 123 {
		t.Fatal("buffer.ReadUint16()")
	}

	if buffer.ReadUint32(LittleEndian) != 123456789 {
		t.Fatal("buffer.ReadUint32()")
	}

	if buffer.ReadUint64(LittleEndian) != 1234567890 {
		t.Fatal("buffer.ReadUint64()")
	}

	if !bytes.Equal(buffer.ReadBytes(10), test) {
		t.Fatal("buffer.ReadBytes()")
	}
}
