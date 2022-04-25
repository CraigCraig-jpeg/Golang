package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{4, 7, 3, 42, 99, 18, 16, 56, 12}
	sort.Ints(xi)
	xs := []string{"James", "Q", "M", "Moneypenny", "Dr. No"}
	sort.Strings(xs)

	fmt.Println(xi)
	fmt.Println(xs)

}
