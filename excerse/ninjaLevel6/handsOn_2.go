package main 

import "fmt"

func main() {
	x := []int {1 , 2 , 3 , 4 , 5 , 6 ,}
	y := foo(x...)
	fmt.Println(y)
}

func foo(x ...int) int{
	var y int
	for _ , i := range x {
		y += i
	}
	return y 
}