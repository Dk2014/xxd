package net

import "net"

// Listen announces on the local network address.
// The 'network' must be a stream-oriented network: "tcp", "tcp4", "tcp6", "unix" or "unixpacket".
// See net.Dial() for the syntax of laddr.
func Listen(network, address string, protocol PacketProtocol) (*Listener, error) {
	listener, err := net.Listen(network, address)
	if err != nil {
		return nil, err
	}
	return NewListener(listener, protocol), nil
}

// The packet-oriented listener
type Listener struct {
	listener net.Listener   // The listener
	protocol PacketProtocol // Packet protocol
	settings                // Settings
}

func NewListener(listener net.Listener, protocol PacketProtocol) *Listener {
	l := &Listener{
		listener: listener,
		protocol: protocol,
	}
	return l
}

// Get the listener.
func (l *Listener) Listener() net.Listener {
	return l.listener
}

// Close the listener.
func (l *Listener) Close() error {
	return l.listener.Close()
}

// Addr returns the listener's network address.
func (l *Listener) Addr() net.Addr {
	return l.listener.Addr()
}

// Accept waits for and returns the next connection to the listener.
func (l *Listener) Accept() (*Conn, error) {
	conn, err := l.listener.Accept()
	if err != nil {
		return nil, err
	}
	pconn := NewConn(conn, l.protocol)
	pconn.settings = l.settings
	return pconn, nil
}

// Get the packet protocol.
func (l *Listener) GetProtocol() PacketProtocol {
	return l.protocol
}

// Set the packet protocol.
func (l *Listener) SetProtocol(protocol PacketProtocol) {
	l.protocol = protocol
}
