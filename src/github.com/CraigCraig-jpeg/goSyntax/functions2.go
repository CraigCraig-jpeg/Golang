package main

import "fmt"

type person struct {
	first string
	last string
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) name() () {
	fmt.Println(s.first , s.last)
}

func main() {
	x := secretAgent {
		person : person {
			first : "craig",
			last : "samuelson",
		},
		ltk : true,
	}
	x.name()
	fmt.Println("cool man" , x)
}
