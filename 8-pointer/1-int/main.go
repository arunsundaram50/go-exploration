package main

import (
	"fmt"
)

func f(a int) { // takes a int value, local to this function
	a = 100
}

func g(a *int) { // takes an address where int is stored
	*a = 100 // *a means write the value 100 at the address pointed to by a (100,000,000)
}

func main() {

	x := 10 // Say RAM location = 100,000,000
	fmt.Printf("%d\n", x)

	f(x) // 10 is passed. i.e. a copy of the value stored in x is passed
	fmt.Printf("%d\n", x)

	g(&x) // 100,000,000 is passed: i.e. a copy of the address where x is stored is passed
	fmt.Printf("%d\n", x)
}
