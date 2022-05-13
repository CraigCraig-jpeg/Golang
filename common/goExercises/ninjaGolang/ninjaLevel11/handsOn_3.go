package main

import (
	"log"
)

type customError struct {
	errors string
}

func (e *customError) Error() string {
	return e.errors
}

func (a *customError)foo(s string) string {
	a.errors = s
	return a.Error()
}

func main() {
	x := new(customError)
	err := x.foo("rawrrrr")
	log.Println(err)
}