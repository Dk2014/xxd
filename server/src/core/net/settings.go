package net

import "time"

// Connection settings
type settings struct {
	rTimeout time.Duration // Read timeout
	wTimeout time.Duration // Write timeout
	rMaxSize uint          // Max read size
	wMaxSize uint          // Max write size
}

// Set timeout for reading.
func (s *settings) SetReadTimeout(timeout time.Duration) {
	s.rTimeout = timeout
}

// Get timeout for reading.
func (s *settings) GetReadTimeout() time.Duration {
	return s.rTimeout
}

// Set timeout for writing.
func (s *settings) SetWriteTimeout(timeout time.Duration) {
	s.wTimeout = timeout
}

// Get timeout for writing.
func (s *settings) GetWriteTimeout() time.Duration {
	return s.wTimeout
}

// Set max packet size for reading.
func (s *settings) SetMaxReadSize(size uint) {
	s.rMaxSize = size
}

// Get max packet size for reading.
func (s *settings) GetMaxReadSize() uint {
	return s.rMaxSize
}

// Set max packet size for writing.
func (s *settings) SetMaxWriteSize(size uint) {
	s.wMaxSize = size
}

// Set max packet size for writing.
func (s *settings) GetMaxWriteSize() uint {
	return s.wMaxSize
}
