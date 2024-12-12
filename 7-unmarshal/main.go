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

// print all white color items in myfile.csv
func main() {
	allItems := []KitchenItem{}

	bytes, err := os.ReadFile("myfile.csv")
	if err != nil {
		panic(err)
	}
	err = csvutil.Unmarshal(bytes, &allItems)
	if err != nil {
		panic(err)
	}

	for _, item := range allItems {
		if item.Color == "white" {
			fmt.Printf("%v\n", item)
		}
	}
}
