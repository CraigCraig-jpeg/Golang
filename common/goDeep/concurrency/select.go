package main

import (
	"fmt"
)

func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	//send a message
	go send()
}

func send(a chan int, b chan int , c chan int) {
	
}