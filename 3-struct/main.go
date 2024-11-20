package main

import (
	"fmt"
)

type UserPii struct {
	year   int    `csv:"year"`
	salary int    `csv:"salary"`
	hobby  string `csv:"hobby"`
}

func printHobby(year int) {
	for _, up := range xyz {
		if up.year == year {
			println(up.hobby)
		}
	}
}

var xyz = []UserPii{
	{2021, 1000, "Basket ball"},
	{2022, 2000, "Cricket"},
	{2023, 3000, "Cricket"},
	{2024, 4000, "Soccer"},
	{2025, 5000, "Volleyball"},
}

func main() {
	up1 := UserPii{2021, 1000, "Basket ball"}
	fmt.Printf("%v\n", up1)
	up2 := UserPii{2022, 2000, "Cricket"}
	fmt.Printf("%v\n", up2)

	printHobby(2024)

	fmt.Printf("xyz = %v\n", xyz)
}
