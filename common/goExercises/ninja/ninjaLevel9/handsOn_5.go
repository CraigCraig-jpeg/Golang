package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Your operating system" , runtime.GOOS)
	fmt.Println("Your operating architecture" , runtime.GOARCH)
}