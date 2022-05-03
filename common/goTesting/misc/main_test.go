package main

import (
	"testing"
)

/*#########################
#		  myAdd()		  #
#########################*/

func TestMyAdd(t *testing.T) {
	x := myAdd(2 , 3)
	if x != 5 {
		t.Error("expected" , 5 , "got" , x)
	}
}
func TestMyAdd2(t *testing.T) {
	x := myAdd(2 , 3)
	if x != 5 {
		t.Error("expected" , 5 , "got" , x)
	}
}