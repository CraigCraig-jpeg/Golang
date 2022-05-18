package dog

import (
	"testing"
	// "log"
)

func TestYears(t *testing.T) {
	got := 2
	desired := 14

	x := Years(got)
	if x != desired {
		t.Errorf("Unexpected year: %vv, expected %v", x, desired)
	}
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(i)
	}
}

func TestYearsTwo(t *testing.T) {
	got := 3
	expected := 21
	test := YearsTwo(got)
	if test != expected {
		t.Errorf("Unexpected year: %v, expected %v",got , expected)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(i)
	}
}