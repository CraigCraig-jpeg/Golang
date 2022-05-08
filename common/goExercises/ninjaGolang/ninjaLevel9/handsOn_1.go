package main 

import (
	"fmt"
	"sync"
)

var in interface{
	printer()
}

var wg sync.WaitGroup
func main() {
	wg.Add(2)
	go printer()
	go printer()
	wg.Wait()
}

func printer() {
	fmt.Println("Hello, world!")
	wg.Done()
}