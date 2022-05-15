package goTest_1

import (
	"testing"
	// "log"
)

func TestMySum(t *testing.T) {
	var err bool = false
	
	type test struct {
		data []int 
		answer int
	}

	tests := []test{
		test{data: []int{1 , 2 , 3} , answer: 6},
		test{data: []int{3, 4, 5, 6} , answer: 7},
		test{data: []int{4 ,4 } , answer: 8},
		test{data: []int{1 , 2 , 3 , 4 , 5} , answer: 9},
		test{data: []int{9 , 2 ,3} , answer: 10},
		test{data: []int{10 , 1} , answer: 11},
	}
	
	for _, r := range tests {
		var sum int
		var index int
		for i , rs := range r.data {
			sum += rs 
			index = i
		}
		if sum != r.answer{
			t.Log("expected" , r.answer , "got" , sum , "at index" , index )
			err = true
		}
	}
	if err == true{
		t.Error("One or more tests have failed at TestMySum")	
	}
}