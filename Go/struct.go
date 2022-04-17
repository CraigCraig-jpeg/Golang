package main

import "fmt"

type foo struct{
	 name string
	 middle string
	 surname string
}

func main(){
	i := foo { 
		name : "Craig",
		middle : "",
		surname : "Samuelson",
	}
	j := foo { 
		name : "James",
		middle : "jule",
		surname : "Howard",
	}
	fmt.Println(i.name , i.surname)
	fmt.Println(j.name , j.middle , j.surname)
}