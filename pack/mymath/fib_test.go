package mymath

import "testing"

func TestFibRecur(t *testing.T) {
	num := []int{1, 2, 3, 4, 5, 6, 7}
	data := []int{1, 1, 2, 3, 5, 8, 13}

	for i, v := range num {
		got := FibRecur(v)
		if got != data[i] {
			t.Errorf("FibRecur(%d) expects %d, but got %d", v, data[i], got)
		}
	}
}

func TestFibDp(t *testing.T) {
	num := []int{1, 2, 3, 4, 5, 6, 7}
	data := []int{1, 1, 2, 3, 5, 8, 13}

	for i, v := range num {
		got := FibDp(v)
		if got != data[i] {
			t.Errorf("FibRecur(%d) expects %d, but got %d", v, data[i], got)
		}
	}
}

func BenchmarkFibRecur(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibRecur(20)
	}
}

func BenchmarkFibDp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibDp(20)
	}
}
