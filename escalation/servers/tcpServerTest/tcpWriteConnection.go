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
		io.WriteString(conn, "Hi")
		log.Printf("adress: %v, adress_network: %v,adress_string: %v" , ln.Addr(),ln.Addr().Network(),ln.Addr().String() )
		conn.Close()
	}
}