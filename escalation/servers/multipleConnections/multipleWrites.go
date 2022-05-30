package main

import (
	"net"
	"io"
	"fmt"
	"bufio"
	"os"
)

func main() {
	listener, err := net.Listen("tcp",":8080")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn){
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	input := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		io.WriteString(conn,input.Text())
	}
}