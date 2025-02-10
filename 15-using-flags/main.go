package main

import (
	"flag"
	"fmt"
)

func main() {
	personAddress := flag.String("person", "", "Person name")
	ageAddress := flag.Int("age", 0, "Person åge")
	flag.Parse()

	fmt.Printf("hello, %d year old %s\n", *ageAddress, *personAddress)
}
