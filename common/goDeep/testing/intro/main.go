package goTest_1

import (
	"log"
)

func main() {
	log.Println(mySum(1 , 2))
}

func mySum(x ...int) int{
	var sum int
	for _ , r := range x {
		sum += r
	}
	return sum
}