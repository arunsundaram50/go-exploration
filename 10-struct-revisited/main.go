package main

import "fmt"

type UserPii struct {
	year   int
	salary int
	hobby  string
}

func (u UserPii) print() {
	fmt.Printf("Year = %d, Salary = %d, Hobby = %s\n", u.year, u.salary, u.hobby)
}

func print2(u UserPii) {
	fmt.Printf("Year = %d, Salary = %d, Hobby = %s\n", u.year, u.salary, u.hobby)
}

func main() {
	u1 := UserPii{salary: 10000, hobby: "skiing", year: 2002}

	fmt.Printf("%v\n", u1)

	u1.print()
	print2(u1)
}
