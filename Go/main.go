package main 

import "fmt" 

var xy int 
var xz [10] int 
type foo int 
type bar [] int 

func main() {
	x := [] int {1 , 2 , 3}
	fmt.Println(x[2:])
	for i, v := range x {
		fmt.Println(i , v)
	}
}