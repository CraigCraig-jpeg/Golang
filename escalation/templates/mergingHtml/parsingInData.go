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

func main() {
	file , err := template.ParseGlob("./exampleFiles/*")
	check(err)
	createdFile, err := os.Create("./exampleFiles/parsingInData.gohtml")
	err = file.ExecuteTemplate(createdFile, "pasing.gohtml", "Craig")
	check(err)
}