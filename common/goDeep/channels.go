package main

import (
	"fmt"
)

// func main() {
// 	c := make(chan int)
// 	go func() {
// 		c <- 1
// 	}()
// 	fmt.Println(<-c)
// }

func main() {
	c := make(chan int , 2)

	c <- 1 
	c <- 3

	fmt.Println(<-c)
	fmt.Println(<-c)
}  