package main

import (
	"flag"
	"fmt"
)

func main() {
	person := flag.String("person", "", "Person name")
	age := flag.Int("age", 0, "Person åge")
	flag.Parse()

	fmt.Printf("hello, %d year old %s\n", *age, *person)
}
