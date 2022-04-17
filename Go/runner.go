package main 

import "fmt"

func main() {
	var z []int 
	z = append (z , 1 , 3, 4 , 5)
	fmt.Println(z)

	x := make([]int ,5 , 5 )
	x[0] = 12 
	for i , v := range x {
		fmt.Println(i , v )
	}
	fmt.Printf("%T\n" , x )

	y := [10]int{1 , 2 , 3 , 4 , 5 , 6 , 2 , 4 , 5 , 2 }
	for i , v := range y {
		fmt.Println(i , v)
	}
	fmt.Printf("%T\n" , y )

	b := append(y[5:6] , y[6:7]...)
	c := append(y[5:6] , y[6:7]...)

	fmt.Println(b)
	fmt.Println(c)

	t := []int{1 , 2 , 3 , 4 , 5 , 6 , 2 , 4 , 5 , 2 }
	g := []int{43 , 42 , 55 , 56}
	g = append(g , 52 , 53 , 54 , 55 )
	g = append(g, t ...)
	fmt.Println(g)
	g = append(g[:5] , g[7:]...)
	fmt.Println(g)
	fmt.Println(t)

	f := make([]string , 5 , 5)
	f = append(f , "five")
	fmt.Println(f)
	fmt.Println(cap(f))
	fmt.Println(len(f))
	
	

}