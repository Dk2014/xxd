package main

import (
	"fmt"
	"net"
	"os"
)

const (
	API_SERVER = "127.0.0.1"
	TCP_PORT   = 6668
)

func main() {
	socket, err := net.ListenTCP("tcp4", &net.TCPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: TCP_PORT,
	})

	if err != nil {
		fmt.Println("Listen失败", err)
		return
	}

	defer socket.Close()

	for {
		conn, err := socket.Accept()

		if err != nil {
			fmt.Println("服务器连接失败", err)
			continue
		}

		go readdata(conn)
	}
}

func readdata(conn net.Conn) {
	connFrom := conn.RemoteAddr().String()
	println("Connection from: ", connFrom)
	for {
		data := make([]byte, 2048)
		length, err := conn.Read(data[0:2048])
		switch err {
		case nil:
			handleMsg(length, err, data)

		default:
			goto DISCONNECT
		}
	}
DISCONNECT:
	err := conn.Close()
	println("Closed connection:", connFrom)
	checkError(err, "Close:")
}

func handleMsg(length int, err error, msg []byte) {
	if length > 0 {
		for i := 0; ; i++ {
			if msg[i] == 0 {
				break
			}
		}
		fmt.Printf("Received data: %v", string(msg[0:length]))
		fmt.Println("   length:", length)
		f, _ := os.OpenFile("/Users/tangweichen/log.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
		defer f.Close()
		f.Write(msg[:length])
		f.WriteString("\n")
	}

}

func checkError(error error, info string) {
	if error != nil {
		panic("ERROR: " + info + " " + error.Error()) // terminate program
	}
}
