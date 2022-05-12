package main

import (
	"log"
)

type max struct {
	name string
	age int64
}

func main() {
	x := fullname()
	log.Println(*&x.age)
	log.Println(*&x.name)
}

func fullname() *max {
	return &max{name: "Alex", age:10}
}