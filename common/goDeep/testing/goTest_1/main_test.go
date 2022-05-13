package goTest_1

import (
	"testing"
	"log"
)

func TestMySum(t *testing.T) {
	log.Println("Test on my Sum")
	value , want := 2 , 2
	if want != mySum(value) {
		t.Error(" Error: expected 2 , got" , value)
	}
}