package main

import (
	"net"
	"log"
	"bufio"
)

func check(err error) error {
	if err != nil {
		panic(err)
	}else{
		return nil
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	check(err)
	for {
		conn, err := listener.Accept()
		check(err)
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn){
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		out := scanner.Text()
		log.Println(out)
	}
	defer conn.Close()
}