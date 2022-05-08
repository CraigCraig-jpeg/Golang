package main 

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("emit macho dwarf: elf header corrupted")
	err2 := errors.New("mmm yes: elf header corrupted")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(err2)
}