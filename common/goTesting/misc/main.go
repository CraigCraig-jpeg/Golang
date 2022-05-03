package main 

import (
	"fmt"
)

func main() {
	x := myAdd(2 , 3 , 4 , 5 , 6 , 7 , 8 , 9 )
	fmt.Println(x)
}

func myAdd(l ...int) int {
	var r int
	for _ , v := range l {
		r += v
	}
	return r
}