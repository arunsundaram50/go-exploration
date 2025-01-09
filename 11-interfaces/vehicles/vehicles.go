package vehicles

import "fmt"

type Car struct {
	Make  string
	Model string
	Year  int
	Mpg   float64
	Seats int
}

type Truck struct {
	Make       string
	Model      string
	Year       int
	Mpg        float64
	WheelCount int
}

func (car Car) GetType() string {
	return "car"
}

func (car Car) Details() string {
	s := fmt.Sprintf("Make = %s, Model = %s, Year = %d, Mpg = %f, Seats = %d\n", car.Make, car.Model, car.Year, car.Mpg, car.Seats)
	return s
}

func (truck Truck) Details() string {
	s := fmt.Sprintf("Make = %s, Model = %s, Year = %d, Mpg = %f, WheelCount = %d\n", truck.Make, truck.Model, truck.Year, truck.Mpg, truck.WheelCount)
	return s
}

func (truck Truck) GetType() string {
	return "truck"
}
