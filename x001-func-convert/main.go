package main

import (
	"errors"
	"fmt"
)

type binFunc func(int, int) (int, error)

func (f binFunc) Error() string {
	return "binFunc error"
}

func divide(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("Can't be divided by zero")
	}
	return x / y, nil
}

func main() {
	var bf = binFunc(divide)
	o, err := bf(3, 0)
	if err != nil {
		fmt.Println(bf)
		fmt.Println(err)
	} else {
		fmt.Println("The result is :", o)
	}
}
