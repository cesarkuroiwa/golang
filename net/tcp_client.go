package main

import "net"
import "flag"

func main() {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "localhost:9999")
	tcpConn, _ := net.DialTCP("tcp", nil, tcpAddr)
	flag.Parse()
	tcpConn.Write([]byte(flag.Arg(0)))
}