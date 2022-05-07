package main

import (
	"fmt"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	end := make(chan int)

	go send(even, odd, end)
	go receive(even, odd, end)
}

func receive(even, odd, end <-chan int){
	for {
		select {
		case v :=  <-even:
			fmt.Println(v)
		case v := <-odd:
			fmt.Println(v)
		case v := <-end:
			fmt.Println(v)
			return
		}
	}
}

func send(x, y, z chan<- int){
	for i := 0; i < 100; i++ {
		if i % 2 == 0 {
			x <- i
		}else {
			y <- i
		}
		close(x)
		close(y)
	}
	z <- 0
}