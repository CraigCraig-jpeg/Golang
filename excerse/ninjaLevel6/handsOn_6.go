package main

import "fmt"

func main() {
	func () {
		fmt.Println("hello")
	}()

	x := func () {
		fmt.Println("hello")
	}

	x()
}