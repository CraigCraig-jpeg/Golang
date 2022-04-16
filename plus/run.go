package main 

import "fmt"

const y int = 2
var z [10]int 
type foo [] int

func main() {
	x := []int{1,2,3}
	for ie , edd := range x {
		fmt.Println(ie , edd)
	}
}