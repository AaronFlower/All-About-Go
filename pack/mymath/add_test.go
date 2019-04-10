package mymath

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	got := Add(1, 2)
	if 3 != got {
		t.Errorf("Add(1, 2) = %d; want 3", got)
	}
}

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("Hello")
	}
}
