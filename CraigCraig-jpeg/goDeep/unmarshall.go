package main

import "fmt"
import "encoding/json"

type Person struct {
	First string `json:"first"`
	Last string `json:"last"`
	Age int `json:"Age"`
}

func main() {
	p := Person{
		First : "Craig",
		Last : "Samuelson",
		Age : 30,
	}
	v , err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error marshalling")
	}
	
	y := new(Person)
	err = json.Unmarshal(v , y)
	if err != nil {
		fmt.Println("Error unmarshalling")
	}
	fmt.Println(y)

	var z Person
	err = json.Unmarshal(v, &z)
	if err != nil {
		fmt.Println("Error unmarshalling")
	}
	fmt.Println(z.First)

	k :=  `[{"First":"Craig", "Last":"Samu" , "Age":30}]`
	kS := []byte(k)
	
	var l []Person
	lNew := new([]Person)
	
	err = json.Unmarshal(kS , &l)
	if err != nil {
		fmt.Println("Error" , err)
	}
	fmt.Println(l)
	fmt.Println(l[0])
	fmt.Println(l[0].First , l[0].Last , l[0].Age)

	err = json.Unmarshal(kS , &lNew)
	if err != nil {
		fmt.Println("Error" , err)
	}
	fmt.Println(lNew)
	fmt.Println((*lNew)[0].First)
	

}

