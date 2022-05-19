package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl , err := template.ParseFiles("./exampleFiles/pasing.gohtml")
	if err != nil {
		panic(err)
	}
	err = tpl.Execute(os.Stdout , nil)
	if err != nil {
		log.Fatal(err)
	}
}