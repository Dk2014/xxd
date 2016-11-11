package net

import "sync"
import "time"
import "bytes"
import "testing"

func Test_Dial(t *testing.T) {
	listener, err1 := Listen("tcp", "0.0.0.0:0", PacketN(4, BigEndian))
	if err1 != nil {
		t.Fatalf("Setup listener failed, Error = %v", err1)
	}
	addr := listener.Addr().String()
	t.Logf("Listen: %v", addr)

	conn, err2 := DialTimeout("tcp", addr, PacketN(4, BigEndian), time.Second)
	if err2 != nil {
		t.Fatalf("Dial fialed, Error = %v", err2)
	}

	if err := conn.Close(); err != nil {
		t.Fatalf("Close conn returns error, Error = %v", err)
	}

	if err := listener.Close(); err != nil {
		t.Fatalf("Close listener returns error, Error = %v", err)
	}
}

func Test_ReadWrite(t *testing.T) {
	listener, err1 := Listen("tcp", "0.0.0.0:0", PacketN(4, BigEndian))
	if err1 != nil {
		t.Fatalf("Setup listener failed, Error = %v", err1)
	}
	addr := listener.Addr().String()
	t.Logf("Listen: %v", addr)

	conn, err2 := DialTimeout("tcp", addr, PacketN(4, BigEndian), time.Second)
	if err2 != nil {
		t.Fatalf("Dial fialed, Error = %v", err2)
	}

	var (
		serverSideErr  error
		serverSideData []byte
		serverSideWait = new(sync.WaitGroup)
		clientSideData = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	)

	serverSideWait.Add(1)
	go func() {
		defer serverSideWait.Done()

		conn, err1 := listener.Accept()
		if err1 != nil {
			serverSideErr = err1
			return
		}

		t.Logf("Accept: %s", conn.RemoteAddr())

		data, err2 := conn.Read()
		if err2 != nil {
			serverSideErr = err2
			return
		}

		serverSideData = data
	}()
	if err := conn.Write(clientSideData); err != nil {
		t.Fatalf("Write data returns error, Error = %v", err)
	}
	serverSideWait.Wait()

	if !bytes.Equal(clientSideData, serverSideData) {
		t.Fatalf("Server side data not match client side data")
	}

	if err := conn.Close(); err != nil {
		t.Fatalf("Close conn returns error, Error = %v", err)
	}

	if err := listener.Close(); err != nil {
		t.Fatalf("Close listener returns error, Error = %v", err)
	}
}
