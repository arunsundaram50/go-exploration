package main

import "fmt"

var months = []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}

func printMonthsStartingFrom(startingMonth int, numberOfMonths int) {
	fmt.Printf("%v\n", months[startingMonth-1:startingMonth-1+numberOfMonths])
}

func main() {
	printMonthsStartingFrom(3, 2) // March, April
	printMonthsStartingFrom(6, 3) // June, July, Aug
}
