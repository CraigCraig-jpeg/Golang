package main 

import (
	"bufio"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		log.Println(scanner.Text())
	}
}