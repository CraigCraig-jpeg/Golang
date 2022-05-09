package main 

import (
	// "fmt"
	"os"
	"log"
)

func main() {
	err := os.Mkdir("log" , os.ModePerm)
	if err != nil {
		log.Print("Cloud not crete logging dir")
	}
	f , err := os.Create("./log/random.txt")
	defer f.Close()
	if err != nil {
		log.Println(err)
	}
	log.SetOutput(f)

	_, err = os.Open("randosm.txt")
	if err != nil {
		// fmt.Println(err)
		log.Println(err)
	}

}