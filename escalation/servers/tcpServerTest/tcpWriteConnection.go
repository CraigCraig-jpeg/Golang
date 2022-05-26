package main

import (
	// "fmt"
	"io"
	"net"
	"log"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
	panic(err)
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		// go handleConnection(conn)
		io.WriteString(conn, "Connection established\n")
		log.Printf("adress_local: %v, adress_Remote: %v" , conn.LocalAddr(),conn.RemoteAddr())
		conn.Close()
	}
}