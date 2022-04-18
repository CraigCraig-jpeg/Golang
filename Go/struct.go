package main

import "fmt"

type person struct{
	name string
	middle string
	surname string
}
type secretAgent struct{
	person
	ltk bool
}

func main(){
	i := person { 
		name : "Craig",
		middle : "",
		surname : "Samuelson",
	}
	j := person { 
		name : "James",
		middle : "jule",
		surname : "Howard",
	}
	k := secretAgent {
		person : person{
		name : "Craig",
		middle : "",
		surname : "Samuelson",
		},
		ltk : true, 
	}
	l := struct {
		name string 
		age int
	}{
		name : "Craig",
	}
	l.age = 129

	o := map[string]secretAgent{
		k.name : k ,
		k.middle : k ,
		k.surname : k ,
	}

	fmt.Println(i.name , i.surname)
	fmt.Println(j.name , j.middle , j.surname)
	fmt.Println(k.name , k.surname , k.ltk)
	fmt.Println(l.name , l.age )

	for _ , v := range o {
		fmt.Println(v)
	}
}