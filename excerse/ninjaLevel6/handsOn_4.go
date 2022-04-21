package main

import "fmt"

type person struct {
	first string
	last string
	age int
}

func (p person) speak() (){
	fmt.Printf("%v%v%v",p.first,p.last,p.age)
}

func main() {
	x := person{
		first: "Craig",
		last: "Samuelson",
		age: 10,
	}
	x.speak()
}