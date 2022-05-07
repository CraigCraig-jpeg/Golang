package main

import (
	"fmt"
)

func main() {

}

//parameter with constraints
func min(x , y int) int {
	if x < y {
		return x
	}
	return y
}