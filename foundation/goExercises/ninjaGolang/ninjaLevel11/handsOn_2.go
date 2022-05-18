package main

import (
	"encoding/json"
	"fmt"
	"log"
	"errors"
)


type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}
	
	bs, err := toJSON(p1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bs))
	
}

// toJSON needs to return an error also
func toJSON(a interface{}) ([]byte , interface{} ){

	marshallFail := errors.New("Marhsalling failed due to malformed JSON")
	bs, err := json.Marshal(a)
	if err != nil {
		return nil , marshallFail.Error()
	}
	return bs , nil
}