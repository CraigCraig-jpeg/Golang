package main

import (
	"log"
	"os"
	"text/template"
)

func check(err error) any{
	if err != nil {
		panic(err)
	}else{
		return nil
	}
}

type person struct {
	Name string
	Age int
}

func main() {
	craig := person{
		Name: "Craig",
		Age: 10,
	}
	log.Println("created struct")
	file , err := template.ParseGlob("./parseFiles/*")
	check(err)
	createdFile, err := os.Create("./exampleFiles/parsingInStructs.gohtml")
	err = file.ExecuteTemplate(createdFile, "structs.gohtml", craig)
	check(err)
}