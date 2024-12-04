package main

import (
	"flag"
	"fmt"
)

func main() {
	lhsFilename := ""
	rhsFilename := ""

	flag.StringVar(&lhsFilename, "lhs", "", "CSV filename on the left")
	flag.StringVar(&rhsFilename, "rhs", "", "CSV filename on the right")
	flag.Parse()

	fmt.Printf("LHS=%s, RHS=%s\n", lhsFilename, rhsFilename)
}
