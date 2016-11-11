package mysql

import (
	"reflect"
	"unsafe"
)

// StringArena lets you consolidate allocations for a group of strings
// that have similar life length
type StringArena struct {
	buf []byte
	str string
}

// NewStringArena creates an arena of the specified size.
func NewStringArena(size int) *StringArena {
	self := &StringArena{buf: make([]byte, 0, size)}
	pbytes := (*reflect.SliceHeader)(unsafe.Pointer(&self.buf))
	pstring := (*reflect.StringHeader)(unsafe.Pointer(&self.str))
	pstring.Data = pbytes.Data
	pstring.Len = pbytes.Cap
	return self
}

// NewString copies a byte slice into the arena and returns it as a string.
// If the arena is full, it returns a traditional go string.
func (self *StringArena) NewString(b []byte) string {
	if len(self.buf)+len(b) > cap(self.buf) {
		return string(b)
	}
	start := len(self.buf)
	self.buf = append(self.buf, b...)
	return self.str[start : start+len(b)]
}

// SpaceLeft returns the amount of space left in the arena.
func (self *StringArena) SpaceLeft() int {
	return cap(self.buf) - len(self.buf)
}
