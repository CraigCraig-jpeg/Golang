package main 

import (
	"fmt"
)

func main(){
	y := []int {1 , 2 , 3}
	x := sum(y...)
	fmt.Println(x)

	z := even(sum , y...)
	fmt.Println(z)

	r := odd(sum , y...)
	fmt.Println(r)
}

func sum(x ...int) int {
	total := 0
	for _ , v := range x{
		total += v
	}
	return total
}

func even(t func(x ...int) int, i ...int) int{
	var ji []int 
	for _ , g := range i {
		if g % 2 == 0{
			ji = append(ji , g)
		}
	}
	return t(ji...)
}

func odd(t func(x ...int) int, i ...int) int{
	var ji []int 
	for _ , r := range i {
		if r % 2 == 1 {
			ji = append(ji , r)
		}
	}
	return t(ji...)
}