package main

import "fmt"

func main() {
	z := total(tip , 64 )
	fmt.Println(z)

	
}

func tip(x float64) float64 {
	y := x + x * 0.10
	return y
}

func total(x func (x float64) float64 , total float64) float64{
	y := x(10)
	return y 
}