package main

import (
	"fmt"
	"os"
)

func f(d *[]byte) {
	fmt.Println(len(*d))
}

func g(d *[]byte) {
	i := 0
	for _, b := range *d {
		fmt.Println(b)
		if i > 10 {
			break
		}
		i++
	}
}

func main() {
	filename := "/Users/arunsundaram/data/investments/SHARADAR/SEP.csv"
	data, err := os.ReadFile(filename) // copy #1
	if err != nil {
		panic(err)
	}

	f(&data)
	g(&data)
}
