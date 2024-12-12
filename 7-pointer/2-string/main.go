package main

import (
	"fmt"
)

func f(a string) { // take a string
	a = "I am f()"
}

func g(a *string) { // take an string address
	*a = "I am g()" // *a means, the value in address a
	fmt.Printf("address of a is %x\n", a)
}

func main() {

	x := "I am main()" // RAM location = 100,000,000 for example; value is "I am main"
	fmt.Printf("%s\n", x)

	f(x) // "I am main()"
	fmt.Printf("%s\n", x)

	g(&x) // 100,000,000
	fmt.Printf("%s\n", x)
}
