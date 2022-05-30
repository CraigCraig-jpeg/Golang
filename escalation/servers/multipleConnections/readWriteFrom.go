package main

import (
	"log"
	"net"
	"bufio"
	"time"
	// "io"
	"fmt"
)

func check(err error) error {
	if err != nil {
		// log.Panic(err)
		// log.Fatal(err)
		log.Println(err)
		return err
	}else {
		return nil
	}
}

func main() {
	c := make(chan string)
	listener, err := net.Listen("tcp", ":8080")
	defer listener.Close()
	check(err)
	for {
		conn, err := listener.Accept()
		check(err)
		go handleConnection(conn, c)
	}
}

func handleConnection(conn net.Conn, c chan string){
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	check(err)
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	defer conn.Close()
	defer fmt.Println("deadline exceeded for connection", conn.RemoteAddr())
}