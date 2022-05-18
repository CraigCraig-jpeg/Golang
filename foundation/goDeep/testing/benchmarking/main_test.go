package main 

import (
	// "log"
	"testing"
)

func TestGreet(t *testing.T) {
	got := greet("Craig")
	expected := "Craig"

	if got != expected {
		t.Errorf("Function test failed , got %v, want %v", got , expected)
	}
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		greet("james")
	}
}