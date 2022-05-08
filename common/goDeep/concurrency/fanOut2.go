package main

import (
	"fmt"
	"sync"
)

func main() {
	a := make(chan int)
	b := make(chan int)
	go generate(a , 1 , 3 , 4 , 5 , 6 , 7 , 8 , 9)
	go spreadIn(a , b)
	for i := range b {
		fmt.Println(i)
	}
}

func generate(c chan<- int, n ...int) {
	for _ , i := range n {
		c <- i
	}
	close(c)
}

func spreadIn(a <-chan int , b chan<- int){
	var g sync.WaitGroup
	for i := range a {
		g.Add(1)
		go func(r int) {
			b <- r
			g.Done()
		}(i)
	}
	g.Wait()
	close(b)
}