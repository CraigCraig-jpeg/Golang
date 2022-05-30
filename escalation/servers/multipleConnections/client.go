package main

import (
	"bufio"
	"fmt"
	// "io"
	"log"
	"net"
	"os"
)

func check (err error) error {
	if err != nil {
		log.Printf("%v", err)
		return err
	}else{
		return err
	}
}

func main() {
	conn , err := net.Dial("tcp","localhost:8080")
	check(err)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Fprintln(conn ,input.Text())
	}
	defer conn.Close()
}