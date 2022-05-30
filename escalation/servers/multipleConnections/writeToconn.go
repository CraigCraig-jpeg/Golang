package main

import (
	"net"
	"io"
	"fmt"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	conn , err := listener.Accept()
	response := fmt.Sprintf("Hi %v \n" , conn.LocalAddr())
	io.WriteString(conn, response)
	conn.Close()
}
