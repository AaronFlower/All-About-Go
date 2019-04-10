package mymath

// FibRecur returns the Fibonacci Number of n
func FibRecur(n int) int {
	if n <= 2 {
		return 1
	}
	return FibRecur(n-1) + FibRecur(n-2)
}

// FibDp returns the Fibonacci Number of n usuing dp
func FibDp(n int) int {
	if n <= 2 {
		return 1
	}
	a := 1
	b := 1
	for i := 0; i < n-1; i++ {
		b = a + b
		a = b - a
	}
	return a
}
