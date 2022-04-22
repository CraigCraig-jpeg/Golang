package main

import "fmt"

func main(){
	x := foo(1)
	y , z := bar(1 , "lorem ipsum dolor sit amet")
	fmt.Println(x,y,z)
}

func foo(x int ) int {
	return x
}

func bar(x int , y string)(int , string){
	return x , y
} 