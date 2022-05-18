package main 

import (
	"log"
)

func main() {
	log.Println(greet("lol"))
}

func greet(s string) string{
	return s 
}