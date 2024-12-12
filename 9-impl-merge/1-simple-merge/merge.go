package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jszwec/csvutil"
)

type ItemRow struct {
	ItemName  string
	Quantity  int
	IsNeeded  string
	AddedDate string
}

func main() {
	lhsFilename := ""
	rhsFilename := ""

	flag.StringVar(&rhsFilename, "rhs", "", "CSV filename on the right")
	flag.StringVar(&lhsFilename, "lhs", "", "CSV filename on the left")
	flag.Parse()

	fmt.Printf("LHS=%s, RHS=%s\n", lhsFilename, rhsFilename)

	lhsData, err := os.ReadFile(lhsFilename)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s\n", string(lhsData))

	rhsData, err := os.ReadFile(rhsFilename)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Printf("%s\n", string(rhsData))

	// Unmarshal == Put it back in my program structure
	var rows []ItemRow
	fmt.Printf("Rows = %d\n", len(rows))
	err = csvutil.Unmarshal(lhsData, &rows)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	fmt.Printf("Rows = %d\n", len(rows))
	fmt.Printf("3rd row %v\n", rows[2])
	fmt.Printf("All rows %v\n", rows)
}
