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

type human interface {
	speak()
}

func foo(h human){
	switch h.(type) {
	case person:
	fmt.Println("called person with" , h.(person).first )
	case secretAgent:
	fmt.Println("called agent with" , h.(secretAgent).first )
	}
}

func (i secretAgent) speak(){
	fmt.Println("i am agent")
}

func (i person) speak(){
	fmt.Println("i am person")
}

func main(){
	x := secretAgent {
		person : person {
			first : "Craig", 
			last : "Samuelson",
		},
		ltk : true,
	}
	y := person {
		first : "Dr.",
		last : "No",
	}
	func(x secretAgent, y person){
		x.speak()
		y.speak()
		foo(y)
		foo(x)
	}(x , y)
	z := func (x int){
		fmt.Println(x)
	}
	z(2)
	i := fuss("Hi")
	fmt.Println(i)

	g := jull()
	fmt.Printf("%T" , g)
}

func fuss(x string) (string){
	return x
}

func jull() (func() int) {
	return func() int {
		return 78
	}
}