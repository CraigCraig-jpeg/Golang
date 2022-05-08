package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go writee(c)
	for i := range c {
		fmt.Println(i)
	}
}

func writee(c chan int) {
	for i := 0 ; i < 100 ; i++{
		c <- i
	}	
	close(c)
}