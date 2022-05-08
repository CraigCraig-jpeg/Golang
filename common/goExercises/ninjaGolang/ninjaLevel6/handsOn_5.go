package main

import "fmt"

type square struct {
	side float64
}
func (x square) area() float64 { 
	area := x.side * x.side
	return area
}

//----------------------------------------------------------------//

type circle struct {
	radius float64
}
func (x circle) area() float64 {
	pie := 3.141592653589793
	area := pie * x.radius * x.radius
	return area
}

//----------------------------------------------------------------//

type shape interface {
	area() float64
}

func area(x shape) {
	fmt.Println("shape area is" , x.area())
}

func main() {
	Square := square {
		side: 2,
	}
	Circle := circle {
		radius:2,
	}
	area(Square)
	area(Circle)
}