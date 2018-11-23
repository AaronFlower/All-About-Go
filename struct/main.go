package main

import "fmt"

// Human beging
type Human struct {
	name  string
	age   int
	phone string
}

// Student for test
type Student struct {
	Human  // Anonymous
	school string
}

// Employee for test
type Employee struct {
	Human
	company string
}

// SayHi for Human
func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s \n", h.name, h.phone)
}

// SayHi overrides the Human method
func (e *Employee) SayHi() {
	fmt.Printf("Hi, I am %s , I work at %s. Call me on %s \n", e.name, e.company, e.phone)
}

func main() {
	mark := Student{Human{"Mark", 25, "222-2323"}, "MIT"}
	sam := Employee{Human{"Sam", 45, "444-4343"}, "Google"}

	mark.SayHi()
	sam.SayHi()
}
