package main

import (
	"fmt"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	end := make(chan int)

	go send(x chan int, y chan int, z chan int)
}

func send(x, y, z chan<-  int){
	for i := 0; i < 100; i++ {
		if i % 2 == 0 {
			x <- i
		}
		if i % 2 == 1 {
			y <- i
		}
		close(x)
		close(y)
		q <- 0
	}
}