package main

import (
	"fmt"
	"net"
)

func main() {
	conn, _ := net.DialTCP("tcp", nil,
		&net.TCPAddr{
			IP:   net.ParseIP("127.0.0.1"),
			Port: 8080})
	var s string
	for {
		fmt.Scanln(&s)
		if s == "q" {
			break
		}
		conn.Write([]byte(s))
	}

	conn.Close()
}
