package main

import "fmt"
import "encoding/json"

type Message struct {
	Name string 
	Body string
	Time int
}

func main() {
	m := Message{
		Name : "jamse",
		Body : "hi",
		Time : 1800,
	}
	b , err := json.Marshal(m)
	if err != nil {
		fmt.Println("fail" , err)
	}

	var s Message
	json.Unmarshal(b, &s)
	fmt.Println(s)

	type k struct {
		Name string
		Body string
		Time int
	}
	K := new(k)
	json.Unmarshal(b , K)
	if err != nil {
		fmt.Println("fail" , err)
	}
	fmt.Println(K)
}