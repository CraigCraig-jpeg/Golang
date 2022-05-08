package main

import (
	"fmt"
)

func main() {
	x := make(chan int)
	go gen(x)
	for i := range x {
		fmt.Println(i)
	}
	fmt.Println("about to exit")
}

func gen(c chan int ) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}
