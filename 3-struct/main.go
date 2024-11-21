package main

import (
	"fmt"
)

type UserPii struct {
	year   int
	salary int
	hobby  string
}

var userPiiList = []UserPii{
	{2021, 1000, "Basket ball"},
	{2022, 2000, "Cricket"},
	{2023, 3000, "Cricket"},
	{2024, 4000, "Soccer"},
	{2025, 5000, "Volleyball"},
}

func printHobby(year int) {
	for _, userPii := range userPiiList {
		if userPii.year == year {
			println(userPii.hobby)
		}
	}
}

func main() {
	up1 := UserPii{2021, 1000, "Basket ball"}
	fmt.Printf("%v\n", up1)
	up2 := UserPii{2022, 2000, "Cricket"}
	fmt.Printf("%v\n", up2)

	printHobby(2024)

	fmt.Printf("xyz = %v\n", userPiiList)
}
