package main

import (
	"bufio"
	"fmt"
	// "io"
	"log"
	"net"
	"os"
)

func Log(s string) {
	log.Println(s)
}

func check (err error) error {
	if err != nil {
		log.Printf("%v", err)
		return err
	}else{
		return err
	}
}

func main() {
	conn , err := net.Dial("tcp","localhost:8081")
	check(err)
	server := bufio.NewScanner(conn)
	input := bufio.NewScanner(os.Stdin)
		
	c := make(chan string)
	d := make(chan string)

	c <- server.Text()
	d <- input.Text()

	for conn != nil {
		select {
		case <- c :
			fmt.Println(server.Text())
		case <- d :
			fmt.Fprintln(conn, input.Text())
		}
	}

	
	defer conn.Close()
}