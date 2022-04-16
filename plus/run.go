package main 

import "fmt"

type man [5] int

func main() {
	var foo man
	foo[0] = 1
	foo[1] = len(foo)
	var x [10]int
	x[0] = 2 
	fmt.Println(x[0])
	fmt.Println(foo[1])

}