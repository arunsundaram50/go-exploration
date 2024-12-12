package main

import (
	"fmt"
)

func f(a int) { // take a int
	a = 100
}

func g(a *int) { // take an int address
	*a = 100 // *a means, the value in address a
}

func main() {

	x := 10 // RAM location = 100,000,000 for example; value is 10
	fmt.Printf("%d\n", x)

	f(x) // 10
	fmt.Printf("%d\n", x)

	g(&x) // 100,000,000
	fmt.Printf("%d\n", x)
}
