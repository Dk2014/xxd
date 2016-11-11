package net

import "time"
import "bytes"
import "testing"

func Test_Server(t *testing.T) {
	server, err0 := Serve("tcp", "0.0.0.0:0", PacketN(4, BigEndian))
	if err0 != nil {
		t.Fatalf("Setup server failed, Error = %v", err0)
	}
	addr := server.Addr().String()
	t.Logf("Server: %v", addr)

	var (
		sessionStartCount  int
		sessionCloseCount  int
		sessionMatchFailed bool
		messageCount       int
		messageMatchFailed bool
		message            = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	)

	server.SetSessionStartHook(func(session1 *Session) {
		t.Log("Session start")
		sessionStartCount += 1
		session1.SetRequestHandlerFunc(func(session2 *Session, msg []byte) {
			messageCount += 1
			if session1 != session2 {
				sessionMatchFailed = true
			}
			if !bytes.Equal(msg, message) {
				messageMatchFailed = true
			}
		})
	})

	server.SetSessionCloseHook(func(session *Session) {
		t.Log("Session close")
		sessionCloseCount += 1
	})

	server.Start()

	client1, err1 := DialTimeout("tcp", addr, PacketN(4, BigEndian), time.Second)
	if err1 != nil {
		t.Fatal("Create client1 failed, Error = %v", err1)
	}

	client2, err2 := DialTimeout("tcp", addr, PacketN(4, BigEndian), time.Second)
	if err2 != nil {
		t.Fatal("Create client2 failed, Error = %v", err2)
	}

	client1.Write(message)
	client2.Write(message)
	client1.Write(message)
	client2.Write(message)

	time.Sleep(time.Second)

	// close by manual
	client1.Close()

	server.Stop()

	if sessionStartCount != 2 {
		t.Fatal("Session start count not match")
	}

	if sessionCloseCount != 2 {
		t.Fatal("Session close count not match")
	}

	if sessionMatchFailed {
		t.Fatal("Session match failed")
	}

	if messageCount != 4 {
		t.Fatal("Message count not match")
	}

	if messageMatchFailed {
		t.Fatal("Message match failed")
	}
}
