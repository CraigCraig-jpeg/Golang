package main

import (
	"fmt"
)

type human interface{
	speak()
}
type person struct {
	name string
}
func (p person) speak() {
	fmt.Println("i am" , p )
}
func saySomething(h human){
	h.speak()
}
func main() {
	man := &person{
		name: "dave",
	}
	saySomething(*man)
}