package mymath

import "testing"

func Benchmark_Division(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Division(4, 5)
	}
}

func Benchmark_TimeConsumingFunction(b *testing.B) {
	b.StopTimer() // To stop time counting.
	// so we can handle file, db access some operations.

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Division(4, 5)
	}
}
