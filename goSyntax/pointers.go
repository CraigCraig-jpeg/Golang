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

func foo(x *int) {
	*x = 1
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

	zPtr := new(int)
	*zPtr = 20
	fmt.Println(*zPtr)

	rArray := make([]int , 10 , 10)
	var rPtr *int = &rArray[0]
	*rPtr = 21
	fmt.Println(rArray[0])

	hPtr := new(int)
	foo(hPtr)
	fmt.Println(*hPtr)

	var q int
	var qPtr *int 
	qPtr = &q 
	*qPtr = 10
	fmt.Println(*qPtr)

	var f int
	var fPtr *int = &f 
	*fPtr = 20
	fmt.Println(f)

	ePtr := new([10]int)
	mPtr := &ePtr[0]
	*mPtr = 10
	fmt.Println(ePtr)

	lPtr := new([]int)
	*lPtr = append([]int{} ,0)
	fmt.Println(lPtr)
}

