package main 

import (
	"fmt"
	"runtime"
)

func main() {
	counter := 0
	fmt.Println(runtime.NumCPU())

	const gs = 100 

	for i := 0; i < gs; i++ {
		go func(){
			v := counter
			v ++ 
			counter = v 
		}()
		fmt.Println(counter)
	}
}