package main

import (
	"fmt"
)

func main() {
	fmt.Println(min(1 , 2))
	fmt.Println(miner[int](1 , 2))
	fmt.Println(miner2[int](1 , 2))
}

//generic variadic parameter with constraints
func min[T int](x ...T) int {
	var i int
	for _ , v := range x {
		i += int(v)
	}
	return i
}

//generic parameter with no constraints
func miner[T interface{}](x T , y T) T {

	return x 
}
//generic parameter with constraints
func miner2[T int](x T , y T) T {
	return x + y
}