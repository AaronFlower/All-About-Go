package mymath

import "errors"

// Division returns the quotient of a / b
func Division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("can't be divided by zero")
	}
	return a / b, nil
}
