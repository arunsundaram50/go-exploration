package main

import "fmt"

var months = []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}

func printMonthsStartingFrom(startingMonth int, numberOfMonths int) error {
	if startingMonth-1+numberOfMonths > 12 {
		return fmt.Errorf("You went beyond December")
	}

	fmt.Printf("%v\n", months[startingMonth-1:startingMonth-1+numberOfMonths])
	return nil
}

func main() {
	err1 := printMonthsStartingFrom(10, 3) // Oct, Nov, Dec
	if err1 != nil {
		println(err1)
	}

	err2 := printMonthsStartingFrom(10, 4)
	if err2 != nil {
		println(err2.Error())
	}
}
