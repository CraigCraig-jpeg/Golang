package main

import "fmt"

func main(){
	foo()
	bar("hello")
	str := woo("perfection")
	fmt.Println(str)
	x , y := tang("mike" , "hawk")
	fmt.Println(x)
	fmt.Println(y)
}

func foo(){
	fmt.Println("hello")
}
func bar(s string){
	fmt.Println(s)
}
func woo(s string) string{
	return fmt.Sprint(s)
}
func tang(str string , ln string) (string, bool){
	a := fmt.Sprint(str , ln)
	b := false
	return a , b
}