package net

import "net"
import "errors"
import "sync/atomic"
import "core/log"
import "core/debug"

var (
	SendToClosedError = errors.New("Send to closed session")
	BlockingError     = errors.New("Blocking happened")
)

// Request handler.
type RequestHandler interface {
	// Handle a request from session.
	Handle(*Session, []byte)
}

type SessionStateHandler interface {
	RecordResponse(rsp interface{})
}

type requestHandlerFunc struct {
	callback func(*Session, []byte)
}

func (handler requestHandlerFunc) Handle(session *Session, request []byte) {
	handler.callback(session, request)
}

// Response.
type Response interface {
	// Get recommend byte size for prepare buffer.
	ByteSize() int

	// Encode reponse into buffer.
	Encode(*Buffer)

	GetModuleIdAndActionId() (int8, int8)
}

// Session.
type Session struct {
	id             uint64
	conn           *Conn
	sendChan       chan Response
	sendRawChan    chan []byte
	sendPacketChan chan []byte
	rStopChan      chan int
	wStopChan      chan int
	closeChan      chan int
	closeFlag      int32
	handler        RequestHandler
	buffer         *Buffer
	closeCallback  func(*Session)
	*settings

	// Present your data by a state object.
	State SessionStateHandler
}

func NewSession(id uint64, conn *Conn, sendChanSize uint, closeCallback func(*Session)) *Session {
	return &Session{
		id:             id,
		conn:           conn,
		sendChan:       make(chan Response, sendChanSize),
		sendRawChan:    make(chan []byte, sendChanSize),
		sendPacketChan: make(chan []byte, sendChanSize),
		rStopChan:      make(chan int),
		wStopChan:      make(chan int),
		closeChan:      make(chan int),
		buffer:         NewBuffer(make([]byte, 1024)),
		closeCallback:  closeCallback,
		settings:       &conn.settings,
	}
}

func (session *Session) start() {
	go session.writeLoop()
	go session.readLoop()
}

// Loop and wait incoming requests.
func (session *Session) readLoop() {
	defer func() {
		if err := recover(); err != nil {
			log.Errorf(`core.net.session.readLoop
Error = %v
Stack = 
%s`,
				err,
				debug.Stack(1, "    "),
			)
		}
	}()
	defer func() {
		close(session.rStopChan)
		session.Close()
	}()
	for {
		msg, err := session.conn.Read()
		if err != nil {
			break
		}
		session.handler.Handle(session, msg)
	}
}

// Loop and transport responses.
func (session *Session) writeLoop() {
	defer func() {
		if err := recover(); err != nil {
			log.Errorf(`core.net.session.writeLoop
Error = %v
Stack = 
%s`,
				err,
				debug.Stack(1, "    "),
			)
		}
	}()
	defer func() {
		close(session.wStopChan)
		session.Close()
	}()
L:
	for {
		select {
		case response := <-session.sendChan:
			if err := session.writeResponse(response); err != nil {
				break L
			}
		case data := <-session.sendRawChan:
			if err := session.conn.WriteRaw(data); err != nil {
				break L
			}
		case data := <-session.sendPacketChan:
			if err := session.conn.Write(data); err != nil {
				break L
			}
		case <-session.closeChan:
			break L
		}
	}
}

func (session *Session) writeResponse(response Response) error {
	buffer := session.buffer
	protocol := session.conn.protocol

	// prepare buffer
	headSize := protocol.HeadSize()
	bodySize := response.ByteSize()
	if buffer.Cap() < (bodySize + headSize) {
		buffer.Set(make([]byte, (bodySize + headSize + 512)))
	}

	// write packet body
	buffer.SetWritePosition(headSize)
	response.Encode(buffer)

	// write packet head
	data := buffer.Get()
	protocol.EncodeHead(data, uint(len(data)-headSize))

	return session.conn.WriteRaw(data)
}

// Get session id.
func (session *Session) Id() uint64 {
	return session.id
}

// Get local address.
func (session *Session) LocalAddr() net.Addr {
	return session.conn.LocalAddr()
}

// Get remote address.
func (session *Session) RemoteAdd() net.Addr {
	return session.conn.RemoteAddr()
}

// Get the packet protocol.
func (session *Session) GetProtocol() PacketProtocol {
	return session.conn.GetProtocol()
}

// Set the packet protocol.
func (session *Session) SetProtocol(protocol PacketProtocol) {
	session.conn.SetProtocol(protocol)
}

// Set request handler.
func (session *Session) SetRequestHandler(handler RequestHandler) {
	session.handler = handler
}

func (session *Session) SetRequestHandlerFunc(callback func(*Session, []byte)) {
	session.handler = requestHandlerFunc{callback}
}

// Close session and remove it from api server.
func (session *Session) Close() {
	defer func() {
		if err := recover(); err != nil {
			log.Errorf(`core.net.session.Close
Error = %v
Stack = 
%s`,
				err,
				debug.Stack(1, "    "),
			)
		}
	}()

	if atomic.CompareAndSwapInt32(&session.closeFlag, 0, 1) {
		// if close session without this goroutine
		// deadlock will happen when session close by itself.
		go func() {
			defer func() {
				if err := recover(); err != nil {
					log.Errorf(`core.net.session.Close (2)
Error = %v
Stack = 
%s`,
						err,
						debug.Stack(1, "    "),
					)
				}
			}()
			session.conn.Close()
			// notify write loop session closed
			close(session.closeChan)
			// waitting for read loop stop
			<-session.rStopChan
			// waitting for write loop stop
			<-session.wStopChan
			// remove session from server
			if session.closeCallback != nil {
				session.closeCallback(session)
			}
		}()
	}
}

// Async send a response.
func (session *Session) Send(response Response) error {
	// 记录当前发的协议数据，用于改善较差的网络情况
	session.State.RecordResponse(response)

	if atomic.LoadInt32(&session.closeFlag) != 0 {
		return SendToClosedError
	}

	select {
	case session.sendChan <- response:
		return nil
	default:
		session.Close()
		return BlockingError
	}
}

// Async send a raw data. The raw data is a fully packet include head and body.
func (session *Session) SendRaw(data []byte) error {
	// 记录当前发的协议数据，用于改善较差的网络情况
	session.State.RecordResponse(data)

	if atomic.LoadInt32(&session.closeFlag) != 0 {
		return SendToClosedError
	}

	select {
	case session.sendRawChan <- data:
		return nil
	default:
		session.Close()
		return BlockingError
	}
}

// Async send a packet.
func (session *Session) SendPacket(data []byte) error {
	if atomic.LoadInt32(&session.closeFlag) != 0 {
		return SendToClosedError
	}

	select {
	case session.sendPacketChan <- data:
		return nil
	default:
		session.Close()
		return BlockingError
	}
}
