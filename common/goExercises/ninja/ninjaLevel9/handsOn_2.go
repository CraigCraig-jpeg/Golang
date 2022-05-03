package main

import (
	"fmt"
)

type person struct {
	
}
type human interface{
	speak()
}
func (p *person) speak() {
	fmt.Println("i am" , p )
}
func saySomething(h human){
	h.speak()
}
func main() {
	man := person{
	}
	saySomething(&man)
}