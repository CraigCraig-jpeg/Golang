package main

import "fmt"

func main() {
	s := howl(summer , 1 , 3 , 4)
	fmt.Println(s)
}

func summer(x []int) []int{
	x = append([]int{0}, x...)
	return x
}

func howl(t func (x []int) []int , y ...int) int {
	var sum []int
	var h int
	r := t(y)
	sum = append([]int{}, r ...)
	for _ , i := range sum {
		h += i
	}
	return h
}
