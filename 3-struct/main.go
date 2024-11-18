package main

import "fmt"

type UserPii struct {
	year   int
	salary int
	hobby  string
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
	println(up1.salary)
	up2 := UserPii{2022, 2000, "Cricket"}
	println(up2.hobby)

	fmt.Printf("%v\n", xyz)

	printHobby(2024)
}
