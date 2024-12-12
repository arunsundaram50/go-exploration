package main

import (
	"fmt"
)

func f(a string) { // takes a copy of a string
	a = "I am f()"
}

func g(a *string) { // takes an address where string is stored
	*a = "I am g()" // *a means, the value in address a
	fmt.Printf("address of a is %x\n", a)
}

func main() {

	x := "I am main()" // say RAM location = 100,000,000 for example; value is "I am main"
	fmt.Printf("%s\n", x)

	f(x) // A copy of the value "I am main()" is passed
	fmt.Printf("%s\n", x)

	g(&x) // A copy of the address, 100,000,000 is passed
	fmt.Printf("%s\n", x)
}
