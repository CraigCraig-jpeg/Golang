package main 

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := []byte("MyDarkSecret")
	passwordEncrypt , err := bcrypt.GenerateFromPassword(password , 4)
	if err != nil {
        panic(err)
    }

	password1 := []byte("MyDarkSecret1")
	passwordEncrypt1 , err := bcrypt.GenerateFromPassword(password1 , 4)
	if err != nil {
        panic(err)
    }
	
	fmt.Println("Comparing ...")

	err = bcrypt.CompareHashAndPassword(passwordEncrypt, password)
	if err != nil {
        panic(err)
    }else{
		fmt.Println("Success, youre in !")
	}
}