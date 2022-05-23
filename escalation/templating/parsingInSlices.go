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
	x := []string{"james", "john", "james"}

	file , err := template.ParseGlob("./parseFiles/*")
	check(err)
	createdFile, err := os.Create("./exampleFiles/parsingInSlices.gohtml")
	err = file.ExecuteTemplate(createdFile, "slices.gohtml", x)
	check(err)
}