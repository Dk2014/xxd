package net

import "testing"

func Test_Listener(t *testing.T) {
	listener, err1 := Listen("tcp", "0.0.0.0:0", PacketN(4, BigEndian))
	if err1 != nil {
		t.Fatalf("Setup listener failed, Error = %v", err1)
	}
	addr := listener.Addr().String()
	t.Logf("Listen: %v", addr)

	if err2 := listener.Close(); err2 != nil {
		t.Fatalf("Close listener retuens error, Error = %v", err2)
	}
}
