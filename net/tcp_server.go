package main

import "net"
import "fmt"

func main() {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "localhost:9999")
	tcpListener, _ := net.ListenTCP("tcp", tcpAddr)
	for {
		tcpConn, _ := tcpListener.AcceptTCP()
		buffer := make([]byte, 1024)
		tcpConn.Read(buffer)
		fmt.Println(string(buffer))
	}
}