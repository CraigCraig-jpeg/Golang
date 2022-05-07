package main

import (
	"fmt"
)

type hotdog int 
var g hotdog

func main() {
	fmt.Println(min(1 , 2))
	fmt.Println(miner[int](1 , 2))
	fmt.Println(miner2[int](1 , 2))

	g = 20 
	fmt.Println(miner2(1 , g))
	fmt.Println(miner3(1 , 7))
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
//generic parameter with constraints and support for undelying 
// types 
func miner2[T ~int](x T , y T) T {
	return x + y
}
//generic parameter with variadic constraints 
type variadic interface {
	float64 | int 
}
func miner3[T variadic](x T , y T) T {
	return x + y
}