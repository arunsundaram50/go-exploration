package main

import "fmt"

func f1(xAddress *int) {
	/// change the value at address xAddress
	*xAddress = 11
}

func main() {
	x := 10
	f1(&x)
	fmt.Println(x)
}
