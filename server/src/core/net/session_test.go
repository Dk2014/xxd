package net

import "time"
import "bytes"
import "testing"

type ByteResponse []byte

// Get recommend byte size for prepare buffer.
func (r ByteResponse) ByteSize() int {
	return len(r)
}

// Encode reponse into buffer.
func (r ByteResponse) Encode(b *Buffer) {
	b.WriteBytes(r)
}

func Test_Session(t *testing.T) {
	server, err0 := Serve("tcp", "0.0.0.0:0", PacketN(4, BigEndian))
	if err0 != nil {
		t.Fatalf("Setup server failed, Error = %v", err0)
	}
	addr := server.Addr().String()
	t.Logf("Server: %v", addr)

	var (
		sessionStartCount int
		sessionCloseCount int
		message1          = ByteResponse{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		message2          = ByteResponse{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	)

	server.SetSessionStartHook(func(session *Session) {
		t.Log("Session start")

		sessionStartCount += 1

		session.Send(message1)
		session.Send(message2)
	})

	server.SetSessionCloseHook(func(session *Session) {
		t.Log("Session close")
		sessionCloseCount += 1
	})

	server.Start()

	client1, err1 := DialTimeout("tcp", addr, PacketN(4, BigEndian), time.Second)
	if err1 != nil {
		t.Fatalf("Create client1 failed, Error = %v", err1)
	}
	client1.SetReadTimeout(time.Second)

	client2, err2 := DialTimeout("tcp", addr, PacketN(4, BigEndian), time.Second)
	if err2 != nil {
		t.Fatalf("Create client2 failed, Error = %v", err2)
	}
	client2.SetReadTimeout(time.Second)

	// client1
	if response, err := client1.Read(); err != nil {
		t.Fatalf("Client1 wait message1 failed, Error = %v", err)
	} else if !bytes.Equal(response, []byte(message1)) {
		t.Fatalf("Client1 message1 not match, Message = %v", response)
	}

	if response, err := client1.Read(); err != nil {
		t.Fatalf("Client1 wait message2 failed, Error = %v", err)
	} else if !bytes.Equal(response, []byte(message2)) {
		t.Fatalf("Client1 message2 not match, Message = %v", response)
	}

	// client2
	if response, err := client2.Read(); err != nil {
		t.Fatalf("Client2 wait message1 failed, Error = %v", err)
	} else if !bytes.Equal(response, []byte(message1)) {
		t.Fatalf("Client2 message1 not match, Message = %v", response)
	}

	if response, err := client2.Read(); err != nil {
		t.Fatalf("Client2 wait message2 failed, Error = %v", err)
	} else if !bytes.Equal(response, []byte(message2)) {
		t.Fatalf("Client2 message2 not match, Message = %v", response)
	}

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
}

func Test_Broadcast(t *testing.T) {
	server, err0 := Serve("tcp", "0.0.0.0:0", PacketN(4, BigEndian))
	if err0 != nil {
		t.Fatalf("Setup server failed, Error = %v", err0)
	}
	addr := server.Addr().String()
	t.Logf("Server: %v", addr)

	var (
		sessionStartCount int
		sessionCloseCount int
		channel           = NewChannel()
		message           = ByteResponse{1, 3, 5, 7, 9, 2, 4, 6, 8, 10}
	)

	server.SetSessionStartHook(func(session *Session) {
		t.Log("Session start")

		channel.Join(session)

		sessionStartCount += 1

		if sessionStartCount == 2 {
			server.Broadcast(channel, message)
		}
	})

	server.SetSessionCloseHook(func(session *Session) {
		t.Log("Session close")
		sessionCloseCount += 1
	})

	server.Start()

	client1, err1 := DialTimeout("tcp", addr, PacketN(4, BigEndian), time.Second)
	if err1 != nil {
		t.Fatalf("Create client1 failed, Error = %v", err1)
	}
	client1.SetReadTimeout(time.Second)

	client2, err2 := DialTimeout("tcp", addr, PacketN(4, BigEndian), time.Second)
	if err2 != nil {
		t.Fatalf("Create client2 failed, Error = %v", err2)
	}
	client2.SetReadTimeout(time.Second)

	// client1
	if response, err := client1.Read(); err != nil {
		t.Fatalf("Client1 wait message failed, Error = %v", err)
	} else if !bytes.Equal(response, []byte(message)) {
		t.Fatalf("Client1 message not match, Message = %v", response)
	}

	// client2
	if response, err := client2.Read(); err != nil {
		t.Fatalf("Client2 wait message failed, Error = %v", err)
	} else if !bytes.Equal(response, []byte(message)) {
		t.Fatalf("Client2 message not match, Message = %v", response)
	}

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
}
