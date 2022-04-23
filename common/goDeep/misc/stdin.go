package main

import "fmt"
import "os"
// import "io"
// import "strings"
// import "syscall"
import "bufio"

func main() {
	// io.WriteString(os.Stdout, "man")
	reader := bufio.NewScanner(os.Stdin)
	reader.Scan()
	text := reader.Text()
	fmt.Println(text)

}