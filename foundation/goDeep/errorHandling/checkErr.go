package main

import (
	"fmt"
	"os"
	"strings"
	"io"
)

func main () {
	f , err := os.Create("names.txt")
	if err != nil {
		fmt.Println("Failed to create" , err )
		return
	}
	defer f.Close()
	r := strings.NewReader("lol")

	io.Copy(r, f) 
}