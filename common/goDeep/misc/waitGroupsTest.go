package main 

import (
	"sync"
	"fmt"
)

var wait sync.WaitGroup

func main() {
	wait.Add(2)

	go run()
	go run()

	wait.Wait()
}

func run() {
	fmt.Println("hello world")
	wait.Done()
}