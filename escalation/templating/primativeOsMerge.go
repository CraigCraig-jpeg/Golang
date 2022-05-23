package main

import (
	"os"
	"fmt"
	"log"
)

func main() {
	name := os.Args[1]
	surname := os.Args[2]
	htm := fmt.Sprintf(`
	<!DOCTYPE html>
	<html>
	<body>
	<h1>` + name + surname +`</h1> 
	</body>
	</html>
	`)

	file , err := os.Create("./files/osMerge.html")
	if err != nil {
		panic(err)
	}
	n , err := file.WriteString(htm)
	log.Println(n)
	defer file.Close()
}