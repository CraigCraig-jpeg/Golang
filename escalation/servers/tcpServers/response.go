package main

import (
	"net"
	"log"
	"bufio"
	"fmt"
	"strings"
	"text/template"
	// "os"
)

func check(err error) error {
	if err != nil {
		log.Fatalf("error: %v", err)
		return err
	}else{
		return nil
	}

}

func main() {
	listener, err := net.Listen("tcp", ":8081")
	check(err)
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		check(err)
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	// read connection
	request(conn)

	// write connection
	respond(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			// request line
			m := strings.Fields(ln)[0]
			fmt.Println("***METHOD", m)
		}
		if ln == "" {
			// headers are done
			break
		}
		i++
	}
}

func respond(conn net.Conn) {

	// body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Hello World</strong></body></html>`
	file , err := template.ParseFiles("../../templating/parseFiles/normal.gohtml")
	check(err)
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	// fmt.Fprintf(conn, "Content-Length: %d\r\n", len(file.Name))
	// fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\n")
	err = file.ExecuteTemplate(conn , "normal.gohtml" , "Craig")
	// fmt.Fprint(conn, body)
}