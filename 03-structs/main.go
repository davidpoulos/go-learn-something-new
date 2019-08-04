package main

import "fmt"

//Human definition
type Human struct {
	age       int
	firstName string
	lastName  string
}

//SetAge changes the Human's age to the age (pointer reference)
func (h *Human) SetAge(newAge int) {
	h.age = newAge
}

//SetAgeByCopy copies the Struct and modifies (When would you use this ?)
func (h Human) SetAgeByCopy(newAge int) {
	h.age = newAge
}

func main() {

	// Initliaze

	// With `new` returns a pointer
	h1 := new(Human)
	h1.firstName = "David"

	fmt.Println(h1)

	// With a `struct literal`
	h2 := Human{
		25,
		"David",
		"Poulos",
	}

	h2.SetAge(26)
	fmt.Println(h2)

	// Var Null/0 Initialize
	var h3 Human

	fmt.Println(h3)

	//With a `struct literal` pointer reference
	h4 := &Human{
		24,
		"David",
		"Poulos",
	}

	fmt.Println(h4)
}
