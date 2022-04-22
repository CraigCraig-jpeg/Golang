package main

import "fmt"

type person struct {
	name string
}

func run(y int) int{
	var x int
	x = y
	return x
}

func main() {
	a := 42 
	var aPtr *int = &a
	fmt.Println(aPtr)
	fmt.Printf( "%T\n" , aPtr)
	fmt.Println(* aPtr)
	*aPtr = 43
	fmt.Println(* aPtr)

	x := new(person)
	*x = person{name : "david"}
	fmt.Println(x.name)

	// y := make([]int , 10 , 10)
	// // *y = 10
	// fmt.Println(x.name)

	zPtr := new(int)
	*zPtr = 20
	fmt.Println(*zPtr)

	rArray := make([]int , 10 , 10)
	var rPtr *int = &rArray[0]
	*rPtr = 21
	fmt.Println(rArray[0])
}