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
	
	// for h := range r {
		// 	fmt.Println(h)
		// }
		
	for i := 0; i < 100; i++ {
		fmt.Println(<-r)
	}
	close(r)
	fmt.Println("end of routines")
}