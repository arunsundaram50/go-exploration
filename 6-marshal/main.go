package main

import (
	"fmt"
	"os"

	"github.com/jszwec/csvutil"
)

type KitchenItem struct {
	Name           string
	Material       string
	WeightInLbs    float32
	Color          string
	HeightInInches float32
}

var kitchenItems = []KitchenItem{
	{"spoon#1", "silver", .3, "silver", 10},
	{"spoon#2", "plastic", .1, "white", 5},
	{"fork#1", "plastic", .1, "white", 5},
}

func main() {
	csvBytes, err := csvutil.Marshal(kitchenItems)
	if err != nil {
		fmt.Printf("error %v\n", err)
	} else {
		fmt.Printf("CSV = %s\n", string(csvBytes))
	}
	// prefix: 0o - oct, 0b -bin, 0x -hex, dec (32+128+256)
	os.WriteFile("myfile.csv", csvBytes, 0o644)
}
