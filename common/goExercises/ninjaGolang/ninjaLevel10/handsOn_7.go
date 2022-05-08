package main

import (
	"fmt"
)
const routines int = 10
func main() {
	r := make(chan int)
	for i := 0; i < routines; i++ {
		go func(c chan int){
			for j := 0; j < routines; j++ {
				c<-j
			}
		}(r)
	}


	// close(r)
	for h := range r {
		fmt.Println(h)
	}
	fmt.Println("end of routines")
}