package main 

import (
	"fmt"
)

func main(){
	x := factorail(4)
	fmt.Println(x)
}

func factorail(n int) int {
	if n == 0{
		return 1
	}
	return n * factorail(n-1)
}