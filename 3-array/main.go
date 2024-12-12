package main

import "fmt"

var years = []int{2022, 2023, 2024}
var hobbies = []string{"shuttle cock", "basket ball", "cricket"} // array/slice
var salaries = []int{100000, 150000, 200000}

func findYearIndex(givenYear int) int {
	for i, year := range years {
		if year == givenYear {
			return i
		}
	}
	panic("year not found")
}

func getHobby(year int) string {
	index := findYearIndex(year)
	return hobbies[index]
}

func getHobbyByIndex(index int) string {
	return hobbies[index]
}

func getSalary(year int) int {
	index := findYearIndex(year)
	return salaries[index]
}

func main() {

	hobby := getHobby(2023)
	println("2023 hobby = ", hobby)

	pos := findYearIndex(2023)
	h := getHobbyByIndex(pos)
	println("2023 hobby = ", h)

	name := "Dhana" // scalar
	println(name)

	println(hobbies[0])             // prints scalar
	println(hobbies[1:])            // prints slice's address
	fmt.Printf("%v\n", hobbies[1:]) // prints slice
	fmt.Printf("%s\n", name)

	println(salaries[0])
	fmt.Printf("%v\n", salaries)
}
