package net

import "net"
import "sync"
import "sync/atomic"
import "core/log"
import "core/debug"

type SessionList interface {
	Fetch(func(*Session))
}

// Server.
type Server struct {
	listener         *Listener
	sendChanBuff     uint
	serverStopHook   func(*Server)
	sessionStartHook func(*Session)
	sessionCloseHook func(*Session)
	maxSessionId     uint64
	sessions         map[uint64]*Session
	sessionMutex     sync.Mutex
	stopChan         chan int
	stopFlag         int32
	started          bool
	*settings

	// Present your data by a state object.
	State interface{}
}

// Listen and serve.
func Serve(network, address string, protocol PacketProtocol) (*Server, error) {
	listener, err := Listen(network, address, protocol)
	if err != nil {
		return nil, err
	}
	return NewServer(listener), nil
}

// Create a server.
func NewServer(listener *Listener) *Server {
	return &Server{
		listener:     listener,
		sendChanBuff: 1024,
		maxSessionId: 0,
		sessions:     make(map[uint64]*Session),
		stopChan:     make(chan int),
		settings:     &listener.settings,
	}
}

// Get listener address.
func (server *Server) Addr() net.Addr {
	return server.listener.Addr()
}

// Set session send channel buffer size setting.
// New setting will effect on new sessions.
func (server *Server) SetSendChanBuff(size uint) {
	server.sendChanBuff = size
}

// Get current session send chan buffer size setting.
func (server *Server) GetSendChanBuff() uint {
	return server.sendChanBuff
}

// Set server stop hook. Hook will invoked when server stop and all session closed.
func (server *Server) SetServerStopHook(hook func(*Server)) {
	server.serverStopHook = hook
}

// Set session start hook. Hook will invoked when a new session start.
func (server *Server) SetSessionStartHook(hook func(*Session)) {
	server.sessionMutex.Lock()
	defer server.sessionMutex.Unlock()
	server.sessionStartHook = hook
}

// Set session close hook. Hook will invoked when a session closed.
func (server *Server) SetSessionCloseHook(hook func(*Session)) {
	server.sessionMutex.Lock()
	defer server.sessionMutex.Unlock()
	server.sessionCloseHook = hook
}

// Start server.
func (server *Server) Start() {
	if !server.started {
		server.started = true
		go server.acceptLoop()
	}
}

// Stop server.
func (server *Server) Stop() {
	if atomic.CompareAndSwapInt32(&server.stopFlag, 0, 1) {
		server.listener.Close()

		// waitting for accept loop exit
		<-server.stopChan

		// waitting for all session closed
		wg := new(sync.WaitGroup)
		server.closeSessions(wg)
		wg.Wait()

		if server.serverStopHook != nil {
			server.serverStopHook(server)
		}
	}
}

// Broadcast to sessions. The response only encoded one time
// so performance better than send response one by one.
func (server *Server) Broadcast(sessions SessionList, response Response) {
	protocol := server.listener.protocol
	headSize := protocol.HeadSize()

	buffer := NewBuffer(
		make([]byte, headSize+response.ByteSize()),
	)

	// write packet body
	buffer.SetWritePosition(headSize)
	response.Encode(buffer)

	// write packet head
	data := buffer.Get()
	protocol.EncodeHead(data, uint(len(data)-headSize))

	sessions.Fetch(func(session *Session) {
		// // TODO：目前效率很低（为了解决加密后的广播会导致其他玩家不能进入游戏）
		// buf := make([]byte, len(data))
		// copy(buf, data)
		// session.SendRaw(buf)
		session.SendRaw(data)
	})
}

// Close all sessions.
func (server *Server) closeSessions(wg *sync.WaitGroup) {
	server.sessionMutex.Lock()
	defer server.sessionMutex.Unlock()

	// override session close hook to update wait group
	originSessionCloseHook := server.sessionCloseHook
	server.sessionCloseHook = func(session *Session) {
		defer wg.Done()
		if originSessionCloseHook != nil {
			originSessionCloseHook(session)
		}
	}

	for _, session := range server.sessions {
		wg.Add(1)
		go session.Close()
	}
}

// Loop and accept connections until get an error.
func (server *Server) acceptLoop() {
	defer func() {
		if err := recover(); err != nil {
			log.Errorf(`core.net.server.acceptLoop
Error = %v
Stack = 
%s`,
				err,
				debug.Stack(1, "    "),
			)
		}
	}()

	defer func() {
		close(server.stopChan)
		server.Stop()
	}()

	for {
		conn, err := server.listener.Accept()
		if err != nil {
			break
		}
		go server.startSession(conn)
	}
}

// Start a session to present the connection.
func (server *Server) startSession(conn *Conn) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorf(`core.net.server.startSession
Error = %v
Stack = 
%s`,
				err,
				debug.Stack(1, "    "),
			)
		}
	}()

	sessionId := atomic.AddUint64(&server.maxSessionId, 1)
	session := NewSession(sessionId, conn, server.GetSendChanBuff(), server.closeSession)
	startHook := server.putSession(session)
	if startHook != nil {
		startHook(session)
	}
	session.start()
}

// Close  and remove a session from server.
func (server *Server) closeSession(session *Session) {
	closeHook := server.delSession(session)
	if closeHook != nil {
		closeHook(session)
	}
}

// Put a session into session list
func (server *Server) putSession(session *Session) func(*Session) {
	server.sessionMutex.Lock()
	defer server.sessionMutex.Unlock()
	server.sessions[session.id] = session
	return server.sessionStartHook
}

// Delete a session from session list
func (server *Server) delSession(session *Session) func(*Session) {
	server.sessionMutex.Lock()
	defer server.sessionMutex.Unlock()
	delete(server.sessions, session.id)
	// In this scope, server have not way to override close hook.
	//
	// But if this method returns. Server may be take the lock
	// and override close hook before the closing session invoke it.
	//
	// So we returns current close hook.
	// But we not invoke hook in this scope.
	// Because the callback may be execute slowly, it will let server
	// can't accept new session or close other session.
	//
	// See Server.Stop().
	return server.sessionCloseHook
}
