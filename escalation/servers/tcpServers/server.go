package main

import (
	"fmt"
	"net"
	"os"
	"bufio"
)

func check(err error) error {
	if err != nil {
		panic(err)
		return err
	}else{
		return nil
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8081")
	check(err)
	for {
		conn, err := listener.Accept()
		check(err)
		handle(conn)
	}
}

func handle(conn net.Conn){
	fmt.Fprintln(conn , "Hi")
	output := bufio.NewScanner(conn)
	client := bufio.NewScanner(os.Stdin)
	fmt.Println(output.Text())
	fmt.Fprintln(conn, client.Text())

	defer conn.Close()
}