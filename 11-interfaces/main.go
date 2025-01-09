package main

import (
	"11-interfaces/vehicles"
	"fmt"
)

type Vehicle interface {
	GetType() string
}

func CalculateTolls(vehicles []Vehicle) int {
	totalToll := 0
	for _, vehicle := range vehicles {
		if vehicle.GetType() == "car" {
			totalToll = totalToll + 10
		} else if vehicle.GetType() == "truck" {
			totalToll = totalToll + 50
		}
	}
	return totalToll
}

func main() {
	corolla := vehicles.Car{
		Make:  "Toyota",
		Model: "Corolla",
		Year:  2024,
		Mpg:   50,
		Seats: 4,
	}

	fmt.Println(corolla.Details())
	fmt.Println(corolla.GetType())

	xyzTruck := vehicles.Truck{
		Make:       "Chrysler",
		Model:      "XYZ",
		Year:       2023,
		Mpg:        18,
		WheelCount: 18,
	}
	fmt.Println(xyzTruck.Details())
	fmt.Println(xyzTruck.GetType())

	civic := vehicles.Car{
		Make:  "Honda",
		Model: "civic",
		Year:  2024,
		Mpg:   50,
		Seats: 4,
	}

	total := CalculateTolls([]Vehicle{corolla, civic, xyzTruck})
	fmt.Println("Total = ", total)
}
