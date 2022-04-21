package main

import "fmt"

func main() {
	defer foo("first")
	bar("second")
}

func foo(x string) { 
	fmt.Println(x)
}

func bar(x string) {
	fmt.Println(x)
}
