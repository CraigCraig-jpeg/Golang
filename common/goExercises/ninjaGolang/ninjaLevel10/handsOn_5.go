package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	// add(c)
	go func() {
		c<-2
	}()
	
	v, ok := <- c
	if !ok {
		}else {	
			fmt.Println(v, ok)
	}
	
	close(c)
	
	v, ok = <- c
	if !ok {
		fmt.Println(v, ok)
	}else {	

	}
}

func add(c chan<- int) {
	go func (){
		for i := 0; i < 100 ; i++ {
			c <- i
		}
	}()
}