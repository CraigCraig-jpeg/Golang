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
	d := make(chan <- int , 2)
	// e := make(<- chan int , 2)

	d <- 3
	d <- 4 

	c <- 1 
	c <- 3

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Printf("%T" ,c )
}  