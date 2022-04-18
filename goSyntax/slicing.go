package main 

import "fmt" 

var xy int 
var xz [10] int 
type foo int 
type bar [] int 

func main() {
	x := [] int {1 , 2 , 3}
	fmt.Println(x)
	x = append(x, 77 )
	fmt.Println(x)
	y := []int {123 , 456}
	y = append(x , y...)
	fmt.Println(y)
	y = append(x[2:], x[4:]...)
	fmt.Println(y)
}