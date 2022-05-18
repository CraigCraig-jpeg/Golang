package main 

import "fmt"

type person struct {
	first string
}

func changeMe(x *person) {
	x.first = "Craig"
}

func (z *person) changeMe2 (x string) {
	z.first = x
}

func main() {
	x := new(person)
	changeMe(x)
	fmt.Println(x.first)


	y := person{
		first: "john",
	}
	changeMe(&y)
	fmt.Println(y.first)

	y.changeMe2("john")
	fmt.Println(y.first)
}