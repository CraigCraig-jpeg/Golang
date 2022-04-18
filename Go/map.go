package main 

import "fmt" 

func main() {
	x := make ([]int,10 , 100 )
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	foo := map[string]int{
		"James":12,
		"money":32,
	}
	fmt.Println(foo["James"])
	foo["cliffy"] = 32
	fmt.Println(foo["cliffy"])
	delete(foo, "cliffy")
	fmt.Println(foo["cliffy"])
}