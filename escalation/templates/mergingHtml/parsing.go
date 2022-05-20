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
	nf , err := os.Create("./exampleFiles/parsing.gothml")
	if err != nil {
		panic(err)
	}
	tpl , err = tpl.ParseFiles("./exampleFiles/pasing.gohtml" , "./exampleFiles/osMerge.html")
	// err = tpl.Execute(os.Stdout , nil)
	// err = tpl.Execute(nf , nil)
	// err = tpl.ExecuteTemplate(os.Stdout , "pasing.gohtml" , nil)
	tpl , err = tpl.ParseGlob("./exampleFiles/*.gohtml")
	err = tpl.ExecuteTemplate(os.Stdout , "pasing.gohtml" , nil)
	if err != nil {
		log.Fatal(err)
	}
	rar := template.Must(template.ParseFiles("./exampleFiles/pasing.gohtml"))
	rar.Execute(os.Stdout , nil)
	defer nf.Close()
}