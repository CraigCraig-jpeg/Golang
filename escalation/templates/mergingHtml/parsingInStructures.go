package main

import (
	// "log"
	"os"
	"text/template"
)

func check(err error) any{
	if err != nil {
		return err
	}else{
		return nil
	}
}

type person struct {
	name string
	age int
}

func main() {
	craig := person{
		name: "Craig",
		age: 10,
	}
	file , err := template.ParseGlob("./exampleFiles/*")
	check(err)
	createdFile, err := os.Create("./exampleFiles/parsingInData.gohtml")
	err = file.ExecuteTemplate(createdFile, "pasing.gohtml", craig)
	check(err)
}