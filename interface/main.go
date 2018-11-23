package main

import "fmt"
import "reflect"

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
	loan   float32
}

// Employee for test
type Employee struct {
	Human
	company string
	salary  float32
}

// SayHi for Human
// to implements interface, the receiver can not be a pointer.
func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s \n", h.name, h.phone)
}

// Sing implements the Men mehtod.
func (h Human) Sing(lyrics string) {
	fmt.Println("La la la ...", lyrics)
}

// SayHi overrides the Human method
func (e *Employee) SayHi() {
	fmt.Printf("Hi, I am %s , I work at %s. Call me on %s \n", e.name, e.company, e.phone)
}

// Men interface for Human, Student and Employee
type Men interface {
	SayHi()
	Sing(lyrics string)
}

func main() {
	mark := Student{Human{"Mark", 25, "222-2323"}, "MIT", 0.00}
	paul := Student{Human{"Mark", 25, "222-2323"}, "MIT", 0.00}
	sam := Employee{Human{"Sam", 45, "444-4343"}, "Google", 1000}
	tom := Employee{Human{"Tom", 25, "222-2323"}, "MIT", 5.00}

	var i Men

	i = &mark
	fmt.Println("This is Mike , a Student:")
	i.SayHi()
	i.Sing("Friend \n")

	i = &tom
	fmt.Println("This is Tom , a Student:")
	i.SayHi()
	i.Sing("Beat It \n")

	// Declare a slice Men
	fmt.Println("Let's use a slice of Men and see what happens")
	x := make([]Men, 3)
	x[0], x[1], x[2] = &paul, &sam, &mark

	for _, value := range x {
		value.SayHi()
	}

	// reflection
	var n = 3.4
	v := reflect.ValueOf(n)
	fmt.Println("Type:", v.Type())
	fmt.Println("value:", v.Float())
	fmt.Println("kind is float64:", v.Kind(), v.Kind() == reflect.Float64)
}
