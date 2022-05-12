package main 

import (
	"log"
	"errors"
	"fmt"
)

func main() {
	log.Println("hi")
	err := errors.New("hi")
	err2 := fmt.Errorf("error2: [%w]", err)
	fmt.Println(err)
	fmt.Println(errors.Unwrap(err2))

	value := errors.New("the value should be good(positive)")
	log.Println(value.Error())
}